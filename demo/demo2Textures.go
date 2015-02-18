package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (

	// images
	GOPHER_RUN = "gopherrun"
	/*JAVA_DUKE    = "javaduke"
	FLOOR_IMAGE  = "floorimage"
	CLOUD1_IMAGE = "cloud1image"
	CLOUD2_IMAGE = "cloud2image"
	CLOUD3_IMAGE = "cloud3image"
	CLOUD4_IMAGE = "cloud4image"*/

)

var demo2Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: GOPHER_RUN, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
	//	{Id: JAVA_DUKE, FilePath: "./assets/images/sprites/java-duke.png"},
	//	{Id: FLOOR_IMAGE, FilePath: "./assets/images/sprites/floor.png"},
	//	{Id: CLOUD1_IMAGE, FilePath: "./assets/images/sprites/cloud1.png"},
	//	{Id: CLOUD2_IMAGE, FilePath: "./assets/images/sprites/cloud2.png"},
	//	{Id: CLOUD3_IMAGE, FilePath: "./assets/images/sprites/cloud3.png"},
	//	{Id: CLOUD4_IMAGE, FilePath: "./assets/images/sprites/cloud4.png"},
}

var angle = 0.0
var sizeX = int32(32)
var sizeY = int32(32)
var sizeVelocity = int32(5)
var minSize = int32(5)
var maxSize = int32(350)

// animation
var rotTextAnimSched *gogame.FunctionScheduler
var rotateTextureSpeed = time.Duration(20 * time.Millisecond)

var textureGrid1 = sdl.Rect{100, 200, 64, 64}
var textureGrid2 = sdl.Rect{200, 200, 64, 64}
var textureGrid3 = sdl.Rect{300, 200, 64, 64}
var textureGrid4 = sdl.Rect{400, 200, 64, 64}
var textureGrid5 = sdl.Rect{500, 200, 350, 350}

// init assets for demo 2
func initDemo2Assets() error {
	fmt.Printf("Loading Demo2 assets\n")

	for _, imageAsset := range demo2Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	startDemo2Animation()

	return nil
}

// render screen for demo 2
func demo2RenderCallback() {

	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderGridInRect(textureGrid1, 4, 4, black)
	renderController.RenderGridInRect(textureGrid2, 4, 4, black)
	renderController.RenderGridInRect(textureGrid3, 4, 4, black)
	renderController.RenderGridInRect(textureGrid4, 4, 4, black)
	renderController.RenderGridInRect(textureGrid5, 4, 4, black)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{100, 200}, 32, 32)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{200, 200}, 64, 32)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{300, 200}, 32, 64)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{400, 200}, 64, 64)
	renderController.RenderRotatedTexture(GOPHER_RUN, sdl.Point{500, 200}, angle, sizeX, sizeY)

	return
}

func unloadDemo2Assets() error {
	fmt.Printf("Unloading Demo2 assets\n")

	for _, imageAsset := range demo2Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	return nil
}

func startDemo2Animation() {

	// init animation vars
	angle = 0.0
	/*	textLabels = make([]textTyping, 8)

		textLabels[0] = textTyping{actualText: "Horizontal alignment: LEFT", typedText: ""}
		textLabels[1] = textTyping{actualText: "Horizontal alignment: CENTER", typedText: ""}
		textLabels[2] = textTyping{actualText: "Horizontal alignment: ABS_CENTER", typedText: ""}
		textLabels[3] = textTyping{actualText: "Horizontal alignment: RIGHT", typedText: ""}
		textLabels[4] = textTyping{actualText: "Vertical alignment: TOP", typedText: ""}
		textLabels[5] = textTyping{actualText: "Vertical alignment: ABS_MIDDLE", typedText: ""}
		textLabels[6] = textTyping{actualText: "Vertical alignment: MIDDLE", typedText: ""}
		textLabels[7] = textTyping{actualText: "Vertical alignment: BOTTOM", typedText: ""}
		currentLabel = 0
	*/
	rotTextAnimSched = gogame.NewFunctionScheduler(rotateTextureSpeed, -1, demo2AnimateRotation)

	rotTextAnimSched.Start()

}

func demo2AnimateRotation() {

	// rotate texture
	angle++
	if angle > 360 {
		angle = angle - 360
	}

	// increase & decrease texture size
	sizeX = sizeX + sizeVelocity
	sizeY = sizeY + sizeVelocity

	if sizeX > maxSize {
		sizeVelocity = sizeVelocity * -1
	}
	if sizeX < minSize {
		sizeVelocity = sizeVelocity * -1
	}
}
