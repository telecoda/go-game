package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// fonts
	DEMO0_FONT_8   = "droidsans8"
	DEMO0_FONT_48  = "droidsans48"
	DEMO0_FONT_128 = "droidsans128"
)

var demo0fonts = []gogame.FontAsset{
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_8, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 8},
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_48, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 48},
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_128, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 128},
}

var titleString = "go-game"
var byString = "by: @telecoda"

// init assets for demo 0
func initDemo0Assets() error {

	fmt.Printf("Loading Demo0 assets\n")

	for _, fontAsset := range demo0fonts {
		err := assetHandler.AddFontAsset(fontAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	return nil
}

func unloadDemo0Assets() error {

	fmt.Printf("Unloading Demo0 assets\n")

	for _, fontRes := range demo0fonts {
		err := fontRes.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading a font asset:%s", err)
		}

	}

	return nil
}

// render screen for demo 0
func demo0RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	cx := int32(demoWidth / 2)
	cy := int32(demoHeight / 2)

	// shadows
	renderController.RenderText(DEMO0_FONT_128, titleString, sdl.Point{X: int32(cx + 5), Y: int32(cy + 5)}, darkGrey, gogame.MIDDLE, gogame.CENTER)
	renderController.RenderText(DEMO0_FONT_48, byString, sdl.Point{X: int32(cx + 50), Y: int32(cy + 100)}, darkGrey, gogame.MIDDLE, gogame.LEFT)
	// titles
	renderController.RenderText(DEMO0_FONT_128, titleString, sdl.Point{X: 0, Y: 0}, black, gogame.ABS_MIDDLE, gogame.ABS_CENTER)
	renderController.RenderText(DEMO0_FONT_48, byString, sdl.Point{X: int32(cx + 48), Y: int32(cy + 98)}, black, gogame.MIDDLE, gogame.LEFT)

}
