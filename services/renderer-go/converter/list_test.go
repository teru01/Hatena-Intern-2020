package converter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	testCases := []TC{
		TC{
			in: `- hoge
- 北海道
`,
			out: `<ul><li>hoge</li><li>北海道</li></ul>
`,
		},
		TC{
			in: `- hoge
	- bar
`,
			out: `<ul><li>hoge</li><ul><li>bar</li></ul></ul>
`,
		},
		TC{
			in: `- hoge
	- foo
- qwe
`,
			out: `<ul><li>hoge</li><ul><li>foo</li></ul><li>qwe</li></ul>
`,
		},
		TC{
			in: `- hoge
		- qwe
`,
			out: `<ul><li>hoge</li><ul><ul><li>qwe</li></ul></ul></ul>
`,
		},
		TC{
			in: `	- hoge
		- qwe
`,
			out: `<ul><ul><li>hoge</li><ul><li>qwe</li></ul></ul></ul>
`,
		},
		TC{
			in: `- hoge
	- qwe
hoge
bar
`,
			out: `<ul><li>hoge</li><ul><li>qwe</li></ul></ul>
hoge
bar
`,
		},
	}

	lc := NewListConverter()
	for _, testCase := range testCases {
		result, err := lc.convertText(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, result)
	}
}
