package gogame

import (
	"fmt"
	"math"
	"runtime"
	"time"

	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
	mix "github.com/veandco/go-sdl2/sdl_mixer"
)

const (
	timeStep       = 1.0 / 60
	RadToDeg       = 180 / math.Pi
	DegToRad       = math.Pi / 180
	SYSTEM_FONT_ID = "system-font-id"
	ImmovableMass  = math.MaxFloat64
)

var lastFrame = time.Now()

var gameAssets *assets

var rendCont *renderController
var eventHnd *eventHandler
var audioPlyr *audioPlayer

var FramesPerSecond = 0.0

func init() {
	// init global library assets

	gameAssets = &assets{}
	gameAssets.Initialize()
}

func NewGame(winTitle string, winWidth, winHeight int, renderCallback RenderFunction, eventCallback EventReceiverFunction) (AssetManager, RenderController, AudioPlayer, EventHandler, error) {
	window, _ := sdl.CreateWindow(
		winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)

	if winTitle == "" {
		return nil, nil, nil, nil, fmt.Errorf("Error: window must have a title.")
	}

	if winWidth < 1 {
		return nil, nil, nil, nil, fmt.Errorf("Error: window width must be greater than 0.")
	}

	if winHeight < 1 {
		return nil, nil, nil, nil, fmt.Errorf("Error: window height must be greater than 0.")
	}

	if window == nil {
		return nil, nil, nil, nil, fmt.Errorf("Error: window not created")
	}

	//window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)

	// try acceleration first
	renderer, _ := sdl.CreateRenderer(window, -2, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		// revert to software
		renderer, _ := sdl.CreateRenderer(window, -2, sdl.RENDERER_SOFTWARE)
		if renderer == nil {
			return nil, nil, nil, nil, fmt.Errorf("Error: rendered not created")
		}
	}

	// init audio
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("Failed to init SDL audio. Error: %s\n", err)
	}
	err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("Failed to open audio. Error: \n", err)
	}

	gravity := b2d.Vec2{0.0, 10.0}
	iterations := 10
	world := b2d.NewWorld(gravity, iterations)
	world.Clear()

	// destroy old assets first
	gameAssets.Destroy()
	err = gameAssets.Initialize()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	rendCont = &renderController{
		Window:          window,
		Renderer:        renderer,
		renderCallback:  renderCallback,
		world:           world,
		width:           winWidth,
		height:          winHeight,
		spriteLayers:    make(SpriteLayers),
		RenderDebugInfo: false,
	}

	audioPlyr = &audioPlayer{}

	eventHnd = &eventHandler{
		eventCallback: eventCallback,
	}

	return gameAssets, rendCont, audioPlyr, eventHnd, nil
}

func (r *renderController) SetCallback(callback RenderFunction) {
	r.renderCallback = callback
}

func (r *renderController) SetDebugInfo(enabled bool) {
	r.RenderDebugInfo = enabled
}

// allow low level access to SDL renderer
func (r *renderController) GetRenderer() *sdl.Renderer {
	return r.Renderer
}

// remove all physics bodies
func (r *renderController) ClearWorld() {
	r.world.Clear()
}

func (e *eventHandler) SetCallback(callback EventReceiverFunction) {
	e.eventCallback = callback
}

func Destroy() {
	// free assets
	gameAssets.Destroy()

	// free SDL assets
	rendCont.Renderer.Destroy()
	rendCont.Window.Destroy()
}

func EventLoop() {
	t1 := sdl.GetTicks()

	for {
		doEvents()

		t2 := sdl.GetTicks()
		onUpdate(t2 - t1)
		onRender()
		t1 = t2

		sdl.Delay(1000 / 120)
		present()

		if rendCont.quit {
			break
		}
	}
}

func doEvents() {
	for {
		e := sdl.PollEvent()
		if e == nil {
			break
		}
		processEvent(e)
	}
}

func processEvent(e interface{}) {

	if eventHnd.eventCallback != nil {
		eventHnd.eventCallback(e)
	}

	switch t := e.(type) {
	case *sdl.QuitEvent:
		rendCont.quit = true
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_ESCAPE:
			rendCont.quit = true
		}
	}
}

func onUpdate(ms uint32) {
	rendCont.world.Step(timeStep)
}

func onRender() {
	rendCont.Renderer.SetDrawColor(0xe0, 0xff, 0xff, 0x00)
	rendCont.Renderer.Clear()
	FramesPerSecond = calcFps()

	if rendCont.renderCallback != nil {

		rendCont.renderCallback()
		if rendCont.RenderDebugInfo {
			//rendCont.renderJoints()
		}

	}

}

func (r *renderController) SetDefaultFont(fontId string) error {
	if fontId == "" {
		return fmt.Errorf("Error: default fontId cannot be empty.")
	}

	_, err := getFont(fontId)
	if err != nil {
		return err
	}

	rendCont.defaultFontId = fontId
	return nil
}

func present() {
	rendCont.Renderer.Present()
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
