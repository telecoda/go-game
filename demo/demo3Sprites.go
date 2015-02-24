package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D3_GOPHER_RUN = "gopherrun"

	// sprites
	D3_GOPHER_RUN_SPRITE = "gopherrun-sprite"
)

var demo3Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D3_GOPHER_RUN, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

var d3GopherSprite *gogame.Sprite

// init assets for demo 3
func initDemo3Assets() error {
	fmt.Printf("Loading Demo3 assets\n")

	for _, imageAsset := range demo3Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create sprite
	d3GopherSprite = &gogame.Sprite{Id: D3_GOPHER_RUN_SPRITE, ImageAssetId: D3_GOPHER_RUN, Pos: sdl.Point{100, 200}, Width: 32, Height: 32, Rotation: 0.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(D3_GOPHER_RUN_SPRITE, d3GopherSprite)

	startDemo3Animation()

	return nil
}

// render screen for demo 3
func demo3RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderSprite(D3_GOPHER_RUN_SPRITE)
}

func unloadDemo3Assets() error {
	fmt.Printf("Unloading Demo3 assets\n")

	for _, imageAsset := range demo3Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	rotTextAnimSched.Destroy()

	return nil
}

func startDemo3Animation() {

	// init animation vars
	angle = 0.0
	rotTextAnimSched = gogame.NewFunctionScheduler(rotateTextureSpeed, -1, demo2AnimateRotation)

	rotTextAnimSched.Start()

}

func stopDemo3Animation() {

	rotTextAnimSched.Destroy()

}

func demo3AnimateRotation() {

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
