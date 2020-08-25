package converter

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TC struct {
	in  string
	out string
}

func TestLink(t *testing.T) {
	testCases := []TC{
		TC{
			in:  "[hoge](http://google.com)",
			out: "<a href=http://google.com>hoge</a>",
		},
		TC{
			in:  "[ほげ](http://google.com)",
			out: "<a href=http://google.com>ほげ</a>",
		},
		TC{
			in:  "[hoge](http://お名前.com)",
			out: "<a href=http://お名前.com>hoge</a>",
		},
		TC{
			in:  "1つ目は[hoge](http://google.com)で，2つ目は[hoge](http://google.com)です．",
			out: "1つ目は<a href=http://google.com>hoge</a>で，2つ目は<a href=http://google.com>hoge</a>です．",
		},
	}

	lc := &LinkConverter{
		Pattern: regexp.MustCompile(`\[(.[^\]]*)\]\((https?://.[^\)]*)\)`),
	}
	for _, testCase := range testCases {
		result, err := lc.convert(testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, result)
	}
}
