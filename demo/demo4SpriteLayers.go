package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D4_GOPHER_RUN = "gopherrun"
)

var demo4Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D4_GOPHER_RUN, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

// sprite layers
var d4Layer0 *gogame.SpriteLayer
var d4Layer1 *gogame.SpriteLayer
var d4Layer2 *gogame.SpriteLayer

// animation
var d4ScrollAnimSched *gogame.FunctionScheduler
var d4ScrollSpeed = time.Duration(10 * time.Millisecond)

// init assets for demo 4
func initDemo4Assets() error {
	fmt.Printf("Loading Demo4 assets\n")

	var err error

	for _, imageAsset := range demo4Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create sprites

	d4Layer0, err = demo4CreateLayer(0)
	if err != nil {
		return err
	}

	d4Layer0.Wrap = false

	d4Layer1, err = demo4CreateLayer(1)
	if err != nil {
		return err
	}

	d4Layer1.Wrap = false

	d4Layer2, err = demo4CreateLayer(2)
	if err != nil {
		return err
	}

	d4Layer2.Wrap = true
	d4Layer2.Width = int32(demoWidth + 64)

	startDemo4Animation()

	return nil
}

func demo4CreateLayer(layerId int) (*gogame.SpriteLayer, error) {
	layer, err := renderController.CreateSpriteLayer(layerId, sdl.Point{0, 0})
	if err != nil {
		return nil, err

	}
	for x := 0; x < 25; x++ {
		for y := 0; y < 10; y++ {

			spriteId := fmt.Sprintf("d4sprite x:%2d y:%2d l:%2d", x, y, layerId)
			size := int32((3 - layerId) * 32)
			sprite := &gogame.Sprite{Id: spriteId, ImageAssetId: D4_GOPHER_RUN, Pos: sdl.Point{int32(x*40) + 32, int32(y*40) + 200}, Width: size, Height: size, Rotation: 0.0, Visible: true}

			// add sprite in sprite bank
			assetHandler.AddSprite(sprite)

			// add sprite to layer
			layer.AddSpriteToLayer(spriteId)

		}
	}

	return layer, nil

}

// render screen for demo 4
func demo4RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderLayers()
}

func unloadDemo4Assets() error {
	fmt.Printf("Unloading Demo4 assets\n")

	renderController.DestroySpriteLayer(0)
	renderController.DestroySpriteLayer(1)
	renderController.DestroySpriteLayer(2)

	for _, imageAsset := range demo4Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	stopDemo4Animation()

	return nil
}

func startDemo4Animation() {

	renderController.SetDebugInfo(false)
	// init animation vars
	//offset = 0.0
	d4ScrollAnimSched = gogame.NewFunctionScheduler(d4ScrollSpeed, -1, demo4AnimateScrolling)

	d4ScrollAnimSched.Start()

}

func stopDemo4Animation() {

	d4ScrollAnimSched.Destroy()

}

func demo4AnimateScrolling() {

	d4Layer0.Pos = sdl.Point{d4Layer0.Pos.X + 3, d4Layer0.Pos.Y}
	d4Layer1.Pos = sdl.Point{d4Layer1.Pos.X + 2, d4Layer1.Pos.Y}
	d4Layer2.Pos = sdl.Point{d4Layer2.Pos.X + 1, d4Layer2.Pos.Y}

}
