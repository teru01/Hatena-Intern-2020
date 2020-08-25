package converter

import (
	"fmt"
	"math"
	"regexp"
)

type HeadingConverter struct {
	AllowedLevel int
}

var pattern *regexp.Regexp

func init() {
	pattern = regexp.MustCompile(`^(#+) .*`)
}

func (c *HeadingConverter) convert(line string) (string, error) {
	matches := pattern.FindStringSubmatchIndex(line)
	if len(matches) == 0 {
		return line, nil
	}
	h := c.sharpNumToHeadNum(matches[3])
	return fmt.Sprintf("<h%d>", h) + line[matches[3]+1:] + fmt.Sprintf("</h%d>", h), nil
}

func (c *HeadingConverter) sharpNumToHeadNum(n int) int {
	return int(math.Min(float64(n), float64(c.AllowedLevel)))
}
