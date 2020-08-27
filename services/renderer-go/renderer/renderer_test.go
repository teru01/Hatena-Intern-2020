package renderer

import (
	"context"
	"testing"

	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter"
	"github.com/stretchr/testify/assert"
)

type TC struct {
	in  string
	out string
}

func TestRender(t *testing.T) {
	testCases := []TC{
		TC{
			in: `こんにちは，hello, world!
`,
			out: `こんにちは，hello, world!
`,
		},
		TC{
			in: `こんにちは，hello, world!🤔
`,
			out: `こんにちは，hello, world!🤔
`,
		},
		TC{
			in: `こんにちは，hello, world!

			
ハロー
`,
			out: `こんにちは，hello, world!

			
ハロー
`,
		},
		TC{
			in: `こんにちは，hello, world!
こんにちは，hello, world!
	こんにちは，hello, world!
`,
			out: `こんにちは，hello, world!
こんにちは，hello, world!
	こんにちは，hello, world!
`,
		},
		TC{
			in: `こんにちは，[hoge](http://お名前.com).
クァwsedrftgyふじこ[ここをクリック](http://yahoo.com)
`,
			out: `こんにちは，<a href="http://お名前.com">hoge</a>.
クァwsedrftgyふじこ<a href="http://yahoo.com">ここをクリック</a>
`,
		},
		TC{
			in: `# タイトル
1つ目は[hoge](http://google.com)で，2つ目は[hoge](http://google.com)です．
`,
			out: `<h1>タイトル</h1>
1つ目は<a href="http://google.com">hoge</a>で，2つ目は<a href="http://google.com">hoge</a>です．
`,
		},
		TC{
			in: `# タイトル[hoge](http://google.com)
`,
			out: `<h1>タイトル<a href="http://google.com">hoge</a></h1>
`,
		},
		TC{
			in: `- [hoge](http://google.com)
`,
			out: `<ul><li><a href="http://google.com">hoge</a></li></ul>
`,
		},
		TC{
			in: `- [hoge](http://google.com)
	- [bar](http://google.com)
# title
`,
			out: `<ul><li><a href="http://google.com">hoge</a></li><ul><li><a href="http://google.com">bar</a></li></ul></ul>
<h1>title</h1>
`,
		},
		TC{
			in: `- [hoge](http://google.com)
	- [](http://google.com)または[hoge](http://google.com)
# title
`,
			out: `<ul><li><a href="http://google.com">hoge</a></li><ul><li><a href="http://google.com">success</a>または<a href="http://google.com">hoge</a></li></ul></ul>
<h1>title</h1>
`,
		},
	}

	lc, wc := converter.NewConverters(&converter.DummyFetchClient{})
	for _, testCase := range testCases {
		html, err := Render(context.Background(), testCase.in, lc, wc)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}

func TestURLCache(t *testing.T) {
	tc := TC{
		in:  `1つ目は[](http://google.com)で，2つ目は[](http://google.com)です．
3つ目は[](http://google.com)
4つ目は[](http://google.com)
`,
		out: `1つ目は<a href="http://google.com">success</a>で，2つ目は<a href="http://google.com">success</a>です．
3つ目は<a href="http://google.com">success</a>
4つ目は<a href="http://google.com">success</a>
`,
	}
	fc := converter.DummyFetchClient{}
	lc, wc := converter.NewConverters(&fc)
	html, err := Render(context.Background(), tc.in, lc, wc)
	assert.NoError(t, err)
	assert.Equal(t, tc.out, html)
	assert.Equal(t, fc.CallCount(), 1) // 2度目以降は呼ばれずキャッシュが使われる
}

