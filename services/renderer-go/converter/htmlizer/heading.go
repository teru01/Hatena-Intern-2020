package htmlizer

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type HeadingConverter struct {
	AllowedLevel int
}

var pattern *regexp.Regexp

func init() {
	pattern = regexp.MustCompile(`^(#+)`)
}

func (c *HeadingConverter) convert(src string) (string, error) {
	lines := strings.Split(src, "\n")
	var resultLines []string
	for _, line := range lines {
		matches := pattern.FindStringSubmatchIndex(line)
		if len(matches) == 0 {
			resultLines = append(resultLines, line)
			continue
		}
		h := c.sharpNumToHeadNum(matches[1])
		resultLines = append(resultLines, fmt.Sprintf("<h%d>", h)+line[matches[1]:]+fmt.Sprintf("</h%d>", h))
	}
	return strings.Join(resultLines, "\n"), nil
}

func (c *HeadingConverter) sharpNumToHeadNum(n int) int {
	return int(math.Min(float64(n), float64(c.AllowedLevel)))
}
