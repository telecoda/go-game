package main

import (
	"fmt"

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
var maxSize = int32(500)

// init assets for demo 2
func initDemo2Assets() error {
	fmt.Printf("Loading Demo2 assets\n")

	for _, imageAsset := range demo2Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	//startDemo2Animation()

	return nil
}

// render screen for demo 2
func demo2RenderCallback() {

	angle++
	if angle > 360 {
		angle = angle - 360
	}

	sizeX = sizeX + sizeVelocity
	sizeY = sizeY + sizeVelocity

	if sizeX > maxSize {
		sizeVelocity = sizeVelocity * -1
	}
	if sizeX < minSize {
		sizeVelocity = sizeVelocity * -1
	}
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderTexture(GOPHER_RUN, sdl.Point{50, 200}, 32, 32)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{150, 200}, 64, 32)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{250, 200}, 32, 64)
	renderController.RenderTexture(GOPHER_RUN, sdl.Point{350, 200}, 64, 64)
	renderController.RenderRotatedTexture(GOPHER_RUN, sdl.Point{450, 200}, angle, sizeX, sizeY)

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
