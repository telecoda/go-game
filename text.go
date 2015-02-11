package gogame

import (
	"fmt"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func (r renderController) RenderText(assetId string, text string, pos sdl.Point, textColor sdl.Color) error {

	font, err := gameAssets.getFontAsset(assetId)
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

	src := sdl.Rect{0, 0, textSurface.W, textSurface.H}
	dst := sdl.Rect{int32(pos.X), int32(pos.Y), textSurface.W, textSurface.H}
	r.Renderer.Copy(texture, &src, &dst)

	return nil
}
