package gogame

import (
	"math"
	"runtime"
	"time"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	timeStep = 1.0 / 60
	RadToDeg = 180 / math.Pi
	DegToRad = math.Pi / 180
)

var lastFrame = time.Now()

var audioResources audioResourceMap
var fontResources fontResourceMap
var imageResources imageResourceMap
var spriteBank SpriteMap // Sprite bank contains ALL sprites
var spriteLayers SpriteLayers

var game Game

var FramesPerSecond = 0.0

func init() {
	// init global library resources
	audioResources = make(audioResourceMap)
	fontResources = make(fontResourceMap)
	imageResources = make(imageResourceMap)
	spriteBank = make(SpriteMap)
	spriteLayers = make(SpriteLayers)
}

func NewGame(winTitle string, winWidth, winHeight int, renderCallback RenderFunction) error {
	window, _ := sdl.CreateWindow(
		winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
	if window == nil {
		return nil
	}

	renderer, _ := sdl.CreateRenderer(window, -2, sdl.RENDERER_SOFTWARE)
	if renderer == nil {
		return nil
	}

	gravity := b2d.Vec2{0.0, 10.0}
	iterations := 10
	world := b2d.NewWorld(gravity, iterations)
	world.Clear()

	// destroy old resources first
	audioResources.Destroy()
	audioResources = make(audioResourceMap)
	imageResources.Destroy()
	imageResources = make(imageResourceMap)
	fontResources.Destroy()
	fontResources = make(fontResourceMap)

	spriteBank = make(SpriteMap)
	spriteLayers = make(SpriteLayers)

	game = Game{
		Window:         window,
		Renderer:       renderer,
		renderCallback: renderCallback,
		world:          world,
		width:          winWidth,
		height:         winHeight,
		RenderBoxes:    true,
	}

	return nil
}

func Destroy() {
	// free image resources
	imageResources.Destroy()

	// free font resources
	fontResources.Destroy()

	// free audio resources
	audioResources.Destroy()

	// free SDL resources
	game.Renderer.Destroy()
	game.Window.Destroy()
}

func EventLoop() {
	t1 := sdl.GetTicks()

	for {
		DoEvents()

		t2 := sdl.GetTicks()
		OnUpdate(t2 - t1)
		OnRender()
		t1 = t2

		sdl.Delay(16)
		Present()

		if game.quit {
			break
		}
	}
}

func DoEvents() {
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}
		ProcessEvent(e)
	}
}

func ProcessEvent(e interface{}) {

	switch t := e.(type) {
	case *sdl.QuitEvent:
		game.quit = true
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			game.quit = true
		}
	}
}

func OnUpdate(ms uint32) {
	game.world.Step(timeStep)
}

func OnRender() {
	game.Renderer.SetDrawColor(0xe0, 0xff, 0xff, 0x00)
	game.Renderer.Clear()
	FramesPerSecond = calcFps()
	game.renderCallback()

}

func Present() {
	game.Renderer.Present()
}

func init() {
	// http://www.oki-osk.jp/esc/golang/cgo-osx.html#3
	runtime.LockOSThread()
}

func calcFps() float64 {
	now := time.Now()

	diff := now.Sub(lastFrame)
	lastFrame = now

	oneSecond := time.Duration(1 * time.Second)

	return float64(oneSecond.Nanoseconds()) / float64(diff.Nanoseconds())
}
