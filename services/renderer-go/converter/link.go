package converter

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	cache "github.com/patrickmn/go-cache"
)

type LinkConverter struct {
	Pattern       *regexp.Regexp
	fetcherClient pb_fetcher.FetcherClient
	cache         *cache.Cache
}

func NewLinkConverter(fetcherCli pb_fetcher.FetcherClient) *LinkConverter {
	return &LinkConverter{
		Pattern:       regexp.MustCompile(`\[([^\]]*)\]\((https?://[^\)]*)\)|(https?://[^\s]*)`),
		fetcherClient: fetcherCli,
		cache:         cache.New(time.Hour*5, time.Hour*10),
	}
}

func (lc *LinkConverter) convertLine(ctx context.Context, line string) (string, error) {
	matches := lc.Pattern.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return line, nil
	}
	var wg sync.WaitGroup

	fetchTargetUrlSet := extractFetchTargetURL(matches)
	for url, _ := range fetchTargetUrlSet {
		// 先にタイトルを並行に取得してキャッシュにセットしておく
		// replaceと同時にやるとデータ競合しそうなので
		wg.Add(1)
		go func(url string) {
			_, found := lc.cache.Get(url)
			if !found {
				lc.cache.Set(url, lc.fetchTitle(ctx, url), cache.DefaultExpiration)
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
	for _, m := range matches {
		matchTitle := m[1]
		matchURLWithTitle := m[2] // (title)[url]記法で記述された際のurlにマッチするもの
		matchOnlyURL := m[3]      // URLのみ記述した際のURLにマッチ
		if matchOnlyURL != "" {
			// URLリンク直書きにマッチ
			line = strings.Replace(line, matchOnlyURL, fmt.Sprintf(`<a href="%s">%s</a>`, matchOnlyURL, matchOnlyURL), 1)
			continue
		} else if matchTitle == "" {
			// link記法 w/o titleにマッチ
			title, found := lc.cache.Get(matchURLWithTitle)
			if !found {
				title = lc.fetchTitle(ctx, matchURLWithTitle)
			}
			line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, matchURLWithTitle, title), 1)
			continue
		}
		// link記法 w/ titleにマッチ
		line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, matchURLWithTitle, matchTitle), 1)
	}
	return line, nil
}

func (lc *LinkConverter) fetchTitle(ctx context.Context, url string) string {
	reply, err := lc.fetcherClient.Fetch(ctx, &pb_fetcher.FetchRequest{Uri: url})
	title := "unknown title"
	if err == nil {
		title = reply.Title
	}
	return title
}

func extractFetchTargetURL(matches [][]string) map[string]struct{} {
	urlSet := make(map[string]struct{})
	for _, m := range matches {
		if m[1] == "" && m[3] == "" {
			// タイトルが空
			urlSet[m[2]] = struct{}{}
		}
	}
	return urlSet
}
