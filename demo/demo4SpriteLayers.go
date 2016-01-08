package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D4_GOPHER_RUN_IMAGE = "gopherrun"
)

var demo4Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D4_GOPHER_RUN_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

// sprite layers
var d4Layer0 *gogame.SpriteLayer
var d4Layer1 *gogame.SpriteLayer
var d4Layer2 *gogame.SpriteLayer

// animation
var d4ScrollAnimSched *gogame.FunctionScheduler
var d4ScrollSpeed = time.Duration(20 * time.Millisecond)
var d4AnimThreshold = time.Duration(22 * time.Millisecond)

var d4AnimLastTime time.Time
var d4AnimLastDuratiom time.Duration

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

	// nearest layer
	d4Layer0.Offset = sdl.Point{0, 64}
	d4Layer0.AbsPos = sdl.Point{-32, 0}
	d4Layer0.Wrap = true
	d4Layer0.Width = int32(demoWidth + 64)

	d4Layer1, err = demo4CreateLayer(1)
	if err != nil {
		return err
	}

	d4Layer1.Offset = sdl.Point{0, 32}
	d4Layer1.AbsPos = sdl.Point{-32, 0}
	d4Layer1.Wrap = true
	d4Layer1.Width = int32(demoWidth + 64)

	d4Layer2, err = demo4CreateLayer(2)
	if err != nil {
		return err
	}

	d4Layer2.Offset = sdl.Point{0, 0}
	d4Layer2.AbsPos = sdl.Point{-32, 0}
	d4Layer2.Wrap = true
	d4Layer2.Width = int32(demoWidth + 64)

	startDemo4Animation()

	return nil
}

func demo4CreateLayer(layerId int) (*gogame.SpriteLayer, error) {
	layer, err := renderer.CreateSpriteLayer(layerId, sdl.Point{0, 0})
	if err != nil {
		return nil, err

	}
	for x := 0; x < 17; x++ {
		for y := 0; y < 5; y++ {

			spriteId := fmt.Sprintf("d4sprite x:%2d y:%2d l:%2d", x, y, layerId)
			size := int32((3 - layerId) * 16)
			sprite := &gogame.Sprite{Id: spriteId, ImageAssetId: D4_GOPHER_RUN_IMAGE, Pos: sdl.Point{int32(x*64) + 32, int32(y*64 + 200)}, Width: size, Height: size, Rotation: 0.0, Visible: true}

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
	renderer.ClearScreen(demoScreen.Color)

	renderTitle()

	renderer.RenderLayers()
}

func unloadDemo4Assets() error {
	fmt.Printf("Unloading Demo4 assets\n")

	stopDemo4Animation()

	renderer.DestroySpriteLayer(0)
	renderer.DestroySpriteLayer(1)
	renderer.DestroySpriteLayer(2)

	for _, imageAsset := range demo4Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	return nil
}

func startDemo4Animation() {

	renderer.SetDebugInfo(false)
	// init animation vars
	//offset = 0.0
	d4ScrollAnimSched = gogame.NewFunctionScheduler(d4ScrollSpeed, -1, demo4AnimateScrolling)

	d4AnimLastTime = time.Now()

	d4ScrollAnimSched.Start()

}

func stopDemo4Animation() {

	d4ScrollAnimSched.Destroy()

}

func demo4AnimateScrolling() {

	// TEMP code to check animation timings
	currentTime := time.Now()

	d4AnimLastDuratiom = currentTime.Sub(d4AnimLastTime)

	d4AnimLastTime = currentTime

	if d4AnimLastDuratiom > d4AnimThreshold {
		fmt.Printf("Animation delay (nanos): %d expected (nanos): %d diff: %d\n", d4AnimLastDuratiom.Nanoseconds(), d4ScrollSpeed.Nanoseconds(), d4AnimLastDuratiom.Nanoseconds()-d4ScrollSpeed.Nanoseconds())
	}

	d4Layer0.Offset = sdl.Point{d4Layer0.Offset.X + 4, d4Layer0.Offset.Y}
	d4Layer1.Offset = sdl.Point{d4Layer1.Offset.X + 2, d4Layer1.Offset.Y}
	d4Layer2.Offset = sdl.Point{d4Layer2.Offset.X + 1, d4Layer2.Offset.Y}

}
