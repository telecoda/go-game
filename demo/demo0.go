package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// fonts
	DEMO0_FONT_8  = "droidsans8"
	DEMO0_FONT_16 = "droidsans16"
	DEMO0_FONT_48 = "droidsans48"
)

var demo0fonts = []gogame.FontAsset{
	{Id: DEMO0_FONT_8, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf", Size: 8},
	{Id: DEMO0_FONT_16, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf", Size: 16},
	{Id: DEMO0_FONT_48, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf", Size: 48},
}

// init assets for demo 0
func initDemo0Assets() error {

	fmt.Printf("Loading Demo0 assets\n")

	for _, fontRes := range demo0fonts {
		err := assetHandler.AddFontAsset(fontRes, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	return nil
}

func unloadDemo0Assets() error {

	fmt.Printf("Unloading Demo0 assets\n")

	for _, fontRes := range demo0fonts {
		err := assetHandler.UnloadFontAsset(fontRes.Id)
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading a font asset:%s", err)
		}

	}

	return nil
}

// render screen for demo 0
func demo0RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)
	renderController.RenderText(DEMO0_FONT_16, "Test text", sdl.Point{X: 20, Y: 60}, black)
	renderController.RenderText(DEMO0_FONT_48, "Test text", sdl.Point{X: 20, Y: 80}, red)

}
