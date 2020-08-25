package converter

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []string{
		"# hogehoge",
		"### qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		"## あいう",
		"###### abc",
		"##0124",
		"## hoge##",
	}
	expected := []string{
		"<h1>hogehoge</h1>",
		"<h3>qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq</h3>",
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		"<h2>あいう</h2>",
		"<h5>abc</h5>",
		"##0124",
		"<h2>hoge##</h2>",
	}

	h := &HeadingConverter{
		AllowedLevel: 5,
		Pattern:      regexp.MustCompile(`^(#+) .*`),
	}
	for i, _ := range testCases {
		result, err := h.convert(testCases[i])
		assert.NoError(t, err)
		assert.Equal(t, expected[i], result)
	}
}