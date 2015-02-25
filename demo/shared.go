package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// fonts
	SHARED_FONT_8   = "shared-droidsans8"
	SHARED_FONT_16  = "shared-droidsans16"
	SHARED_FONT_24  = "shared-droidsans24"
	SHARED_FONT_48  = "shared-droidsans48"
	SHARED_FONT_128 = "shared-droidsans128"
)

var droidFontPath = "./demo_assets/fonts/droid-sans/DroidSans.ttf"

var sharedFonts = []gogame.FontAsset{
	{BaseAsset: gogame.BaseAsset{Id: SHARED_FONT_8, FilePath: droidFontPath}, Size: 8},
	{BaseAsset: gogame.BaseAsset{Id: SHARED_FONT_16, FilePath: droidFontPath}, Size: 16},
	{BaseAsset: gogame.BaseAsset{Id: SHARED_FONT_24, FilePath: droidFontPath}, Size: 24},
	{BaseAsset: gogame.BaseAsset{Id: SHARED_FONT_48, FilePath: droidFontPath}, Size: 48},
	{BaseAsset: gogame.BaseAsset{Id: SHARED_FONT_128, FilePath: droidFontPath}, Size: 128},
}

// init shared assets
func initSharedAssets() error {

	fmt.Printf("Loading shared assets\n")

	for _, fontAsset := range sharedFonts {
		err := assetHandler.AddFontAsset(fontAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	renderController.SetDefaultFont(SHARED_FONT_8)
	return nil
}

func renderTitle() {

	// Title
	// shadows
	renderController.RenderText(SHARED_FONT_128, demoScreen.Title, sdl.Point{X: 25, Y: 5}, 0.0, darkGrey, gogame.TOP, gogame.LEFT)
	// titles
	renderController.RenderText(SHARED_FONT_128, demoScreen.Title, sdl.Point{X: 20, Y: 10}, 0.0, black, gogame.TOP, gogame.LEFT)

	// Description
	renderController.RenderText(SHARED_FONT_16, demoScreen.Description, sdl.Point{50, 160}, 0.0, black, gogame.TOP, gogame.LEFT)

	// Code sample
	fromLine := 600
	lineSpacing := 20
	codeX := 50
	codeY := fromLine
	for _, codeLine := range demoScreen.CodeSample {
		renderController.RenderText(SHARED_FONT_16, codeLine, sdl.Point{int32(codeX), int32(codeY)}, 0.0, black, gogame.TOP, gogame.LEFT)
		codeY += lineSpacing
	}

	renderFPS()

}

func renderFPS() {
	fps := fmt.Sprintf("FPS:%2.2f", gogame.FramesPerSecond)
	// fps
	renderController.RenderText(SHARED_FONT_8, fps, sdl.Point{X: 5, Y: 0}, 0.0, black, gogame.TOP, gogame.LEFT)

}
