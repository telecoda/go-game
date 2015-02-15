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
	//textY := int32(demoHeight / 2)

	// valign
	fontPos := sdl.Point{textX, 150}
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: LEFT", fontPos, 0.0, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 200}
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: CENTER", fontPos, 0.0, black, gogame.TOP, gogame.CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 250}
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: ABS_CENTER", fontPos, 0.0, black, gogame.TOP, gogame.ABS_CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 300}
	renderController.RenderText(SHARED_FONT_24, "Horizontal alignment: RIGHT", fontPos, 0.0, black, gogame.TOP, gogame.RIGHT)
	renderFontPos(fontPos)

	// halign
	fontPos = sdl.Point{textX, 350}
	renderController.RenderText(SHARED_FONT_24, "Vertical alignment: TOP", fontPos, 0.0, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 400}
	renderController.RenderText(SHARED_FONT_24, "Vertical alignment: ABS_MIDDLE", fontPos, 0.0, black, gogame.ABS_MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 450}
	renderController.RenderText(SHARED_FONT_24, "Vertical alignment: MIDDLE", fontPos, 0.0, black, gogame.MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 550}
	renderController.RenderText(SHARED_FONT_24, "Vertical alignment: BOTTOM", fontPos, 0.0, black, gogame.BOTTOM, gogame.LEFT)
	renderFontPos(fontPos)

}

func renderFontPos(pos sdl.Point) {
	renderController.GetRenderer().SetDrawColor(255, 0, 0, 255)
	rect := sdl.Rect{pos.X, pos.Y, 5, 5}
	renderController.GetRenderer().FillRect(&rect)

}
