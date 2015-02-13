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
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_8, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 8},
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_16, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 16},
	{BaseAsset: gogame.BaseAsset{Id: DEMO0_FONT_48, FilePath: "./demo_assets/fonts/droid-sans/DroidSans.ttf"}, Size: 48},
}

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

	textX := int32(gameWidth / 2)
	textY := int32(gameHeight / 2)

	// valign
	renderController.RenderText(DEMO0_FONT_16, "Horizontal alignment: LEFT", sdl.Point{X: textX, Y: 100}, black, gogame.TOP, gogame.LEFT)
	renderController.RenderText(DEMO0_FONT_16, "Horizontal alignment: CENTER", sdl.Point{X: textX, Y: 120}, black, gogame.TOP, gogame.CENTER)
	renderController.RenderText(DEMO0_FONT_16, "Horizontal alignment: ABS_CENTER", sdl.Point{X: textX, Y: 140}, black, gogame.TOP, gogame.ABS_CENTER)
	renderController.RenderText(DEMO0_FONT_16, "Horizontal alignment: RIGHT", sdl.Point{X: textX, Y: 160}, black, gogame.TOP, gogame.RIGHT)

	// halign
	renderController.RenderText(DEMO0_FONT_16, "Vertical alignment: TOP", sdl.Point{X: 0, Y: textY}, black, gogame.TOP, gogame.LEFT)
	renderController.RenderText(DEMO0_FONT_16, "Vertical alignment: MIDDLE", sdl.Point{X: textX / 2, Y: textY}, black, gogame.MIDDLE, gogame.LEFT)
	renderController.RenderText(DEMO0_FONT_16, "Vertical alignment: ABS_MIDDLE", sdl.Point{X: textX, Y: textY}, black, gogame.ABS_MIDDLE, gogame.LEFT)
	renderController.RenderText(DEMO0_FONT_16, "Vertical alignment: BOTTOM", sdl.Point{X: textX/2 + textX, Y: textY}, black, gogame.BOTTOM, gogame.LEFT)

	//renderController.RenderText(DEMO0_FONT_16, "Test text", sdl.Point{X: 20, Y: 60}, black, gogame.TOP, gogame.LEFT)
	//renderController.RenderText(DEMO0_FONT_16, "Test text", sdl.Point{X: 20, Y: 60}, black, gogame.TOP, gogame.LEFT)
	//renderController.RenderText(DEMO0_FONT_48, "Test text", sdl.Point{X: 20, Y: 80}, red, gogame.MIDDLE, gogame.CENTER)

}
