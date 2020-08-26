package converter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type LineConverter interface {
	convertLine(src string) (string, error)
}

type WholeConverter interface {
	convertText(src string) (string, error)
}

func Execute(text string, lineConverters []LineConverter, wholeConverters []WholeConverter) (string, error) {
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
			line, err = lc.convertLine(strings.TrimRight(line, "\n"))
			if err != nil {
				return "", err
			}
		}
		fmt.Fprintf(&builder, line+"\n")
	}
	convertedText := builder.String()

	for _, wc := range wholeConverters {
		var err error
		convertedText, err = wc.convertText(convertedText)
		if err != nil {
			return "", err
		}
	}
	return convertedText, nil
}

// ここに書くのはあまり良く無いかも
func NewConverters(f pb_fetcher.FetcherClient) ([]LineConverter, []WholeConverter) {
	return []LineConverter{
			NewHeadingConverter(5),
			NewLinkConverter(f),
		},
		[]WholeConverter{
			NewListConverter(),
		}
}
