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

	return nil
}

// Add a font asset & loads it into memory
func (a *assets) AddFontAsset(asset FontAsset, load bool) error {

	if asset.Size < 1 {
		return fmt.Errorf("Error: font size must be larger than %d", asset.Size)
	}

	asset.loaded = false

	if load {
		err := asset.Load()
		if err != nil {
			return err
		}
	}

	fmt.Printf("Adding font:%s \n", asset.Id)

	return asset.Save()

}

func (a *FontAsset) Add(load bool) error {

	if a.Size < 1 {
		return fmt.Errorf("Error: font size must be larger than %d", a.Size)
	}

	a.loaded = false

	if load {
		err := a.Load()
		if err != nil {
			return err
		}
	}

	fmt.Printf("Adding font:%s \n", a.Id)

	return a.Save()

}

func (a *FontAsset) Save() error {

	gameAssets.fontAssets[a.Id] = a

	return nil

}

func (a *FontAsset) Load() error {

	// already loaded
	if a.loaded {
		return nil
	}

	font, err := ttf.OpenFont(a.FilePath, a.Size)
	if err != nil {
		return fmt.Errorf("Error in LoadFontAsset:%s", err)
	}

	if font != nil {
		a.font = font
		a.loaded = true
	}

	return a.Save()
}

func (a *FontAsset) Unload() error {

	if a.loaded {
		if a.font != nil {
			fmt.Printf("Unloading font asset:%s\n", a.Id)
			a.font.Close()
			a.loaded = false
			err := a.Save()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getFontAsset(assetId string) (*FontAsset, error) {
	asset, ok := gameAssets.fontAssets[assetId]
	if !ok {
		return nil, fmt.Errorf("Error: unknown font asset:%s\n ", assetId)
	}
	return asset, nil
}

func getFont(assetId string) (*ttf.Font, error) {

	asset, err := getFontAsset(assetId)
	if err != nil {
		return nil, err
	}

	if asset.font == nil {
		return nil, fmt.Errorf("Error: font not loaded:%s\n ", assetId)
	}

	return asset.font, nil
}

// close all font assets
func (f fontAssetMap) Destroy() {

	for _, asset := range f {
		fmt.Printf("Destroying font asset:%s\n", asset.Id)
		asset.Unload()
	}

}

// Add an image asset & loads it into memory
func (a *assets) AddImageAsset(asset ImageAsset, load bool) error {

	asset.loaded = false

	if load {
		err := asset.Load()
		if err != nil {
			return err
		}

	}

	return asset.Save()

}

func (a *ImageAsset) Save() error {

	gameAssets.imageAssets[a.Id] = a

	return nil

}

func (a *ImageAsset) Load() error {

	// already loaded
	if a.loaded {
		return nil
	}

	image, err := img.Load(a.FilePath)
	if err != nil {
		return fmt.Errorf("Failed to load image: %s\n", err)

	}
	texture, err := rendCont.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %s\n", err)
	}

	a.image = image
	a.texture = texture
	a.loaded = true

	return a.Save()
}

func (a *ImageAsset) Unload() error {

	if a.loaded {
		if a.image != nil {
			fmt.Printf("Unloading image asset:%s\n", a.Id)
			a.image.Free()
			a.texture.Destroy()
			a.loaded = false
			err := a.Save()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getImageAsset(assetId string) (*ImageAsset, error) {
	asset, ok := gameAssets.imageAssets[assetId]
	if !ok {
		return nil, fmt.Errorf("Error: unknown image asset:%s\n ", assetId)
	}
	return asset, nil
}

func getImage(assetId string) (*sdl.Surface, *sdl.Texture, error) {

	asset, ok := gameAssets.imageAssets[assetId]
	if !ok {
		return nil, nil, fmt.Errorf("Warning: unknown image asset:%\n ", assetId)
	}

	if asset.image == nil {
		return nil, nil, fmt.Errorf("Warning: image not loaded:%\n ", assetId)
	}

	if asset.texture == nil {
		return nil, nil, fmt.Errorf("Warning: texture not loaded:%\n ", assetId)
	}

	return asset.image, asset.texture, nil
}

func (a audioAssetMap) Destroy() {

	for _, res := range a {
		fmt.Printf("Freeing audio asset:%s\n", res.Id)
	}

}

func (i imageAssetMap) Destroy() {

	for _, asset := range i {
		fmt.Printf("Freeing image asset:%s\n", asset.Id)
		asset.Unload()
	}

}

func (s spriteMap) Destroy() {

	for _, sprite := range s {
		fmt.Printf("Freeing sprite asset:%s\n", sprite.Id)
	}

}
