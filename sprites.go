package gogame

import (
	"fmt"
	"sort"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"

	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	ratio                        = 32.0
	spriteToPhysicsRatio float64 = 1.0 / ratio // 1/32nd
	physicsToSpriteRatio float64 = ratio / 1.0 // 32 times
)

func CreateSprite(spriteId string, sprite *Sprite) error {

	if sprite == nil {
		return fmt.Errorf("Error: sprite pointer is nil")
	}

	if sprite.ImageResourceId != "" {
		err := sprite.SetImage(sprite.ImageResourceId)
		if err != nil {
			return err
		}
	}

	sprite.applyPhysics = false
	sprite.mass = 0.0

	spriteBank[spriteId] = sprite

	return nil
}

func (s *Sprite) EnablePhysics(mass float64) {
	s.mass = mass

	sizeOfBody := b2d.Vec2{float64(s.Width) * spriteToPhysicsRatio, float64(s.Height) * spriteToPhysicsRatio}

	body := b2d.Body{}
	body.Set(&sizeOfBody, mass)

	posOfBody := b2d.Vec2{float64(s.Pos.X) * spriteToPhysicsRatio, float64(s.Pos.Y) * spriteToPhysicsRatio}

	body.Position = posOfBody
	body.Rotation = s.Rotation * DegToRad

	s.applyPhysics = true

	s.body = &body

	game.world.AddBody(&body)

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

	return renderSpriteWithOffset(sprite, sdl.Point{0.0, 0.0})

}

func (s *Sprite) renderBox(centre b2d.Vec2, rotInRadians float64) {

	rotation := b2d.Mat22ByAngle(rotInRadians)

	half := b2d.Vec2{float64(s.Width / 2), float64(-s.Height / 2)}

	v1 := centre.Add(rotation.MulV(b2d.Vec2{-half.X, -half.Y}))
	v2 := centre.Add(rotation.MulV(b2d.Vec2{half.X, -half.Y}))
	v3 := centre.Add(rotation.MulV(b2d.Vec2{half.X, half.Y}))
	v4 := centre.Add(rotation.MulV(b2d.Vec2{-half.X, half.Y}))

	game.Renderer.DrawLine(int(v1.X), int(v1.Y), int(v2.X), int(v2.Y))
	game.Renderer.DrawLine(int(v2.X), int(v2.Y), int(v3.X), int(v3.Y))
	game.Renderer.DrawLine(int(v3.X), int(v3.Y), int(v4.X), int(v4.Y))
	game.Renderer.DrawLine(int(v4.X), int(v4.Y), int(v1.X), int(v1.Y))
}

func renderSpriteWithOffset(sprite *Sprite, offset sdl.Point) error {
	if sprite == nil {
		return fmt.Errorf("Error sprite pointer is nil")
	}

	if !sprite.Visible {
		// don't render it
		return nil
	}

	var pos sdl.Point
	var rotInRadians float64

	if sprite.applyPhysics {
		// use body co-ords for rendering
		game.Renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
		pos = sdl.Point{int32(sprite.body.Position.X * physicsToSpriteRatio), int32(sprite.body.Position.Y * physicsToSpriteRatio)}
		rotInRadians = sprite.body.Rotation
	} else {
		game.Renderer.SetDrawColor(0x00, 0x00, 0xff, 0xff)
		pos = sprite.Pos
		rotInRadians = sprite.Rotation * DegToRad
	}

	relativePos := sdl.Point{pos.X + offset.X, pos.Y + offset.Y}

	centre := b2d.Vec2{float64(relativePos.X + sprite.Width/2), float64(relativePos.Y + sprite.Height/2)}

	if game.RenderBoxes {
		// render outline box of sprite
		sprite.renderBox(centre, rotInRadians)
	}

	return renderRotatedTexture(sprite.texture, relativePos, sprite.Rotation, sprite.image.W, sprite.image.H, sprite.Width, sprite.Height)

}
