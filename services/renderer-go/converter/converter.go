package converter

import (
	"bufio"
	"io"
	"strings"

	"github.com/golang/go/src/fmt"
	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter/htmlizer"
)

func Execute(text string, lineConverters []htmlizer.LineConverter, wholeConverters []htmlizer.WholeConverters) (string, error) {
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
			convertedLine, err := lc.convert(line)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&builder, convertedLine + "\n")
		}
	}
	convertedText := builder.String()

	for _, wc := range wholeConverters {
		var err error
		convertedText, err = wc.convert(convertedText)
		if err != nil {
			return "", err
		}
	}
	return convertedText, nil
}
