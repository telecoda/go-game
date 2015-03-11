package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D6_GOPHER_RUN_IMAGE  = "gopherrun"
	D6_FLOOR_SHORT_IMAGE = "floor-short"
	D6_FLOOR_IMAGE       = "floor"
)

var demo6Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D6_GOPHER_RUN_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
	{BaseAsset: gogame.BaseAsset{Id: D6_FLOOR_SHORT_IMAGE, FilePath: "./demo_assets/images/sprites/floor-short.png"}},
	{BaseAsset: gogame.BaseAsset{Id: D6_FLOOR_IMAGE, FilePath: "./demo_assets/images/sprites/floor.png"}},
}

var d6GopherCount int

var d6Layer0 *gogame.SpriteLayer

var d6MainFloorSprite *gogame.Sprite

var d6FloorSprite1 *gogame.Sprite
var d6FloorSprite2 *gogame.Sprite
var d6FloorSprite3 *gogame.Sprite
var d6FloorSprite4 *gogame.Sprite
var d6FloorSprite5 *gogame.Sprite

// animation

var d6EnabledPhysicsSched *gogame.FunctionScheduler
var d6GopherDropperSched *gogame.FunctionScheduler

// init assets for demo 6
func initDemo6Assets() error {
	fmt.Printf("Loading Demo6 assets\n")

	var err error

	for _, imageAsset := range demo6Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create layer
	d6Layer0, err = renderController.CreateSpriteLayer(0, sdl.Point{0, 0})
	if err != nil {
		return err

	}

	d6GopherCount = 0

	// create sprites
	d6MainFloorSprite = &gogame.Sprite{Id: "main-floor-sprite", ImageAssetId: D6_FLOOR_IMAGE, Pos: sdl.Point{512, 750}, Width: 1024, Height: 34, Rotation: 0.0, Visible: true}

	d6FloorSprite1 = &gogame.Sprite{Id: "floor-sprite-1", ImageAssetId: D6_FLOOR_SHORT_IMAGE, Pos: sdl.Point{512 - 80, 275}, Width: 132, Height: 21, Rotation: 25.0, Visible: true}
	d6FloorSprite2 = &gogame.Sprite{Id: "floor-sprite-2", ImageAssetId: D6_FLOOR_SHORT_IMAGE, Pos: sdl.Point{512 + 80, 350}, Width: 132, Height: 21, Rotation: -25.0, Visible: true}
	d6FloorSprite3 = &gogame.Sprite{Id: "floor-sprite-3", ImageAssetId: D6_FLOOR_SHORT_IMAGE, Pos: sdl.Point{512 - 80, 425}, Width: 132, Height: 21, Rotation: 25.0, Visible: true}
	d6FloorSprite4 = &gogame.Sprite{Id: "floor-sprite-4", ImageAssetId: D6_FLOOR_SHORT_IMAGE, Pos: sdl.Point{512 + 80, 500}, Width: 132, Height: 21, Rotation: -25.0, Visible: true}
	d6FloorSprite5 = &gogame.Sprite{Id: "floor-sprite-5", ImageAssetId: D6_FLOOR_SHORT_IMAGE, Pos: sdl.Point{512 - 80, 575}, Width: 132, Height: 21, Rotation: 25.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(d6MainFloorSprite)
	assetHandler.AddSprite(d6FloorSprite1)
	assetHandler.AddSprite(d6FloorSprite2)
	assetHandler.AddSprite(d6FloorSprite3)
	assetHandler.AddSprite(d6FloorSprite4)
	assetHandler.AddSprite(d6FloorSprite5)

	// add to layer
	d6Layer0.AddSpriteToLayer(d6MainFloorSprite.Id)
	d6Layer0.AddSpriteToLayer(d6FloorSprite1.Id)
	d6Layer0.AddSpriteToLayer(d6FloorSprite2.Id)
	d6Layer0.AddSpriteToLayer(d6FloorSprite3.Id)
	d6Layer0.AddSpriteToLayer(d6FloorSprite4.Id)
	d6Layer0.AddSpriteToLayer(d6FloorSprite5.Id)

	startDemo6Animation()

	return nil
}

// render screen for demo 6
func demo6RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderLayers()
}

func unloadDemo6Assets() error {
	fmt.Printf("Unloading Demo6 assets\n")

	stopDemo6Animation()

	renderController.DestroySpriteLayer(0)

	for _, imageAsset := range demo6Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}



	return nil
}

func startDemo6Animation() {

	// init animation vars

	renderController.SetDebugInfo(false)
	d6EnabledPhysicsSched = gogame.NewFunctionScheduler(time.Duration(2*time.Second), 1, func() { demo6EnablePhysics() })
	d6GopherDropperSched = gogame.NewFunctionScheduler(time.Duration(2*time.Second), 100, demo6DropGopher)
	d6EnabledPhysicsSched.Start()
}

func stopDemo6Animation() {

	d6EnabledPhysicsSched.Destroy()
	d6GopherDropperSched.Destroy()
}

func demo6EnablePhysics() {

	d6MainFloorSprite.EnablePhysics(gogame.ImmovableMass)

	d6FloorSprite1.EnablePhysics(gogame.ImmovableMass)
	d6FloorSprite2.EnablePhysics(gogame.ImmovableMass)
	d6FloorSprite3.EnablePhysics(gogame.ImmovableMass)
	d6FloorSprite4.EnablePhysics(gogame.ImmovableMass)
	d6FloorSprite5.EnablePhysics(gogame.ImmovableMass)

	// start gopher dropper
	d6GopherDropperSched.Start()
}

func demo6DropGopher() {

	d6GopherCount++
	d6GopherSprite := &gogame.Sprite{Id: "gopher-sprite-" + strconv.Itoa(d6GopherCount), ImageAssetId: D6_GOPHER_RUN_IMAGE, Pos: sdl.Point{512 - 80, 200}, Width: 26, Height: 32, Rotation: 0.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(d6GopherSprite)

	// add to layer
	d6Layer0.AddSpriteToLayer(d6GopherSprite.Id)

	d6GopherSprite.EnablePhysics(0.2)

}
