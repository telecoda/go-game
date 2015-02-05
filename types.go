package gogame

import (
	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type FontResource struct {
	Id       string
	FilePath string
	Size     int
	loaded   bool
	font     *ttf.Font
}

type ImageResource struct {
	Id       string
	FilePath string
	loaded   bool
	image    *sdl.Surface
	texture  *sdl.Texture
}

type AudioResource struct {
	Id       string
	FilePath string
	loaded   bool
}

type audioResourceMap map[string]*AudioResource
type fontResourceMap map[string]*FontResource
type imageResourceMap map[string]ImageResource

type RenderFunction func()

type Game struct {
	Window          *sdl.Window
	Renderer        *sdl.Renderer
	renderCallback  RenderFunction
	quit            bool
	world           *b2d.World
	width           int
	height          int
	RenderBoxes     bool
	FramesPerSecond float64
}

type Sprite struct {
	Id              string
	Pos             sdl.Point
	Width           int32
	Height          int32
	Rotation        float64
	Visible         bool
	ImageResourceId string
	image           *sdl.Surface
	texture         *sdl.Texture

	applyPhysics bool
	mass         float64
	body         *b2d.Body
}

type SpriteMap map[string]*Sprite

type SpriteLayer struct {
	Pos     sdl.Point
	Visible bool
	Sprites SpriteMap
}
type SpriteLayers map[int]*SpriteLayer
