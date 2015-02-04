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

// Add a font resource & loads it into memory
func AddFontResource(resource FontResource) error {

	resource.loaded = false

	font, err := ttf.OpenFont(resource.FilePath, resource.Size)
	if err != nil {
		return fmt.Errorf("Error in AddFontResource:%s", err)
	}

	if font != nil {
		resource.font = font
		resource.loaded = true
	}

	fontResources[resource.Id] = &resource

	return nil

}

func GetFontResource(resourceId string) (*ttf.Font, error) {

	res, ok := fontResources[resourceId]
	if !ok {
		return nil, fmt.Errorf("Error: unknown font resource:%\n ", resourceId)
	}

	if res.font == nil {
		return nil, fmt.Errorf("Error: font not loaded:%s\n ", resourceId)
	}

	return res.font, nil
}

// Add an image resource & loads it into memory
func AddImageResource(resource ImageResource) error {

	resource.loaded = false

	image, err := img.Load(resource.FilePath)
	if err != nil {
		return fmt.Errorf("Failed to load image: %s\n", err)

	}
	texture, err := game.Renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %s\n", err)
	}

	resource.image = image
	resource.texture = texture
	resource.loaded = true

	imageResources[resource.Id] = resource

	return nil

}

func getImageResource(resourceId string) (*sdl.Surface, *sdl.Texture, error) {

	res, ok := imageResources[resourceId]
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
