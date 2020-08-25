package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []string{
		"#hogehoge",
		"###qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		`#hogehoge
##aaawwws`,
	}
	expected := []string{
		"<h1>hogehoge</h1>",
		"<h3>qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq</h3>",
		"qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
		`<h1>hogehoge</h1>
<h2>aaawwws</h2>`,
	}

	h := &HeadingConverter{
		AllowedLevel: 5,
	}
	for i, _ := range testCases {
		result, err := h.convert(testCases[i])
		assert.NoError(t, err)
		assert.Equal(t, expected[i], result)
	}
}
