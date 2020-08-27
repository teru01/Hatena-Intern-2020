package converter

import (
	"context"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyUploader struct {
}

func (du *DummyUploader) Upload(data image.Image) (string, error) {
	return "http://locahost/test.png", nil
}

func TestImageConverter(t *testing.T) {
	testCases := []TC{
		TC{
			in:  `hoge\aaa/bar`,
			out: `hoge<img src=http://locahost/test.png alt=aaa>bar`,
		},
		TC{
			in:  `hoge\ほげ/bar`,
			out: `hoge<img src=http://locahost/test.png alt=ほげ>bar`,
		},
		TC{
			in:  `hoge\aaa/bar\hgoehogehoge/aaa`,
			out: `hoge<img src=http://locahost/test.png alt=aaa>bar<img src=http://locahost/test.png alt=hgoehogehoge>aaa`,
		},
	}

	ic := NewImageConverter(&DummyUploader{})
	for _, testCase := range testCases {
		output, err := ic.convertLine(context.Background(), testCase.in)
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, output)
	}
}
