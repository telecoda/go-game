package gogame

import (
	"fmt"
	"sort"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func (r renderController) CreateSpriteLayer(layerId int, pos sdl.Point) (*SpriteLayer, error) {

	layer, ok := r.spriteLayers[layerId]
	if !ok {
		// add new layer
		layer = newSpriteLayer(sdl.Point{0, 0})
		r.spriteLayers[layerId] = layer
		layer.Sprites = make(spriteMap)
		return layer, nil
	} else {
		return nil, fmt.Errorf("Error sprite layer :%d already exists")
	}

}

func newSpriteLayer(pos sdl.Point) *SpriteLayer {

	layer := SpriteLayer{
		Pos:     pos,
		Visible: true,
		Sprites: make(spriteMap),
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
func (r renderController) RenderLayers() error {

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
		err := renderSpriteWithOffset(l.Sprites[id], l.Pos)
		if err != nil {
			return err
		}
	}

	return nil
}