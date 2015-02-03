package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (g *Game) AddSprite(spriteId string, sprite Sprite) error {

	g.SpriteBank[spriteId] = &sprite

	return nil
}

func (g *Game) GetSprite(spriteId string) (*Sprite, error) {

	sprite, ok := g.SpriteBank[spriteId]
	if !ok {
		return nil, fmt.Errorf("Warning: unknown sprite resource:%\n ", spriteId)
	}

	return sprite, nil
}

func (s *Sprite) SetImage(image *sdl.Surface, texture *sdl.Texture) error {

	s.image = image
	s.texture = texture

	return nil
}

func (g *Game) RenderSprite(spriteId string) error {

	sprite, ok := g.SpriteBank[spriteId]
	if !ok {
		return fmt.Errorf("Warning: unknown sprite resource:%\n ", spriteId)
	}

	if sprite == nil {
		return nil
	}

	if !sprite.Visible {
		// don't render it
		return nil
	}
	fmt.Printf("Rendering sprite:%v\n", sprite)

	return g.renderRotatedTexture(sprite.texture, sprite.Pos, sprite.Rotation, sprite.image.W, sprite.image.H, sprite.Width, sprite.Height)

}
