package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D7_GOPHER_RUN_IMAGE = "gopher-run"

	// sprites
	D7_GOPHER_RUN_SPRITE = "gopher-run-sprite"
)

var demo7Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D7_GOPHER_RUN_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
}

var d7Layer0 *gogame.SpriteLayer

var d7RunSprite1 *gogame.Sprite
var d7RunSprite2 *gogame.Sprite
var d7RunSprite3 *gogame.Sprite
var d7RunSprite4 *gogame.Sprite
var d7RunSprite5 *gogame.Sprite

// animation
var d7AnimSched *gogame.FunctionScheduler
var d7RotateSpeed = time.Duration(20 * time.Millisecond)

// init assets for demo 7
func initDemo7Assets() error {
	fmt.Printf("Loading Demo7 assets\n")

	for _, imageAsset := range demo7Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	d7Layer0, err := renderController.CreateSpriteLayer(0, sdl.Point{0, 0})
	if err != nil {
		return err

	}

	gopherWidth := int32(52)
	gopherHeight := int32(64)

	xOffset := gopherWidth / 2
	yOffset := gopherHeight / 2

	d7RunSprite1 = &gogame.Sprite{Id: D7_GOPHER_RUN_SPRITE + "1", ImageAssetId: D7_GOPHER_RUN_IMAGE, Pos: sdl.Point{512, 250}, Width: gopherWidth, Height: gopherHeight, Rotation: 0.0, Visible: true}
	d7RunSprite2 = &gogame.Sprite{Id: D7_GOPHER_RUN_SPRITE + "2", ImageAssetId: D7_GOPHER_RUN_IMAGE, Pos: sdl.Point{512 + gopherWidth, 250}, Width: gopherWidth, Height: gopherHeight, Rotation: 0.0, Visible: true}
	d7RunSprite3 = &gogame.Sprite{Id: D7_GOPHER_RUN_SPRITE + "3", ImageAssetId: D7_GOPHER_RUN_IMAGE, Pos: sdl.Point{512 + gopherWidth*2, 250}, Width: gopherWidth, Height: gopherHeight, Rotation: 0.0, Visible: true}
	d7RunSprite4 = &gogame.Sprite{Id: D7_GOPHER_RUN_SPRITE + "4", ImageAssetId: D7_GOPHER_RUN_IMAGE, Pos: sdl.Point{512 + gopherWidth*3, 250}, Width: gopherWidth, Height: gopherHeight, Rotation: 0.0, Visible: true}
	d7RunSprite5 = &gogame.Sprite{Id: D7_GOPHER_RUN_SPRITE + "5", ImageAssetId: D7_GOPHER_RUN_IMAGE, Pos: sdl.Point{512 + gopherWidth*4, 250}, Width: gopherWidth, Height: gopherHeight, Rotation: 0.0, Visible: true}

	assetHandler.AddSprite(d7RunSprite1)
	assetHandler.AddSprite(d7RunSprite2)
	assetHandler.AddSprite(d7RunSprite3)
	assetHandler.AddSprite(d7RunSprite4)
	assetHandler.AddSprite(d7RunSprite5)

	d7Layer0.AddSpriteToLayer(d7RunSprite1.Id)
	d7Layer0.AddSpriteToLayer(d7RunSprite2.Id)
	d7Layer0.AddSpriteToLayer(d7RunSprite3.Id)
	d7Layer0.AddSpriteToLayer(d7RunSprite4.Id)
	d7Layer0.AddSpriteToLayer(d7RunSprite5.Id)

	d7RunSprite1.EnablePhysics(gogame.ImmovableMass)
	d7RunSprite2.EnablePhysics(100)
	d7RunSprite3.EnablePhysics(100)
	d7RunSprite4.EnablePhysics(100)
	d7RunSprite5.EnablePhysics(100)

	d7RunSprite1.JoinTo(d7RunSprite2, sdl.Point{int32(xOffset), int32(yOffset)})
	d7RunSprite2.JoinTo(d7RunSprite3, sdl.Point{int32(xOffset), int32(-yOffset)})
	d7RunSprite3.JoinTo(d7RunSprite4, sdl.Point{int32(xOffset), int32(yOffset)})
	d7RunSprite4.JoinTo(d7RunSprite5, sdl.Point{int32(xOffset), int32(-yOffset)})
	startDemo7Animation()

	return nil

}

// render screen for demo 7
func demo7RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderLayers()
}

func unloadDemo7Assets() error {
	fmt.Printf("Unloading Demo7 assets\n")

	for _, imageAsset := range demo5Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	stopDemo7Animation()
	return nil
}

func startDemo7Animation() {

	d7AnimSched = gogame.NewFunctionScheduler(d7RotateSpeed, -1, demo7AnimateRotation)

	renderController.SetDebugInfo(true)

	d7AnimSched.Start()
}

func stopDemo7Animation() {

	renderController.SetDebugInfo(false)

	d7AnimSched.Destroy()

}

func demo7AnimateRotation() {

}
