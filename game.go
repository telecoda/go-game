package game

import (
	"runtime"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
)

const (
	timeStep = 1.0 / 60
)

func NewGame(winTitle string, winWidth, winHeight int) *Game {
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

	return &Game{
		Window:    window,
		Renderer:  renderer,
		resources: make(ResourceMap, 0),
	}
}

func NewPhysicsGame(winTitle string, winWidth, winHeight int) *PhysicsGame {

	game := NewGame(winTitle, winWidth, winHeight)
	if game == nil {
		return nil
	}

	gravity := b2d.Vec2{0.0, -10.0}
	iterations := 10
	world := b2d.NewWorld(gravity, iterations)

	phy := PhysicsGame{}
	phy.Window = game.Window
	phy.Renderer = game.Renderer
	phy.resources = game.resources
	phy.World = world

	return &phy
}

func (game *Game) Destroy() {
	// free image resources
	//gopherImage.Free()
	//gopherTexture.Destroy()

	// free font resources
	//font.Close()

	// free SDL resources
	game.Renderer.Destroy()
	game.Window.Destroy()
}

func (game *Game) EventLoop() {
	t1 := sdl.GetTicks()

	for {
		game.DoEvents()

		t2 := sdl.GetTicks()
		game.OnUpdate(t2 - t1)
		game.OnRender()
		t1 = t2

		sdl.Delay(16)
		game.Present()

		if game.quit {
			break
		}
	}
}

func (pGame *PhysicsGame) EventLoop() {
	t1 := sdl.GetTicks()

	for {
		pGame.DoEvents()

		t2 := sdl.GetTicks()
		pGame.OnUpdate(t2 - t1)
		pGame.OnRender()
		t1 = t2

		sdl.Delay(16)
		pGame.Present()

		if pGame.quit {
			break
		}
	}
}

func (game *Game) DoEvents() {
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}
		game.ProcessEvent(e)
	}
}

func (game *Game) ProcessEvent(e interface{}) {

	switch t := e.(type) {
	case *sdl.QuitEvent:
		game.quit = true
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_SPACE:
			//app.AddGopher(b2d.Vec2{0.0, 4.0})
		}
	}
}

func (game *Game) OnUpdate(ms uint32) {
	// do nothing
}

func (pGame *PhysicsGame) OnUpdate(ms uint32) {
	pGame.World.Step(timeStep)
}

func (game *Game) OnRender() {
	game.Renderer.SetDrawColor(0xee, 0xee, 0xee, 0x00)
	game.Renderer.Clear()

}

func (pGame *PhysicsGame) OnRender() {
	pGame.Renderer.SetDrawColor(0xff, 0x00, 0x00, 0x00)
	pGame.Renderer.Clear()

	/*
		for _, b := range app.World.Bodies {
			app.RenderTexturedBody(b, gopherTexture)
			app.RenderBody(b)
		}
		for _, j := range app.World.Joints {
			app.RenderJoint(j)
		}

		text := fmt.Sprintf("Gopher count:%d", len(app.World.Bodies)-1)
		app.RenderText(text, b2d.Vec2{20, 20})

		frameRate := calcFps()

		fpsText := fmt.Sprintf("Frame rate:%3.0f", frameRate)
		app.RenderText(fpsText, b2d.Vec2{20, 60})
	*/
}

func (game *Game) Present() {
	game.Renderer.Present()
}

func init() {
	// http://www.oki-osk.jp/esc/golang/cgo-osx.html#3
	runtime.LockOSThread()
}
