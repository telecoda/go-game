package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (r renderController) RenderTexture(assetId string, pos sdl.Point, textureWidth, textureHeight int32) error {

	image, texture, err := gameAssets.getImageAsset(assetId)
	if err != nil {
		return err
	}

	width := image.W
	height := image.H

	src := sdl.Rect{0, 0, width, height}

	dst := sdl.Rect{pos.X, pos.Y, textureWidth, textureHeight}
	r.Renderer.Copy(texture, &src, &dst)

	return nil
}

func (r renderController) RenderRotatedTexture(assetId string, pos sdl.Point, rotation float64, textureWidth, textureHeight int32) error {

	image, texture, err := gameAssets.getImageAsset(assetId)
	if err != nil {
		return err
	}

	width := image.W
	height := image.H

	return renderRotatedTexture(texture, pos, rotation, width, height, textureWidth, textureHeight)

}

func renderRotatedTexture(texture *sdl.Texture, pos sdl.Point, rotation float64, width, height, textureWidth, textureHeight int32) error {

	if texture == nil {
		return fmt.Errorf("Error: texture is nil")
	}
	src := sdl.Rect{0, 0, width, height}
	centre := sdl.Point{textureWidth / 2, textureHeight / 2}
	dst := sdl.Rect{pos.X, pos.Y, textureWidth, textureHeight}
	rendCont.Renderer.CopyEx(texture, &src, &dst, rotation, &centre, sdl.FLIP_NONE)

	return nil
}

func (r renderController) ClearScreen(color sdl.Color) {
	r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	r.Renderer.Clear()
}
