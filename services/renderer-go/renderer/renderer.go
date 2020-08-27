package renderer

import (
	"context"

	"github.com/hatena/Hatena-Intern-2020/services/renderer-go/converter"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string, lcs []converter.LineConverter, wcs []converter.WholeConverter) (string, error) {
	return converter.Execute(ctx, src, lcs, wcs)
}
