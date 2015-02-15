package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// fonts
	DEMO1_FONT_8  = "droidsans8"
	DEMO1_FONT_16 = "droidsans16"
	DEMO1_FONT_48 = "droidsans48"
)

var demo1fonts = []gogame.FontAsset{}

// init assets for demo 1
func initDemo1Assets() error {

	fmt.Printf("Loading Demo1 assets\n")

	for _, fontAsset := range demo1fonts {
		err := assetHandler.AddFontAsset(fontAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	return nil
}

func unloadDemo1Assets() error {

	fmt.Printf("Unloading Demo1 assets\n")

	for _, fontRes := range demo1fonts {
		err := fontRes.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading a font asset:%s", err)
		}

	}

	return nil
}

// render screen for demo 1
func demo1RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	textX := int32(demoWidth / 2)
	textY := int32(demoHeight / 2)

	// valign
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: LEFT", sdl.Point{X: textX, Y: 150}, 0.0, black, gogame.TOP, gogame.LEFT)
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: CENTER", sdl.Point{X: textX, Y: 180}, 0.0, black, gogame.TOP, gogame.CENTER)
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: ABS_CENTER", sdl.Point{X: textX, Y: 210}, 0.0, black, gogame.TOP, gogame.ABS_CENTER)
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: RIGHT", sdl.Point{X: textX, Y: 240}, 0.0, black, gogame.TOP, gogame.RIGHT)

	// halign
	renderController.RenderText(SHARED_FONT_16, "Vertical alignment: TOP", sdl.Point{X: 0, Y: textY}, 0.0, black, gogame.TOP, gogame.LEFT)
	renderController.RenderText(SHARED_FONT_16, "Vertical alignment: MIDDLE", sdl.Point{X: textX / 2, Y: textY}, 0.0, black, gogame.MIDDLE, gogame.LEFT)
	renderController.RenderText(SHARED_FONT_16, "Vertical alignment: ABS_MIDDLE", sdl.Point{X: textX, Y: textY}, 0.0, black, gogame.ABS_MIDDLE, gogame.LEFT)
	renderController.RenderText(SHARED_FONT_16, "Vertical alignment: BOTTOM", sdl.Point{X: textX/2 + textX, Y: textY}, 0.0, black, gogame.BOTTOM, gogame.LEFT)

}
