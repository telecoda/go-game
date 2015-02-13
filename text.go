package gogame

import (
	"fmt"

	sdl "github.com/veandco/go-sdl2/sdl"
)

type VAlign string

const (
	TOP        VAlign = "top"
	ABS_MIDDLE VAlign = "absmiddle"
	MIDDLE     VAlign = "middle"
	BOTTOM     VAlign = "bottom"
)

type HAlign string

const (
	LEFT       HAlign = "left"
	ABS_CENTER HAlign = "abscenter"
	CENTER     HAlign = "center" // <--- centre
	RIGHT      HAlign = "right"
)

func (r renderController) RenderText(assetId string, text string, pos sdl.Point, textColor sdl.Color, vAlign VAlign, hAlign HAlign) error {

	font, err := getFont(assetId)
	if err != nil {
		return err
	}
	textSurface := font.RenderText_Solid(text, textColor)
	texture, err := r.Renderer.CreateTextureFromSurface(textSurface)
	defer textSurface.Free()
	defer texture.Destroy()

	if err != nil {
		return fmt.Errorf("Error rendering text:%s\n", err)
	}

	var textY, textX int32

	textWidth := textSurface.W
	textHeight := textSurface.H

	switch vAlign {
	case TOP:
		textY = pos.Y
	case MIDDLE:
		textY = pos.Y - textHeight/2
	case ABS_MIDDLE:
		textY = int32(r.height/2) - textHeight/2
	case BOTTOM:
		textY = pos.Y - textHeight
	}

	switch hAlign {
	case LEFT:
		textX = pos.X
	case CENTER:
		textX = pos.X - textWidth/2
	case ABS_CENTER:
		textX = int32(r.width/2) - textWidth/2
	case RIGHT:
		textX = pos.X - textWidth
	}

	src := sdl.Rect{0, 0, textSurface.W, textSurface.H}
	dst := sdl.Rect{textX, textY, textSurface.W, textSurface.H}
	r.Renderer.Copy(texture, &src, &dst)

	return nil
}
