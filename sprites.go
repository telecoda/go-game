package gogame

import (
	"fmt"
	"sort"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func CreateSprite(spriteId string, sprite Sprite) error {

	if sprite.ImageResourceId != "" {
		err := sprite.SetImage(sprite.ImageResourceId)
		if err != nil {
			return err
		}
	}

	spriteBank[spriteId] = &sprite

	return nil
}

func CreateSpriteLayer(layerId int, pos sdl.Point) (*SpriteLayer, error) {

	layer, ok := spriteLayers[layerId]
	if !ok {
		// add new layer
		layer = newSpriteLayer(sdl.Point{0, 0})
		spriteLayers[layerId] = layer
		layer.Sprites = make(SpriteMap)
		return layer, nil
	} else {
		return nil, fmt.Errorf("Error sprite layer :%d already exists")
	}

}

func newSpriteLayer(pos sdl.Point) *SpriteLayer {

	layer := SpriteLayer{
		Pos:     pos,
		Visible: true,
		Sprites: make(SpriteMap),
	}

	return &layer
}

func (l *SpriteLayer) AddSpriteToLayer(spriteId string) error {
	// lookup sprite pointer
	sprite, err := GetSprite(spriteId)
	if err != nil {
		return err
	}

	// store pointer
	l.Sprites[spriteId] = sprite
	return nil
}

func RenderLayers() error {

	furthest := len(spriteLayers)
	if furthest == 0 {
		// no layers
		return fmt.Errorf("Error: no layers to render")
	}
	for l := furthest - 1; l >= 0; l-- {
		layer := spriteLayers[l]
		renderLayer(layer)
	}

	return nil
}

func renderLayer(layer *SpriteLayer) error {
	if !layer.Visible {
		// don't render
		return nil
	}

	ids := make([]string, len(layer.Sprites))
	i := 0

	// extract sprite id's
	for id, _ := range layer.Sprites {
		ids[i] = id
		i++
	}

	sort.Strings(ids)
	for _, id := range ids {
		err := renderSpriteWithOffset(layer.Sprites[id], layer.Pos)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetSprite(spriteId string) (*Sprite, error) {

	sprite, ok := spriteBank[spriteId]
	if !ok {
		return nil, fmt.Errorf("Warning: unknown sprite resource:%\n ", spriteId)
	}

	if sprite == nil {
		return nil, fmt.Errorf("Error: pointer for sprite:%s is nil\n ", spriteId)
	}

	return sprite, nil
}

func (s *Sprite) SetImage(resourceId string) error {

	image, texture, err := getImageResource(resourceId)
	if err != nil {
		return err
	}

	s.ImageResourceId = resourceId
	s.image = image
	s.texture = texture

	return nil
}

func RenderSprite(spriteId string) error {

	sprite, ok := spriteBank[spriteId]
	if !ok {
		return fmt.Errorf("Warning: unknown sprite resource:%\n ", spriteId)
	}

	return renderSprite(sprite)

}

func renderSprite(sprite *Sprite) error {
	if sprite == nil {
		return fmt.Errorf("Error sprite pointer is nil")
	}

	if !sprite.Visible {
		// don't render it
		return nil
	}

	return renderRotatedTexture(sprite.texture, sprite.Pos, sprite.Rotation, sprite.image.W, sprite.image.H, sprite.Width, sprite.Height)

}

func renderSpriteWithOffset(sprite *Sprite, offset sdl.Point) error {
	if sprite == nil {
		return fmt.Errorf("Error sprite pointer is nil")
	}

	if !sprite.Visible {
		// don't render it
		return nil
	}

	relativePos := sdl.Point{sprite.Pos.X + offset.X, sprite.Pos.Y + offset.Y}

	return renderRotatedTexture(sprite.texture, relativePos, sprite.Rotation, sprite.image.W, sprite.image.H, sprite.Width, sprite.Height)

}
