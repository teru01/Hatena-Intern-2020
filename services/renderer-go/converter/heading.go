package converter

import (
	"fmt"
	"math"
	"regexp"
)

type HeadingConverter struct {
	AllowedLevel int
	Pattern      *regexp.Regexp
}

func (hc *HeadingConverter) convert(line string) (string, error) {
	matches := hc.Pattern.FindStringSubmatchIndex(line)
	if len(matches) == 0 {
		return line, nil
	}
	h := hc.sharpNumToHeadNum(matches[3])
	return fmt.Sprintf("<h%d>", h) + line[matches[3]+1:] + fmt.Sprintf("</h%d>", h), nil
}

func (hc *HeadingConverter) sharpNumToHeadNum(n int) int {
	return int(math.Min(float64(n), float64(hc.AllowedLevel)))
}
