package gogame

import (
	"fmt"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"

	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	ratio                        = 32.0
	spriteToPhysicsRatio float64 = 1.0 / ratio // 1/32nd
	physicsToSpriteRatio float64 = ratio / 1.0 // 32 times
)

func (a *assets) AddSprite(sprite *Sprite) error {

	if sprite == nil {
		return fmt.Errorf("Error: sprite pointer is nil")
	}

	if sprite.Id == "" {
		return fmt.Errorf("Error: sprite.Id is empty")
	}

	if sprite.ImageAssetId != "" {
		err := sprite.SetImage(sprite.ImageAssetId)
		if err != nil {
			return err
		}
	}

	sprite.applyPhysics = false
	sprite.mass = 0.0

	a.spriteBank[sprite.Id] = sprite

	return nil
}

func (a *assets) GetSprite(spriteId string) (*Sprite, error) {

	sprite, ok := a.spriteBank[spriteId]
	if !ok {
		return nil, fmt.Errorf("Warning: unknown sprite asset:%s\n ", spriteId)
	}

	if sprite == nil {
		return nil, fmt.Errorf("Error: pointer for sprite:%s is nil\n ", spriteId)
	}

	return sprite, nil
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

	rendCont.world.AddBody(&body)

}

func (sprite1 *Sprite) JoinTo(sprite2 *Sprite, atPoint sdl.Point) error {

	// both sprites MUST have physics enabled
	if !sprite1.applyPhysics {
		return fmt.Errorf("Error sprite:%s does not have physics enabled", sprite1.Id)
	}

	if !sprite2.applyPhysics {
		return fmt.Errorf("Error sprite:%s does not have physics enabled", sprite2.Id)
	}

	joint := b2d.Joint{}

	jointPos := &b2d.Vec2{float64(atPoint.X+sprite1.Pos.X) * spriteToPhysicsRatio, float64(atPoint.Y+sprite1.Pos.Y) * spriteToPhysicsRatio}
	fmt.Printf("atPoint:%v\n", atPoint)
	fmt.Printf("joint pos:%v\n", jointPos)

	joint.Set(sprite1.body, sprite2.body, jointPos)

	fmt.Printf("sprite1 body:%v\n", sprite1.body)
	fmt.Printf("sprite2 body:%v\n", sprite2.body)
	fmt.Printf("Joint details:%v\n", joint)

	rendCont.world.AddJoint(&joint)

	return nil
}

func (s *Sprite) SetImage(assetId string) error {

	image, texture, err := getImage(assetId)
	if err != nil {
		return err
	}

	s.ImageAssetId = assetId
	s.image = image
	s.texture = texture

	return nil
}

func (r renderer) RenderSprite(spriteId string) error {

	sprite, ok := gameAssets.spriteBank[spriteId]
	if !ok {
		return fmt.Errorf("Warning: unknown sprite asset:%s\n ", spriteId)
	}

	return sprite.render()

}

func (s *Sprite) render() error {
	if s == nil {
		return fmt.Errorf("Error sprite pointer is nil")
	}

	if !s.Visible {
		// don't render it
		return nil
	}

	return renderSpriteWithOffset(s, sdl.Point{0.0, 0.0})

}

func (r renderer) renderJoints() {

	rendCont.Renderer.SetDrawColor(0x00, 0xff, 0x00, 0xff)

	for _, joint := range rendCont.world.Joints {

		joint1X := (joint.LocalAnchor1.X + joint.Body1.Position.X) * physicsToSpriteRatio
		joint1Y := (joint.LocalAnchor1.Y + joint.Body1.Position.Y) * physicsToSpriteRatio

		//TODO - adjust rendering to compensate for rotations

		rect := sdl.Rect{int32(joint1X - 1), int32(joint1Y - 1), 3, 3}
		rendCont.Renderer.FillRect(&rect)

	}

}

func (s *Sprite) renderBox(centre b2d.Vec2, rotInRadians float64) {

	rotation := b2d.Mat22ByAngle(rotInRadians)

	half := b2d.Vec2{float64(s.Width / 2), float64(-s.Height / 2)}

	v1 := centre.Add(rotation.MulV(b2d.Vec2{-half.X, -half.Y}))
	v2 := centre.Add(rotation.MulV(b2d.Vec2{half.X, -half.Y}))
	v3 := centre.Add(rotation.MulV(b2d.Vec2{half.X, half.Y}))
	v4 := centre.Add(rotation.MulV(b2d.Vec2{-half.X, half.Y}))

	rendCont.Renderer.DrawLine(int(v1.X), int(v1.Y), int(v2.X), int(v2.Y))
	rendCont.Renderer.DrawLine(int(v2.X), int(v2.Y), int(v3.X), int(v3.Y))
	rendCont.Renderer.DrawLine(int(v3.X), int(v3.Y), int(v4.X), int(v4.Y))
	rendCont.Renderer.DrawLine(int(v4.X), int(v4.Y), int(v1.X), int(v1.Y))

	// render centre point
	rendCont.Renderer.SetDrawColor(255, 0, 0, 255)
	rect := sdl.Rect{int32(centre.X - 1), int32(centre.Y - 1), 3, 3}
	rendCont.Renderer.FillRect(&rect)

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
	var rotInDegrees float64

	if sprite.applyPhysics {
		// use body co-ords for rendering
		rendCont.Renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
		pos = sdl.Point{int32(sprite.body.Position.X * physicsToSpriteRatio), int32(sprite.body.Position.Y * physicsToSpriteRatio)}
		rotInRadians = sprite.body.Rotation
		rotInDegrees = sprite.body.Rotation * RadToDeg
	} else {
		rendCont.Renderer.SetDrawColor(0x00, 0x00, 0xff, 0xff)
		pos = sprite.Pos
		rotInRadians = sprite.Rotation * DegToRad
		rotInDegrees = sprite.Rotation
	}

	relativePos := sdl.Point{pos.X + offset.X, pos.Y + offset.Y}

	// offset from middle of sprite
	relativePos.X -= sprite.Width / 2
	relativePos.Y -= sprite.Height / 2

	centre := b2d.Vec2{float64(relativePos.X + sprite.Width/2), float64(relativePos.Y + sprite.Height/2)}

	err := renderRotatedTexture(sprite.texture, relativePos, rotInDegrees, sprite.image.W, sprite.image.H, sprite.Width, sprite.Height)

	if err != nil {
		return err
	}

	if rendCont.RenderDebugInfo {
		// render outline box of sprite
		sprite.renderBox(centre, rotInRadians)

		// only render if default font set
		xPos := int32(centre.X) - sprite.Width/2
		yPos := int32(centre.Y) + sprite.Height/2
		debugTextPos := sdl.Point{xPos, yPos}
		textColour := sdl.Color{R: 0, G: 0, B: 0, A: 255}
		spriteIdText := fmt.Sprintf("ID:%s", sprite.Id)
		rendCont.RenderText(rendCont.defaultFontId, spriteIdText, debugTextPos, 0.0, textColour, TOP, LEFT)
		spritePosText := fmt.Sprintf("X:%d Y:%d", sprite.Pos.X, sprite.Pos.Y)
		debugTextPos = sdl.Point{xPos, debugTextPos.Y + 10}
		rendCont.RenderText(rendCont.defaultFontId, spritePosText, debugTextPos, 0.0, textColour, TOP, LEFT)
		spriteSizeText := fmt.Sprintf("W:%d H:%d", sprite.Width, sprite.Height)
		debugTextPos = sdl.Point{xPos, debugTextPos.Y + 10}
		rendCont.RenderText(rendCont.defaultFontId, spriteSizeText, debugTextPos, 0.0, textColour, TOP, LEFT)
		spriteRotText := fmt.Sprintf("rot:%2.5f", sprite.Rotation)
		debugTextPos = sdl.Point{xPos, debugTextPos.Y + 10}
		rendCont.RenderText(rendCont.defaultFontId, spriteRotText, debugTextPos, 0.0, textColour, TOP, LEFT)
		spritePhysicsText := fmt.Sprintf("physics:%t", sprite.applyPhysics)
		debugTextPos = sdl.Point{xPos, debugTextPos.Y + 10}
		rendCont.RenderText(rendCont.defaultFontId, spritePhysicsText, debugTextPos, 0.0, textColour, TOP, LEFT)
		if sprite.applyPhysics {
			// also render physics info
			var spritePhysicsText string
			if sprite.body.Mass == ImmovableMass {
				spritePhysicsText = fmt.Sprintf("body x:%2.5f y:%2.5f mass:ImmovableMass", sprite.body.Position.X, sprite.body.Position.Y)
			} else {
				spritePhysicsText = fmt.Sprintf("body x:%2.5f y:%2.5f mass:%2.5f", sprite.body.Position.X, sprite.body.Position.Y, sprite.body.Mass)
			}
			debugTextPos = sdl.Point{xPos, debugTextPos.Y + 10}
			rendCont.RenderText(rendCont.defaultFontId, spritePhysicsText, debugTextPos, 0.0, textColour, TOP, LEFT)

		}

	}

	return nil

}
