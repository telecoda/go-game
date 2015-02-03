package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (g *Game) RenderTexture(resourceId string, pos sdl.Point, textureWidth, textureHeight int32) error {

	image, texture, err := g.GetImageResource(resourceId)
	if err != nil {
		return err
	}

	width := image.W
	height := image.H

	src := sdl.Rect{0, 0, width, height}

	dst := sdl.Rect{pos.X, pos.Y, textureWidth, textureHeight}
	g.Renderer.Copy(texture, &src, &dst)

	return nil
}

func (g *Game) RenderRotatedTexture(resourceId string, pos sdl.Point, rotation float64, textureWidth, textureHeight int32) error {

	image, texture, err := g.GetImageResource(resourceId)
	if err != nil {
		return err
	}

	width := image.W
	height := image.H

	return g.renderRotatedTexture(texture, pos, rotation, width, height, textureWidth, textureHeight)

}

func (g *Game) renderRotatedTexture(texture *sdl.Texture, pos sdl.Point, rotation float64, width, height, textureWidth, textureHeight int32) error {

	if texture == nil {
		return fmt.Errorf("Error: texture is nil")
	}
	src := sdl.Rect{0, 0, width, height}
	centre := sdl.Point{width / 2, height / 2}
	dst := sdl.Rect{pos.X, pos.Y, textureWidth, textureHeight}
	g.Renderer.CopyEx(texture, &src, &dst, rotation, &centre, sdl.FLIP_NONE)

	return nil
}
