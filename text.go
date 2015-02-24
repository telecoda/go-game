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

func (r renderController) RenderText(fontAssetId string, text string, pos sdl.Point, angle float64, textColor sdl.Color, vAlign VAlign, hAlign HAlign) error {

	font, err := getFont(fontAssetId)
	if err != nil {
		return err
	}

	//font.SetOutline(2)
	//textSurface := font.RenderText_Blended(text, textColor)
	textSurface := font.RenderText_Blended(text, textColor)

	//backColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	//textSurface := font.RenderText_Shaded(text, textColor, backColor)

	texture, err := r.Renderer.CreateTextureFromSurface(textSurface)
	defer textSurface.Free()
	defer texture.Destroy()

	if err != nil {
		return fmt.Errorf("Error rendering text:%s\n", err)
	}

	var textY, textX, centreX, centreY int32

	textWidth := textSurface.W
	textHeight := textSurface.H

	// calc Y pos
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

	// calc X pos
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

	// calc centre Y
	switch vAlign {
	case TOP:
		centreY = 0
	case MIDDLE:
		centreY = textHeight / 2
	case ABS_MIDDLE:
		centreY = textHeight / 2
	case BOTTOM:
		centreY = textHeight
	}

	// calc centre X
	switch hAlign {
	case LEFT:
		centreX = 0
	case CENTER:
		centreX = textWidth / 2
	case ABS_CENTER:
		centreX = textWidth / 2
	case RIGHT:
		centreX = textWidth
	}

	center := sdl.Point{centreX, centreY}

	src := sdl.Rect{0, 0, textSurface.W, textSurface.H}
	dst := sdl.Rect{textX, textY, textSurface.W, textSurface.H}
	//r.Renderer.Copy(texture, &src, &dst)
	//angle := 0.0
	r.Renderer.CopyEx(texture, &src, &dst, angle, &center, sdl.FLIP_NONE)

	return nil
}
