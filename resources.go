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

func (a *assets) Initialize() {
	a.audioAssets = make(audioResourceMap)
	a.fontAssets = make(fontResourceMap)
	a.imageAssets = make(imageResourceMap)
	a.spriteBank = make(spriteMap)
}

// Add a font resource & loads it into memory
func (a *assets) AddFontResource(resource FontResource) error {

	resource.loaded = false

	font, err := ttf.OpenFont(resource.FilePath, resource.Size)
	if err != nil {
		return fmt.Errorf("Error in AddFontResource:%s", err)
	}

	if resource.Size < 1 {
		return fmt.Errorf("Error: font size must be larger than %d", resource.Size)
	}

	if font != nil {
		resource.font = font
		resource.loaded = true
	}

	a.fontAssets[resource.Id] = &resource

	return nil

}

func (a *assets) getFontResource(resourceId string) (*ttf.Font, error) {

	res, ok := a.fontAssets[resourceId]
	if !ok {
		return nil, fmt.Errorf("Error: unknown font resource:%\n ", resourceId)
	}

	if res.font == nil {
		return nil, fmt.Errorf("Error: font not loaded:%s\n ", resourceId)
	}

	return res.font, nil
}

// Add an image resource & loads it into memory
func (a *assets) AddImageResource(resource ImageResource) error {

	resource.loaded = false

	image, err := img.Load(resource.FilePath)
	if err != nil {
		return fmt.Errorf("Failed to load image: %s\n", err)

	}
	texture, err := rendCont.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %s\n", err)
	}

	resource.image = image
	resource.texture = texture
	resource.loaded = true

	a.imageAssets[resource.Id] = resource

	return nil

}

func (a *assets) getImageResource(resourceId string) (*sdl.Surface, *sdl.Texture, error) {

	res, ok := a.imageAssets[resourceId]
	if !ok {
		return nil, nil, fmt.Errorf("Warning: unknown image resource:%\n ", resourceId)
	}

	if res.image == nil {
		return nil, nil, fmt.Errorf("Warning: image not loaded:%\n ", resourceId)
	}

	if res.texture == nil {
		return nil, nil, fmt.Errorf("Warning: texture not loaded:%\n ", resourceId)
	}

	return res.image, res.texture, nil
}

func (a audioResourceMap) Destroy() {

	for _, res := range a {
		fmt.Printf("Freeing audio resource:%s\n", res.Id)
	}

}

func (i imageResourceMap) Destroy() {

	for _, res := range i {
		fmt.Printf("Freeing image resource:%s\n", res.Id)
		res.image.Free()
		res.texture.Destroy()
	}

}

// close all font resources
func (f fontResourceMap) Destroy() {

	for _, res := range f {
		fmt.Printf("Freeing font resource:%s\n", res.Id)
		res.font.Close()
	}

}

func (s spriteMap) Destroy() {

	for _, sprite := range s {
		fmt.Printf("Freeing sprite resource:%s\n", sprite.Id)
	}

}
