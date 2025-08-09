package resource

import (
	"fmt"
	"image/color"
	_ "image/png"

	"github.com/fogleman/gg"
	"github.com/mechiko/walk"
	"golang.org/x/image/font"
)

func (r *resourceData) Font(name string, points float64) (font.Face, error) {
	switch name {
	default:
		return nil, fmt.Errorf("data:resource not found")
	}
}

func (r *resourceData) ImageFontRune(rn rune, font string, points float64, size int, col color.RGBA) (*walk.Icon, error) {
	dc := gg.NewContext(size, size)
	// dc.SetRGB(1, 1, 1)
	// dc.Clear()
	dc.SetRGB(float64(col.R), float64(col.G), float64(col.B))
	if font, err := r.Font(font, points); err != nil {
		return nil, fmt.Errorf("data:resource error %w", err)
	} else {
		dc.SetFontFace(font)
		dc.DrawString(string(rn), 4, float64(size-10))
		if icon, err := walk.NewIconFromImageForDPI(dc.Image(), 96); err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		} else {
			return icon, nil
		}
	}
}
