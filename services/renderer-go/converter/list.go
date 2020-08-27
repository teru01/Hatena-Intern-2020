package converter

import (
	"context"
	"regexp"
	"strings"
)

type ListConverter struct {
	Pattern *regexp.Regexp
}

func NewListConverter() *ListConverter {
	return &ListConverter{
		Pattern: regexp.MustCompile(`^(\t*)- .*`),
	}
}

// テキスト全体を受け取り，変換して返す
func (lc *ListConverter) convertText(ctx context.Context, text string) (string, error) {
	lines := strings.Split(text, "\n")
	var builder strings.Builder
	prevNumIndent := -1
	currentNumIndent := 0
	isListZone := false
	nLines := len(lines)
	for i, line := range lines {
		r := lc.Pattern.FindStringSubmatchIndex(line)
		if len(r) == 0 {
			if isListZone {
				builder.WriteString(strings.Repeat("</ul>", currentNumIndent+1))
				builder.WriteString("\n")
				isListZone = false
			}
			if i == nLines-1 {
				builder.WriteString(line)
			} else {
				builder.WriteString(line + "\n")
			}
			continue
		}
		isListZone = true
		currentNumIndent = r[3]
		if currentNumIndent > prevNumIndent {
			builder.WriteString(strings.Repeat("<ul>", currentNumIndent-prevNumIndent))
		} else if currentNumIndent < prevNumIndent {
			builder.WriteString(strings.Repeat("</ul>", prevNumIndent-currentNumIndent))
		}
		builder.WriteString("<li>" + line[r[3]+2:] + "</li>")
		prevNumIndent = currentNumIndent
	}
	return builder.String(), nil
}
