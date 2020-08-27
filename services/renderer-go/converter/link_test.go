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

func TestLinkTitleCache(t *testing.T) {
	tc := TC{
		in:  "1つ目は[](http://google.com)で，2つ目は[](http://google.com)です．",
		out: `1つ目は<a href="http://google.com">success</a>で，2つ目は<a href="http://google.com">success</a>です．`,
	}
	lc := NewLinkConverter(&DummyFetchClient{callCount: 0})
	result, err := lc.convertLine(context.Background(), tc.in)
	assert.NoError(t, err)
	assert.Equal(t, tc.out, result)
	_, found := lc.cache.Get("http://google.com")
	assert.Equal(t, true, found)
	assert.Equal(t, lc.fetcherClient.(*DummyFetchClient).callCount, 1) // 2度目は呼ばれずキャッシュが使われる

	tc = TC{
		in:  "1つ目は[](http://google.com)で，2つ目は[](https://google.com/hogehoge)です．",
		out: `1つ目は<a href="http://google.com">success</a>で，2つ目は<a href="https://google.com/hogehoge">success</a>です．`,
	}
	lc = NewLinkConverter(&DummyFetchClient{callCount: 0})
	result, err = lc.convertLine(context.Background(), tc.in)
	assert.NoError(t, err)
	assert.Equal(t, tc.out, result)
	assert.Equal(t, 2, lc.fetcherClient.(*DummyFetchClient).callCount) // 異なるURIならキャッシュは使われない
}

func TestExtractFetchTargetURL(t *testing.T) {
	testCases := [][][]string{
		[][]string{
			[]string{"[title](http://google.com)", "", "http://google.com", ""},
			[]string{"[title](http://yahoo.com)", "", "http://yahoo.com", ""},
			[]string{"[title](http://amazon.com)", "", "http://amazon.com", ""},
		},
		[][]string{
			[]string{"[title](http://google.com)", "", "http://google.com", ""},
			[]string{"[title](http://yahoo.com)", "", "http://yahoo.com", ""},
			[]string{"[title](http://google.com)", "", "http://google.com", ""},
		},
		[][]string{
			[]string{"http://google.com", "", "", "http://google.com"},
			[]string{"[title](http://google.com)", "title", "http://google.com", ""},
			[]string{"[](http://google.com)", "", "http://google.com", ""},
		},
		[][]string{
			[]string{"http://google.com", "", "", "http://google.com"},
			[]string{"http://google.com", "", "", "http://google.com"},
			[]string{"http://google.com", "", "", "http://google.com"},
		},
	}
	expected := []map[string]struct{}{
		map[string]struct{}{
			"http://google.com": struct{}{},
			"http://yahoo.com":  struct{}{},
			"http://amazon.com": struct{}{},
		},
		map[string]struct{}{
			"http://google.com": struct{}{},
			"http://yahoo.com":  struct{}{},
		},
		map[string]struct{}{
			"http://google.com": struct{}{},
		},
		map[string]struct{}{},
	}

	for i, testCase := range testCases {
		urlSet := extractFetchTargetURL(testCase)
		assert.Equal(t, expected[i], urlSet)
	}
}
