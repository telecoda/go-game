package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (

	// images
	D2_GOPHER_RUN = "gopherrun"
)

var demo2Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D2_GOPHER_RUN, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

var angle = 0.0
var sizeX = int32(32)
var sizeY = int32(32)
var sizeVelocity = int32(5)
var minSize = int32(5)
var maxSize = int32(350)

// animation
var d2RotTextAnimSched *gogame.FunctionScheduler
var d2RotateTextureSpeed = time.Duration(20 * time.Millisecond)

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
	renderController.RenderTexture(D2_GOPHER_RUN, sdl.Point{100, 200}, 32, 32)
	renderController.RenderTexture(D2_GOPHER_RUN, sdl.Point{200, 200}, 64, 32)
	renderController.RenderTexture(D2_GOPHER_RUN, sdl.Point{300, 200}, 32, 64)
	renderController.RenderTexture(D2_GOPHER_RUN, sdl.Point{400, 200}, 64, 64)
	renderController.RenderRotatedTexture(D2_GOPHER_RUN, sdl.Point{500, 200}, angle, sizeX, sizeY)

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

	stopDemo2Animation()

	return nil
}

func startDemo2Animation() {

	// init animation vars
	angle = 0.0
	d2RotTextAnimSched = gogame.NewFunctionScheduler(d2RotateTextureSpeed, -1, demo2AnimateRotation)

	d2RotTextAnimSched.Start()

}

func stopDemo2Animation() {

	d2RotTextAnimSched.Destroy()

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
