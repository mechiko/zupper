package resource

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"

	"zupper/domain"
	"zupper/gui/resource/svg"

	_ "github.com/biessek/golang-ico"
	"github.com/mechiko/walk"
)

type resourceData struct {
	domain.Apper
}

func New(app domain.Apper) *resourceData {
	return &resourceData{
		Apper: app,
	}
}

func (r *resourceData) Icon(name string) (*walk.Icon, error) {
	var img image.Image
	var format string
	var err error
	switch name {
	case DefaultTreeNode:
		img, format, err = image.Decode(bytes.NewReader(icoUsers))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case DefaultTreeGroup:
		img, format, err = image.Decode(bytes.NewReader(icoFolder))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case DefaultTreeEmpty:
		img, format, err = image.Decode(bytes.NewReader(icoUsers))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case PngKfh:
		img, format, err = image.Decode(bytes.NewReader(pngKfh))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case PngLogo:
		img, format, err = image.Decode(bytes.NewReader(pngLogo))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case PngRequester:
		img, format, err = image.Decode(bytes.NewReader(pngRequester))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	// case IcoUsers:
	// 	if icon, err := r.ImageFontRune('\uF046', FontMedium, 64, 64, color.RGBA{R: 125, A: 1}); err != nil {
	// 		return nil, fmt.Errorf("data:resource %w", err)
	// 	} else {
	// 		return icon, nil
	// 	}
	// img, format, err = image.Decode(bytes.NewReader(icoUsers))
	// if err != nil {
	// 	return nil, fmt.Errorf("data:resource %w", err)
	// }
	// case IcoFolder:
	// 	if icon, err := r.ImageFontRune('\uEA83', FontMedium, 64, 64, color.RGBA{R: 125, A: 1}); err != nil {
	// 		return nil, fmt.Errorf("data:resource %w", err)
	// 	} else {
	// 		return icon, nil
	// 	}
	// img, format, err = image.Decode(bytes.NewReader(icoFolder))
	// if err != nil {
	// 	return nil, fmt.Errorf("data:resource %w", err)
	// }
	default:
		icon, err := walk.Resources.Icon(name)
		if err != nil {
			return nil, fmt.Errorf("resource:icon %w", err)
		}
		return icon, nil
	}
	r.Logger().Debugf("data:resource %s format %s", name, format)
	icon, err := walk.NewIconFromImageForDPI(img, 96)
	if err != nil {
		// r.app.Logger().Errorf("data:resource walk error %s", err.Error())
		return nil, fmt.Errorf("data:resource %w", err)
	}
	return icon, nil
}

func (r *resourceData) Svg(name string, c color.Color, h, w int) (*walk.Icon, error) {
	var svgDecoder *svg.Decoder
	var img image.Image
	var err error
	switch name {
	case SvgCircle:
		svgDecoder, err = svg.NewDecoder(bytes.NewReader(svg.Colorize(svgCircle, c)))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	case SvgRequest:
		svgDecoder, err = svg.NewDecoder(bytes.NewReader(svg.Colorize(svgRequest, c)))
		if err != nil {
			return nil, fmt.Errorf("data:resource %w", err)
		}
	default:
		return nil, fmt.Errorf("data:resource not found")
	}
	img, err = svgDecoder.Draw(w, h)
	if err != nil {
		return nil, fmt.Errorf("data:resource %w", err)
	}
	icon, err := walk.NewIconFromImageForDPI(img, 96)
	if err != nil {
		return nil, fmt.Errorf("data:resource %w", err)
	}
	return icon, nil
}
