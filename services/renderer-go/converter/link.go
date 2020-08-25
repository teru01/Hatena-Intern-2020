package converter

import (
	"fmt"
	"regexp"
	"strings"
)

type LinkConverter struct {
	Pattern *regexp.Regexp
}

func NewLinkConverter() *LinkConverter {
	return &LinkConverter{
		Pattern: regexp.MustCompile(`\[(.[^\]]*)\]\((https?://.[^\)]*)\)|(https?://.[^\s]*)`),
	}
}

func (lc *LinkConverter) convert(line string) (string, error) {
	matches := lc.Pattern.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return line, nil
	}
	for _, m := range matches {
		if m[3] != "" {
			// URLリンク直書きにマッチ
			line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, m[0], m[0]), 1)
			continue
		}
		// link記法にマッチ
		line = strings.Replace(line, m[0], fmt.Sprintf(`<a href="%s">%s</a>`, m[2], m[1]), 1)
	}
	return line, nil
}
