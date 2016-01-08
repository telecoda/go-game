package gogame

import (
	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
	mix "github.com/veandco/go-sdl2/sdl_mixer"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type Engine struct {
	AssetManager
	renderer     *renderer
	audioPlayer  *audioPlayer
	eventHandler EventHandler
}

type BaseAsset struct {
	Id       string
	FilePath string
	loaded   bool
}

type FontAsset struct {
	BaseAsset
	Size int
	font *ttf.Font
}

type ImageAsset struct {
	BaseAsset
	image   *sdl.Surface
	texture *sdl.Texture
}

type AudioAsset struct {
	BaseAsset
	chunk *mix.Chunk
}

type assets struct {
	audioAssets audioAssetMap
	fontAssets  fontAssetMap
	imageAssets imageAssetMap
	spriteBank  spriteMap
}

type audioAssetMap map[string]*AudioAsset
type fontAssetMap map[string]*FontAsset
type imageAssetMap map[string]*ImageAsset
type spriteMap map[string]*Sprite

type RenderFunction func()

type EventReceiverFunction func(e interface{})

type renderer struct {
	Window          *sdl.Window
	Renderer        *sdl.Renderer
	renderCallback  RenderFunction
	quit            bool
	world           *b2d.World
	width           int
	height          int
	spriteLayers    SpriteLayers
	RenderDebugInfo bool
	FramesPerSecond float64
	defaultFontId   string
}

type audioPlayer struct {
}

type eventHandler struct {
	eventCallback EventReceiverFunction
}

type Sprite struct {
	Id           string
	Pos          sdl.Point
	Width        int32
	Height       int32
	Rotation     float64
	Visible      bool
	ImageAssetId string
	image        *sdl.Surface
	texture      *sdl.Texture

	applyPhysics bool
	mass         float64
	body         *b2d.Body
}

type SpriteLayer struct {
	AbsPos  sdl.Point
	Offset  sdl.Point
	Visible bool
	Wrap    bool
	Width   int32
	Height  int32
	Sprites spriteMap
}
type SpriteLayers map[int]*SpriteLayer
