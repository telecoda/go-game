package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D3_GOPHER_RUN_IMAGE = "gopherrun"

	// sprites
	D3_GOPHER_STILL_SPRITE    = "gopher-still-sprite"
	D3_GOPHER_SPINNING_SPRITE = "gopher-spinning-sprite"
)

var demo3Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D3_GOPHER_RUN_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

var d3StillSprite *gogame.Sprite
var d3SpinningSprite *gogame.Sprite

// animation
var d3RotTextAnimSched *gogame.FunctionScheduler
var d3RotateTextureSpeed = time.Duration(20 * time.Millisecond)

var d3DebugInfoSched *gogame.FunctionScheduler

// init assets for demo 3
func initDemo3Assets() error {
	fmt.Printf("Loading Demo3 assets\n")

	for _, imageAsset := range demo3Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create sprites
	d3StillSprite = &gogame.Sprite{Id: D3_GOPHER_STILL_SPRITE, ImageAssetId: D3_GOPHER_RUN_IMAGE, Pos: sdl.Point{100, 200}, Width: 32, Height: 32, Rotation: 0.0, Visible: true}
	d3SpinningSprite = &gogame.Sprite{Id: D3_GOPHER_SPINNING_SPRITE, ImageAssetId: D3_GOPHER_RUN_IMAGE, Pos: sdl.Point{512, 400}, Width: 32, Height: 32, Rotation: 0.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(d3StillSprite)
	assetHandler.AddSprite(d3SpinningSprite)

	startDemo3Animation()

	return nil
}

// render screen for demo 3
func demo3RenderCallback() {
	renderer.ClearScreen(demoScreen.Color)

	renderTitle()

	renderer.RenderSprite(D3_GOPHER_STILL_SPRITE)
	renderer.RenderSprite(D3_GOPHER_SPINNING_SPRITE)
}

func unloadDemo3Assets() error {
	fmt.Printf("Unloading Demo3 assets\n")

	for _, imageAsset := range demo3Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	stopDemo3Animation()

	return nil
}

func startDemo3Animation() {

	// init animation vars
	angle = 0.0
	d3RotTextAnimSched = gogame.NewFunctionScheduler(d3RotateTextureSpeed, -1, demo3AnimateRotation)

	renderer.SetDebugInfo(false)
	d3DebugInfoSched = gogame.NewFunctionScheduler(time.Duration(5*time.Second), 1, func() { renderer.SetDebugInfo(true) })

	d3RotTextAnimSched.Start()
	d3DebugInfoSched.Start()
}

func stopDemo3Animation() {

	d3RotTextAnimSched.Destroy()
	d3DebugInfoSched.Destroy()

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

	d3SpinningSprite.Rotation = angle
	d3SpinningSprite.Width = sizeX
	d3SpinningSprite.Height = sizeY
}
