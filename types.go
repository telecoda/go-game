package gogame

import (
	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type GameEngine interface {
	NewGame(winTitle string, winWidth, winHeight int, renderCallback RenderFunction) error
	Destroy()
	EventLoop()
}

type FontHandler interface {
	AddFontResource(resource FontResource) error // add a font
	getFontResource(resourceId string) (*ttf.Font, error)
}

type ImageHandler interface {
	AddImageResource(resource ImageResource) error // add an image
	getImageResource(resourceId string) (*sdl.Surface, *sdl.Texture, error)
}

type SpriteHandler interface {
	AddSprite(spriteId string, sprite *Sprite) error
	GetSprite(spriteId string) (*Sprite, error)
}

// Asset handler manages all assets for a game
type AssetHandler interface {
	Initialize() error // initialize all assets
	FontHandler
	ImageHandler
	SpriteHandler
	Destroy() // free all resources
}

type FontRenderer interface {
	RenderText(resourceId string, text string, pos sdl.Point, textColor sdl.Color) error
}

type SpriteRenderer interface {
	RenderSprite(spriteId string) error
}

type LayerRenderer interface {
	CreateSpriteLayer(layerId int, pos sdl.Point) (*SpriteLayer, error)
	RenderLayers() error
}

type GridRenderer interface {
	RenderGrid(xSize, ySize int, color sdl.Color)
}

type TextureRenderer interface {
	RenderTexture(resourceId string, pos sdl.Point, textureWidth, textureHeight int32) error
	RenderRotatedTexture(resourceId string, pos sdl.Point, rotation float64, textureWidth, textureHeight int32) error
}

type RenderController interface {
	FontRenderer
	SpriteRenderer
	GridRenderer
	LayerRenderer
	TextureRenderer
}

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

type assets struct {
	audioAssets audioResourceMap
	fontAssets  fontResourceMap
	imageAssets imageResourceMap
	spriteBank  spriteMap
}

type audioResourceMap map[string]*AudioResource
type fontResourceMap map[string]*FontResource
type imageResourceMap map[string]ImageResource
type spriteMap map[string]*Sprite

type RenderFunction func()

type renderController struct {
	Window          *sdl.Window
	Renderer        *sdl.Renderer
	renderCallback  RenderFunction
	quit            bool
	world           *b2d.World
	width           int
	height          int
	spriteLayers    SpriteLayers
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

type SpriteLayer struct {
	Pos     sdl.Point
	Visible bool
	Sprites spriteMap
}
type SpriteLayers map[int]*SpriteLayer
