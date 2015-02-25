package gogame

import (
	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
	mix "github.com/veandco/go-sdl2/sdl_mixer"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

type GameEngine interface {
	NewGame(winTitle string, winWidth, winHeight int, renderCallback RenderFunction) error
	Destroy()
	EventLoop()
}

type Asset interface {
	save() error
	Load() error
	Unload() error
	Destroy() error
}

type AudioManager interface {
	AddAudioAsset(asset AudioAsset, load bool) error // add a font
}

type FontManager interface {
	AddFontAsset(asset FontAsset, load bool) error // add a font
}

type ImageManager interface {
	AddImageAsset(asset ImageAsset, load bool) error // add an image
}

type SpriteManager interface {
	AddSprite(sprite *Sprite) error
	GetSprite(spriteId string) (*Sprite, error)
}

// Asset manager manages all assets for a game
type AssetManager interface {
	Initialize() error // initialize all assets
	AudioManager
	FontManager
	ImageManager
	SpriteManager
	Destroy() // free all assets
}

type AudioPlayer interface {
	PlayAudio(assetId string, loops int) error
}

type FontRenderer interface {
	RenderText(assetId string, text string, pos sdl.Point, angle float64, textColor sdl.Color, vAlign VAlign, hAlign HAlign) error
}

type SpriteRenderer interface {
	RenderSprite(spriteId string) error
}

type LayerRenderer interface {
	CreateSpriteLayer(layerId int, pos sdl.Point) (*SpriteLayer, error)
	DestroySpriteLayer(layerId int) error
	RenderLayers() error
}

type GridRenderer interface {
	RenderGrid(xSize, ySize int, color sdl.Color)
	RenderGridInRect(rect sdl.Rect, xSize, ySize int, color sdl.Color)
}

type TextureRenderer interface {
	RenderTexture(assetId string, pos sdl.Point, textureWidth, textureHeight int32) error
	RenderRotatedTexture(assetId string, pos sdl.Point, rotation float64, textureWidth, textureHeight int32) error
}

type RenderController interface {
	ClearScreen(color sdl.Color)
	FontRenderer
	SpriteRenderer
	GridRenderer
	LayerRenderer
	TextureRenderer
	GetRenderer() *sdl.Renderer
	SetCallback(callback RenderFunction)
	SetDefaultFont(fontId string) error
	SetDebugInfo(enabled bool)
}

type EventHandler interface {
	SetCallback(calllback EventReceiverFunction)
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

type renderController struct {
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
	Pos     sdl.Point
	Visible bool
	Wrap    bool
	Width   int32
	Height  int32
	Sprites spriteMap
}
type SpriteLayers map[int]*SpriteLayer
