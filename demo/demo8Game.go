package main

import (
	"fmt"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// images
	D8_GOPHER_RUN_IMAGE = "gopherrun"
	D8_FLOOR_IMAGE      = "floor"
)

var demo8Images = []gogame.ImageAsset{
	{BaseAsset: gogame.BaseAsset{Id: D8_GOPHER_RUN_IMAGE, FilePath: "./demo_assets/images/sprites/gopher-run.png"}},
	{BaseAsset: gogame.BaseAsset{Id: D8_FLOOR_IMAGE, FilePath: "./demo_assets/images/sprites/floor.png"}},
}

var d8Layer0 *gogame.SpriteLayer

var d8MainFloorSprite *gogame.Sprite

// init assets for demo 8
func initDemo8Assets() error {
	fmt.Printf("Loading Demo8 assets\n")

	var err error

	for _, imageAsset := range demo8Images {
		err := assetHandler.AddImageAsset(imageAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	// create layer
	d8Layer0, err = renderController.CreateSpriteLayer(0, sdl.Point{0, 0})
	if err != nil {
		return err

	}

	// create sprites
	d8MainFloorSprite = &gogame.Sprite{Id: "main-floor-sprite", ImageAssetId: D6_FLOOR_IMAGE, Pos: sdl.Point{512, 750}, Width: 1024, Height: 34, Rotation: 0.0, Visible: true}

	// add to assets
	assetHandler.AddSprite(d8MainFloorSprite)

	// add to layer
	d8Layer0.AddSpriteToLayer(d8MainFloorSprite.Id)

	startDemo8Animation()

	return nil

}

func startDemo8Animation() {

	// init animation vars
	renderController.SetDebugInfo(false)
	//d6EnabledPhysicsSched = gogame.NewFunctionScheduler(time.Duration(2*time.Second), 1, func() { demo6EnablePhysics() })
	//d6GopherDropperSched = gogame.NewFunctionScheduler(time.Duration(2*time.Second), 100, demo6DropGopher)
	//d6EnabledPhysicsSched.Start()

}

// render screen for demo 8
func demo8RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	renderController.RenderLayers()

}

func unloadDemo8Assets() error {
	fmt.Printf("Unloading Demo8 assets\n")

	stopDemo8Animation()

	err := renderController.DestroySpriteLayer(0)
	if err != nil {
		fmt.Printf("Error: failed to destroy sprite layer: %d Error: %s", 0, err)
		return err
	}

	for _, imageAsset := range demo8Images {
		err := imageAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading an image asset:%s", err)
		}

	}

	// remove physics bodies
	renderController.ClearWorld()

	return nil
}

func stopDemo8Animation() {

	//d6EnabledPhysicsSched.Destroy()
	//d6GopherDropperSched.Destroy()
}

func demo8EnablePhysics() {

	d8MainFloorSprite.EnablePhysics(gogame.ImmovableMass)

	//d6FloorSprite1.EnablePhysics(gogame.ImmovableMass)
	//d6FloorSprite2.EnablePhysics(gogame.ImmovableMass)
	//d6FloorSprite3.EnablePhysics(gogame.ImmovableMass)
	//d6FloorSprite4.EnablePhysics(gogame.ImmovableMass)
	//d6FloorSprite5.EnablePhysics(gogame.ImmovableMass)

	// start gopher dropper
	//d6GopherDropperSched.Start()
}
