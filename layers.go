package gogame

import (
	"fmt"
	"sort"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func (r renderer) CreateSpriteLayer(layerId int, pos sdl.Point) (*SpriteLayer, error) {

	layer, ok := r.spriteLayers[layerId]
	if !ok {
		// add new layer
		layer = newSpriteLayer(sdl.Point{0, 0})
		r.spriteLayers[layerId] = layer
		layer.Sprites = make(spriteMap)
		return layer, nil
	} else {
		return nil, fmt.Errorf("Error sprite layer :%d already exists\n", layerId)
	}

}

func (r *renderer) DestroySpriteLayer(layerId int) error {

	_, ok := r.spriteLayers[layerId]
	if ok {
		// destroy
		delete(r.spriteLayers, layerId)
		fmt.Printf("Destroying sprite layer :%d\n", layerId)
		return nil
	} else {
		return fmt.Errorf("Error sprite layer :%d does not exist\n", layerId)
	}
}

func newSpriteLayer(offset sdl.Point) *SpriteLayer {

	layer := SpriteLayer{
		Offset:  offset,
		Visible: true,
		Sprites: make(spriteMap),
		Width:   int32(rendCont.width),
		Height:  int32(rendCont.height),
	}

	return &layer
}

func (l *SpriteLayer) AddSpriteToLayer(spriteId string) error {
	// lookup sprite pointer
	sprite, err := gameAssets.GetSprite(spriteId)
	if err != nil {
		return err
	}

	// store pointer
	l.Sprites[spriteId] = sprite
	return nil
}

// renders layers
func (r renderer) RenderLayers() error {

	furthest := len(r.spriteLayers)
	if furthest == 0 {
		// no layers
		return fmt.Errorf("Error: no layers to render")
	}
	for l := furthest - 1; l >= 0; l-- {
		layer := r.spriteLayers[l]
		layer.render()
	}

	return nil
}

func (l *SpriteLayer) render() error {
	if !l.Visible {
		// don't render
		return nil
	}

	ids := make([]string, len(l.Sprites))
	i := 0

	// extract sprite id's
	for id, _ := range l.Sprites {
		ids[i] = id
		i++
	}

	sort.Strings(ids)
	for _, id := range ids {

		sprite := l.Sprites[id]
		offset := l.Offset
		if l.Wrap {
			// check if sprite offscreen
			offsetX := sprite.Pos.X + l.Offset.X
			offsetY := sprite.Pos.Y + l.Offset.Y

			wrappedX := offsetX % int32(l.Width)
			wrappedY := offsetY % int32(l.Height)

			offset = sdl.Point{wrappedX - sprite.Pos.X, wrappedY - sprite.Pos.Y}
		}

		pos := sdl.Point{l.AbsPos.X + offset.X, l.AbsPos.Y + offset.Y}

		err := renderSpriteWithOffset(l.Sprites[id], pos)
		if err != nil {
			return err
		}

	}

	return nil
}
