package converter

import (
	"context"
	"fmt"
	"image"
	"os"
	"regexp"
	"strings"

	"github.com/fogleman/gg"
)

type ImageConverter struct {
	Pattern  *regexp.Regexp
	uploader Uploader
}

func NewImageConverter(u Uploader) *ImageConverter {
	return &ImageConverter{
		Pattern:  regexp.MustCompile(`\\(.*?)/`),
		uploader: u,
	}
}

func (ic *ImageConverter) convertLine(ctx context.Context, line string) (string, error) {
	matches := ic.Pattern.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return line, nil
	}

	for _, m := range matches {
		imageStr := m[1]
		img, err := ic.createImage(imageStr)
		if err != nil {
			return line, err
		}
		path, err := ic.uploader.Upload(img)
		if err != nil {
			// 画像にできなかったものは無視
			continue
		}
		line = strings.Replace(line, m[0], fmt.Sprintf("<img src=%s alt=%s>", path, imageStr), 1)
	}
	return line, nil
}

func (ic *ImageConverter) createImage(imageStr string) (image.Image, error) {
	size := 96
	message := imageStr
	im := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size * len([]rune(message))}})
	dc := gg.NewContext(size*len([]rune(message)), size)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(os.Getenv("FONT_FILE_PATH"), float64(size)); err != nil {
		return nil, err
	}

	dc.DrawStringAnchored(message, float64(size*len([]rune(message))/2), float64(size/2), 0.5, 0.5)
	dc.DrawImage(im, 0, 0)
	return dc.Image(), nil
}
