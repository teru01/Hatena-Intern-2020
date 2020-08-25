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
	}

	lc, wc := converter.NewConverters()
	for _, testCase := range testCases {
		html, err := Render(context.Background(), testCase.in, lc, wc)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}
