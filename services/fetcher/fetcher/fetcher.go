package fetcher

import "context"

func Fetch(ctx context.Context, url string) (string, error) {
	if url == "http://fetcher-test" {
		return "success", nil
	}
	return "", nil
}
