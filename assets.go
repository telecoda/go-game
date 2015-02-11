package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func init() {
	// initialise library
	ttf.Init()

}

/*func initSystemFont() error {

	fontRes := FontAsset{Id: SYSTEM_FONT_ID, FilePath: "./test_assets/fonts/droid-sans/DroidSans.ttf", Size: 8}
	err := gameAssets.AddFontAsset(fontRes)
	if err != nil {
		return fmt.Errorf("Failed to load system font. Error:%s", err)
	}

	return nil
}*/

func (a *assets) Destroy() {
	a.audioAssets.Destroy()
	a.fontAssets.Destroy()
	a.imageAssets.Destroy()
	a.spriteBank.Destroy()
}

func (a *assets) Initialize() error {
	a.audioAssets = make(audioAssetMap)
	a.fontAssets = make(fontAssetMap)
	a.imageAssets = make(imageAssetMap)
	a.spriteBank = make(spriteMap)

	//return initSystemFont()
	return nil
}

// Add a font asset & loads it into memory
func (a *assets) AddFontAsset(asset FontAsset) error {

	asset.loaded = false

	font, err := ttf.OpenFont(asset.FilePath, asset.Size)
	if err != nil {
		return fmt.Errorf("Error in AddFontAsset:%s", err)
	}

	if asset.Size < 1 {
		return fmt.Errorf("Error: font size must be larger than %d", asset.Size)
	}

	if font != nil {
		asset.font = font
		asset.loaded = true
	}

	a.fontAssets[asset.Id] = &asset

	fmt.Printf("Font:%s loaded\n", asset.Id)

	return nil

}

func (a *assets) getFontAsset(assetId string) (*ttf.Font, error) {

	res, ok := a.fontAssets[assetId]
	if !ok {
		return nil, fmt.Errorf("Error: unknown font asset:%s\n ", assetId)
	}

	if res.font == nil {
		return nil, fmt.Errorf("Error: font not loaded:%s\n ", assetId)
	}

	return res.font, nil
}

// Add an image asset & loads it into memory
func (a *assets) AddImageAsset(asset ImageAsset) error {

	asset.loaded = false

	image, err := img.Load(asset.FilePath)
	if err != nil {
		return fmt.Errorf("Failed to load image: %s\n", err)

	}
	texture, err := rendCont.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %s\n", err)
	}

	asset.image = image
	asset.texture = texture
	asset.loaded = true

	a.imageAssets[asset.Id] = asset

	return nil

}

func (a *assets) getImageAsset(assetId string) (*sdl.Surface, *sdl.Texture, error) {

	res, ok := a.imageAssets[assetId]
	if !ok {
		return nil, nil, fmt.Errorf("Warning: unknown image asset:%\n ", assetId)
	}

	if res.image == nil {
		return nil, nil, fmt.Errorf("Warning: image not loaded:%\n ", assetId)
	}

	if res.texture == nil {
		return nil, nil, fmt.Errorf("Warning: texture not loaded:%\n ", assetId)
	}

	return res.image, res.texture, nil
}

func (a audioAssetMap) Destroy() {

	for _, res := range a {
		fmt.Printf("Freeing audio asset:%s\n", res.Id)
	}

}

func (i imageAssetMap) Destroy() {

	for _, res := range i {
		fmt.Printf("Freeing image asset:%s\n", res.Id)
		res.image.Free()
		res.texture.Destroy()
	}

}

// close all font assets
func (f fontAssetMap) Destroy() {

	for _, res := range f {
		fmt.Printf("Freeing font asset:%s\n", res.Id)
		res.font.Close()
	}

}

func (s spriteMap) Destroy() {

	for _, sprite := range s {
		fmt.Printf("Freeing sprite asset:%s\n", sprite.Id)
	}

}
