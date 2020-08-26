package converter

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

type LinkConverter struct {
	Pattern       *regexp.Regexp
	fetcherClient pb_fetcher.FetcherClient
}

func NewLinkConverter(fetcherCli pb_fetcher.FetcherClient) *LinkConverter {
	return &LinkConverter{
		Pattern:       regexp.MustCompile(`\[(.[^\]]*)\]\((https?://.[^\)]*)\)|(https?://.[^\s]*)`),
		fetcherClient: fetcherCli,
	}
}

func (lc *LinkConverter) convertLine(ctx context.Context, line string) (string, error) {
	matches := lc.Pattern.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return line, nil
	}
	for _, m := range matches {
		matchTitle := m[1]
		matchURLWithTitle := m[2] // (title)[url]記法で記述された際のurlにマッチするもの
		matchOnlyURL := m[3]      // URLのみ記述した際のURLにマッチ
		if matchOnlyURL != "" {
			// URLリンク直書きにマッチ
			line = strings.Replace(line, matchOnlyURL, fmt.Sprintf(`<a href="%s">%s</a>`, matchOnlyURL, matchOnlyURL), 1)
			continue
		}
		if matchTitle == "" {
			// link記法 w/o titleにマッチ
			reply, _ := lc.fetcherClient.Fetch(ctx, &pb_fetcher.FetchRequest{Uri: matchURLWithTitle})
			line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, matchURLWithTitle, reply.Title), 1)
			continue
		}
		// link記法 w/ titleにマッチ
		line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, matchURLWithTitle, matchTitle), 1)
	}
	return line, nil
}
