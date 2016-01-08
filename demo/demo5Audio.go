package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D5_GOPHER_TALK_IMAGE = "gopher-talk"

	// sprites
	D5_GOPHER_TALK_SPRITE = "gopher-talk-sprite"
)

var demo5Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D5_GOPHER_TALK_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-talk.png"}},
}

var d5TalkSprite *gogame.Sprite

// animation
var d5AnimSched *gogame.FunctionScheduler
var d5RotateSpeed = time.Duration(20 * time.Millisecond)

// init assets for demo 5
func initDemo5Assets() error {
	fmt.Printf("Loading Demo5 assets\n")

	for _, imageAsset := range demo5Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create sprites
	d5TalkSprite = &gogame.Sprite{Id: D5_GOPHER_TALK_SPRITE, ImageAssetId: D5_GOPHER_TALK_IMAGE, Pos: sdl.Point{512, 400}, Width: 132, Height: 100, Rotation: 0.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(d5TalkSprite)

	startDemo5Animation()

	return nil

}

// render screen for demo 5
func demo5RenderCallback() {
	renderer.ClearScreen(demoScreen.Color)

	renderTitle()

	renderer.RenderSprite(D5_GOPHER_TALK_SPRITE)

}

func unloadDemo5Assets() error {
	fmt.Printf("Unloading Demo5 assets\n")

	for _, imageAsset := range demo5Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	stopDemo5Animation()
	return nil
}

func startDemo5Animation() {

	// init animation vars
	d5AnimSched = gogame.NewFunctionScheduler(d5RotateSpeed, -1, demo5AnimateRotation)

	renderer.SetDebugInfo(false)

	d5AnimSched.Start()
}

func stopDemo5Animation() {

	d5AnimSched.Destroy()

}

func demo5AnimateRotation() {

}
