package converter

import (
	"context"
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
			out: `<a href="http://google.com">hoge</a>`,
		},
		TC{
			in:  "[ほげ](http://google.com)",
			out: `<a href="http://google.com">ほげ</a>`,
		},
		TC{
			in:  "[hoge](http://お名前.com)",
			out: `<a href="http://お名前.com">hoge</a>`,
		},
		TC{
			in:  "1つ目は[hoge](http://google.com)で，2つ目は[hoge](http://google.com)です．",
			out: `1つ目は<a href="http://google.com">hoge</a>で，2つ目は<a href="http://google.com">hoge</a>です．`,
		},
		TC{
			in:  "1つ目はhttp://google.com",
			out: `1つ目は<a href="http://google.com">http://google.com</a>`,
		},
		TC{
			in:  "1つ目はhttp://google.com で，2つ目は[hoge](http://google.com)です．",
			out: `1つ目は<a href="http://google.com">http://google.com</a> で，2つ目は<a href="http://google.com">hoge</a>です．`,
		},
		TC{
			in:  "こんにちは，[](http://google.com)です．",
			out: `こんにちは，<a href="http://google.com">success</a>です．`,
		},
		TC{
			in:  "こんにちは，[](http://google.com)です．2つ目は[foo](http://bar.com),3つ目は[](http://bar.com)",
			out: `こんにちは，<a href="http://google.com">success</a>です．2つ目は<a href="http://bar.com">foo</a>,3つ目は<a href="http://bar.com">success</a>`,
		},
	}

	lc := NewLinkConverter(&DummyFetchClient{})
	for _, testCase := range testCases {
		result, err := lc.convertLine(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, result)
	}
}
