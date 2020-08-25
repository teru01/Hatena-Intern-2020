package renderer

import (
	"context"
	"regexp"
	"testing"

	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter"
	"github.com/stretchr/testify/assert"
)

type TC struct {
	in  string
	out string
}

func Test_Render(t *testing.T) {
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
			out: `こんにちは，<a href=http://お名前.com>hoge</a>.
クァwsedrftgyふじこ<a href=http://yahoo.com>ここをクリック</a>
`,
		},
		TC{
			in: `# タイトル
1つ目は[hoge](http://google.com)で，2つ目は[hoge](http://google.com)です．
`,
			out: `<h1>タイトル</h1>
1つ目は<a href=http://google.com>hoge</a>で，2つ目は<a href=http://google.com>hoge</a>です．
`,
		},
		TC{
			in: `# タイトル[hoge](http://google.com)
`,
			out: `<h1>タイトル<a href=http://google.com>hoge</a></h1>
`,
		},
	}

	for _, testCase := range testCases {
		lc, wc := NewConverters()
		html, err := Render(context.Background(), testCase.in, lc, wc)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, html)
	}
}

func NewConverters() ([]converter.LineConverter, []converter.WholeConverter) {
	return []converter.LineConverter{
			&converter.HeadingConverter{
				AllowedLevel: 5,
				Pattern:      regexp.MustCompile(`^(#+) .*`),
			},
			&converter.LinkConverter{
				Pattern: regexp.MustCompile(`\[(.[^\]]*)\]\((https?://.[^\)]*)\)`),
			},
		},
		[]converter.WholeConverter{}
}
