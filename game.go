package gogame

import (
	"runtime"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	timeStep = 1.0 / 60
)

func NewGame(winTitle string, winWidth, winHeight int, renderCallback RenderFunction) *Game {
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

	gravity := b2d.Vec2{0.0, -10.0}
	iterations := 10
	world := b2d.NewWorld(gravity, iterations)

	audio := make(audioResourceMap)
	images := make(imageResourceMap)
	fonts := make(fontResourceMap)
	sprites := make(SpriteMap)

	return &Game{
		Window:         window,
		Renderer:       renderer,
		SpriteBank:     sprites,
		renderCallback: renderCallback,
		audioResources: audio,
		fontResources:  fonts,
		imageResources: images,
		world:          world,
		width:          winWidth,
		height:         winHeight,
	}
}

func (g *Game) Destroy() {
	// free image resources
	g.imageResources.Destroy()

	// free font resources
	g.fontResources.Destroy()

	// free audio resources
	g.audioResources.Destroy()

	// free SDL resources
	g.Renderer.Destroy()
	g.Window.Destroy()
}

func (g *Game) EventLoop() {
	t1 := sdl.GetTicks()

	for {
		g.DoEvents()

		t2 := sdl.GetTicks()
		g.OnUpdate(t2 - t1)
		g.OnRender()
		t1 = t2

		sdl.Delay(16)
		g.Present()

		if g.quit {
			break
		}
	}
}

func (g *Game) DoEvents() {
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}
		g.ProcessEvent(e)
	}
}

func (g *Game) ProcessEvent(e interface{}) {

	switch t := e.(type) {
	case *sdl.QuitEvent:
		g.quit = true
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			g.quit = true
		}
	}
}

func (g *Game) OnUpdate(ms uint32) {
	g.world.Step(timeStep)
}

func (g *Game) OnRender() {
	g.renderCallback()

}

func (g *Game) Present() {
	g.Renderer.Present()
}

func init() {
	// http://www.oki-osk.jp/esc/golang/cgo-osx.html#3
	runtime.LockOSThread()
}
