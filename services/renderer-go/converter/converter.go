package converter

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	pb_fetcher "github.com/hatena/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

type LineConverter interface {
	convertLine(ctx context.Context, src string) (string, error)
}

type WholeConverter interface {
	convertText(ctx context.Context, src string) (string, error)
}

func Execute(ctx context.Context, text string, lineConverters []LineConverter, wholeConverters []WholeConverter) (string, error) {
	reader := bufio.NewReader(strings.NewReader(text))
	var builder strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		for _, lc := range lineConverters {
			line, err = lc.convertLine(ctx, strings.TrimRight(line, "\n"))
			if err != nil {
				return "", err
			}
		}
		fmt.Fprintf(&builder, line+"\n")
	}
	convertedText := builder.String()

	for _, wc := range wholeConverters {
		var err error
		convertedText, err = wc.convertText(ctx, convertedText)
		if err != nil {
			return "", err
		}
	}
	return convertedText, nil
}

// ここに書くのはあまり良く無いかも
func NewConverters(f pb_fetcher.FetcherClient) ([]LineConverter, []WholeConverter, error) {
	uploader, err := NewAWSUploder()
	if err != nil {
		return nil, nil, err
	}
	return []LineConverter{
			NewHeadingConverter(5),
			NewLinkConverter(f),
			NewImageConverter(uploader),
		},
		[]WholeConverter{
			NewListConverter(),
		}, nil
}
