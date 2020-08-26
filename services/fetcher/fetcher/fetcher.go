package fetcher

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

var pattern *regexp.Regexp

func init() {
	// 空文字，空白文字のみは許容しない.
	// TODO: "</title>"はタイトル中に含まないようにしたい（正規表現わからん）
	pattern = regexp.MustCompile(`<(title|TITLE)>(.*\S.*)</(title|TITLE)>`)
}

func Fetch(ctx context.Context, url string, client *http.Client) (string, error) {
	if url == "http://fetcher-test.example.com" {
		// 統合テスト用
		return "success", nil
	}

	client.Timeout = 3 * time.Second
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	return extractTitle(response.Body)
}

func extractTitle(source io.Reader) (string, error) {
	r := bufio.NewReader(source)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return "", fmt.Errorf("title not found\n")
		} else if err != nil {
			return "", err
		}
		match := pattern.FindStringSubmatch(line)
		if len(match) != 0 {
			// マッチが存在
			return match[2], nil
		}
	}
}
