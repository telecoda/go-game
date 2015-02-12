package main

// author: @telecoda

/* demo - to show the various features of the go-game library
 */

import (
	"fmt"

	gogame "github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

var gameWidth = 1024
var gameHeight = 800

var currentDemoScreenId int

var demoScreens map[int]*DemoScreen
var demoScreen *DemoScreen

var assetHandler gogame.AssetManager
var renderController gogame.RenderController
var eventHandler gogame.EventHandler

func main() {

	initDemoScreens()

	var err error

	// init assetHandler and renderController
	assetHandler, renderController, eventHandler, err = gogame.NewGame("go-game demo", gameWidth, gameHeight, nil, demoEventReceiver)
	if err != nil {
		fmt.Printf("Error creating game:%s. Program exit.\n", err)
		return
	}
	defer gogame.Destroy()

	// start with demo 1
	currentDemoScreenId = 0
	demoScreen = demoScreens[currentDemoScreenId]
	err = demoScreen.activate()

	if err != nil {
		fmt.Printf("Error activating screen:%s. Program exit.\n", err)
		return
	}
	// start main event loop
	gogame.EventLoop()

}

// receive events from game
func demoEventReceiver(e interface{}) {
	switch t := e.(type) {
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_SPACE:
			err := nextDemo()
			if err != nil {
				fmt.Printf("Error switching demos :%s", err)
			}
		}
	}
}

func nextDemo() error {
	nextDemoScreenId := currentDemoScreenId + 1

	if nextDemoScreenId > len(demoScreens)-1 {
		currentDemoScreenId = 0
	} else {
		currentDemoScreenId = nextDemoScreenId
	}

	previousDemoScreen := demoScreen
	demoScreen = demoScreens[currentDemoScreenId]
	err := demoScreen.activate()
	if err != nil {
		return err
	}

	err = previousDemoScreen.UnloadAssets()
	if err != nil {
		return err
	}

	gogame.ReportMemoryUsage()

	return nil

}

// make this screen the current one being rendered by the renderer
func (d DemoScreen) activate() error {
	err := demoScreen.InitAssets()
	if err != nil {
		return err
	}
	renderController.SetCallback(d.RenderCallback)

	return nil
}

func initDemoScreens() {

	demoScreens = make(map[int]*DemoScreen)

	demoScreens[0] = &DemoScreen{Id: 0, Title: "Title screen", Color: sdl.Color{R: 255, G: 255, B: 255, A: 255}, InitAssets: initDemo0Assets, UnloadAssets: unloadDemo0Assets, RenderCallback: demo0RenderCallback}
	demoScreens[1] = &DemoScreen{Id: 1, Title: "Text demo", Color: sdl.Color{R: 0, G: 0, B: 0, A: 255}, InitAssets: initDemo1Assets, UnloadAssets: unloadDemo1Assets, RenderCallback: demo1RenderCallback}
	demoScreens[2] = &DemoScreen{Id: 2, Title: "Credits screen", Color: sdl.Color{R: 128, G: 128, B: 128, A: 255}, InitAssets: initDemo2Assets, UnloadAssets: unloadDemo2Assets, RenderCallback: demo1RenderCallback}

}

// Single box
/*
func (app *App) GopherDemo() {
	app.World.Clear()

	var b1 b2d.Body

	// create floor (max mass)
	b1.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
	b1.Position = b2d.Vec2{0.0, -0.7 * b1.Width.Y}
	//b1.Rotation = 12.0 * DegToRad

	app.World.AddBody(&b1)

	// create gopher
	app.AddGopher(b2d.Vec2{0.0, 4.0})
}

func (app *App) AddGopher(pos b2d.Vec2) {
	var body b2d.Body

	body.Set(&b2d.Vec2{1.0, 1.0}, 200.0)
	body.Position = pos
	body.Rotation = 0.0
	app.World.AddBody(&body)

}
*/

/*
func (app *App) RenderBody(b *b2d.Body) {
	app.Renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)

	R := b2d.Mat22ByAngle(b.Rotation)
	x := b.Position
	h := b2d.MulSV(0.5, b.Width)

	o := b2d.Vec2{400, 400}
	S := b2d.Mat22{b2d.Vec2{spriteWidth, 0.0}, b2d.Vec2{0.0, -spriteHeight}}

	v1 := o.Add(S.MulV(x.Add(R.MulV(b2d.Vec2{-h.X, -h.Y}))))
	v2 := o.Add(S.MulV(x.Add(R.MulV(b2d.Vec2{h.X, -h.Y}))))
	v3 := o.Add(S.MulV(x.Add(R.MulV(b2d.Vec2{h.X, h.Y}))))
	v4 := o.Add(S.MulV(x.Add(R.MulV(b2d.Vec2{-h.X, h.Y}))))

	app.Renderer.DrawLine(int(v1.X), int(v1.Y), int(v2.X), int(v2.Y))
	app.Renderer.DrawLine(int(v2.X), int(v2.Y), int(v3.X), int(v3.Y))
	app.Renderer.DrawLine(int(v3.X), int(v3.Y), int(v4.X), int(v4.Y))
	app.Renderer.DrawLine(int(v4.X), int(v4.Y), int(v1.X), int(v1.Y))
}

func (app *App) RenderTexturedBody(b *b2d.Body, t *sdl.Texture) {
	app.Renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)

	R := b2d.Mat22ByAngle(b.Rotation)
	x := b.Position
	h := b2d.MulSV(0.5, b.Width)

	o := b2d.Vec2{400, 400}
	S := b2d.Mat22{b2d.Vec2{spriteWidth, 0.0}, b2d.Vec2{0.0, -spriteHeight}}

	v1 := o.Add(S.MulV(x.Add(R.MulV(b2d.Vec2{-h.X, -h.Y}))))

	src := sdl.Rect{0, 0, spriteWidth, spriteHeight}
	// bottom left
	centre := sdl.Point{0, spriteHeight}

	dst := sdl.Rect{int32(v1.X), int32(v1.Y) - spriteHeight, spriteWidth, spriteHeight}
	app.Renderer.CopyEx(t, &src, &dst, -b.Rotation*RadToDeg, &centre, sdl.FLIP_NONE)

}

func (app *App) RenderText(text string, pos b2d.Vec2) {

	textColour := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	textSurface := font.RenderText_Solid(text, textColour)
	texture, err := app.Renderer.CreateTextureFromSurface(textSurface)

	if err != nil {
		fmt.Printf("Error rendering text:%s\n", err)
	}

	src := sdl.Rect{0, 0, textSurface.W, textSurface.H}
	dst := sdl.Rect{int32(pos.X), int32(pos.Y), textSurface.W, textSurface.H}
	app.Renderer.Copy(texture, &src, &dst)

	textSurface.Free()
	texture.Destroy()

}

func (app *App) RenderJoint(j *b2d.Joint) {
	app.Renderer.SetDrawColor(0x80, 0x80, 0x80, 0x80)

	b1 := j.Body1
	b2 := j.Body2

	R1 := b2d.Mat22ByAngle(b1.Rotation)
	R2 := b2d.Mat22ByAngle(b2.Rotation)

	x1 := b1.Position
	p1 := x1.Add(R1.MulV(j.LocalAnchor1))

	x2 := b2.Position
	p2 := x2.Add(R2.MulV(j.LocalAnchor2))

	o := b2d.Vec2{400, 400}
	S := b2d.Mat22{b2d.Vec2{20.0, 0.0}, b2d.Vec2{0.0, -20.0}}

	x1 = o.Add(S.MulV(x1))
	p1 = o.Add(S.MulV(p1))
	x2 = o.Add(S.MulV(x2))
	p2 = o.Add(S.MulV(p2))

	app.Renderer.DrawLine(int(x1.X), int(x1.Y), int(p1.X), int(p1.Y))
	app.Renderer.DrawLine(int(x2.X), int(x2.Y), int(p2.X), int(p2.Y))
}
*/
/*
func calcFps() float64 {
	now := time.Now()

	diff := now.Sub(lastFrame)
	lastFrame = now

	oneSecond := time.Duration(1 * time.Second)

	return float64(oneSecond.Nanoseconds()) / float64(diff.Nanoseconds())
}*/
