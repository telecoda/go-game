package gogame

import "github.com/veandco/go-sdl2/sdl"

type GameEngine interface {
	Destroy()
	EventLoop()
	GetRenderer() Renderer
	GetAssetManger() AssetManager
	GetAudioPlayer() AudioPlayer
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

type Renderer interface {
	ClearScreen(color sdl.Color)
	ClearWorld()
	FontRenderer
	SpriteRenderer
	GridRenderer
	LayerRenderer
	TextureRenderer
	GetRenderer() *sdl.Renderer
	Present()
	SetCallback(callback RenderFunction)
	SetDefaultFont(fontId string) error
	SetDebugInfo(enabled bool)
}

type EventHandler interface {
	SetCallback(calllback EventReceiverFunction)
}
