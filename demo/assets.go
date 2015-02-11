package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/telecoda/go-game"
	sdl "github.com/veandco/go-sdl2/sdl"
)

// This code sets up the assets for the various demo screens
const (
	// fonts
	DROID_SANS_8  = "droidsans8"
	DROID_SANS_16 = "droidsans16"
	DROID_SANS_48 = "droidsans48"

	// images
	GOPHER_RUN   = "gopherrun"
	JAVA_DUKE    = "javaduke"
	FLOOR_IMAGE  = "floorimage"
	CLOUD1_IMAGE = "cloud1image"
	CLOUD2_IMAGE = "cloud2image"
	CLOUD3_IMAGE = "cloud3image"
	CLOUD4_IMAGE = "cloud4image"

	// sprites
	GOPHER_RUN_SPRITE = "gopherrunsprite"
	JAVA_DUKE_SPRITE  = "javadukesprite"
	FLOOR_SPRITE      = "floorsprite"
	CLOUD1_SPRITE     = "cloud1sprite"
	CLOUD2_SPRITE     = "cloud2sprite"
	CLOUD3_SPRITE     = "cloud3sprite"
	CLOUD4_SPRITE     = "cloud4sprite"
)

var fonts = []gogame.FontAsset{
	{Id: DROID_SANS_8, FilePath: "./assets/fonts/droid-sans/DroidSans.ttf", Size: 8},
	{Id: DROID_SANS_16, FilePath: "./assets/fonts/droid-sans/DroidSans.ttf", Size: 16},
	{Id: DROID_SANS_48, FilePath: "./assets/fonts/droid-sans/DroidSans.ttf", Size: 48},
}

var images = []gogame.ImageAsset{
	{Id: GOPHER_RUN, FilePath: "./assets/images/sprites/gopher-run.png"},
	{Id: JAVA_DUKE, FilePath: "./assets/images/sprites/java-duke.png"},
	{Id: FLOOR_IMAGE, FilePath: "./assets/images/sprites/floor.png"},
	{Id: CLOUD1_IMAGE, FilePath: "./assets/images/sprites/cloud1.png"},
	{Id: CLOUD2_IMAGE, FilePath: "./assets/images/sprites/cloud2.png"},
	{Id: CLOUD3_IMAGE, FilePath: "./assets/images/sprites/cloud3.png"},
	{Id: CLOUD4_IMAGE, FilePath: "./assets/images/sprites/cloud4.png"},
}

var cloudLayerId = 1
var gameLayerId = 0

var black = sdl.Color{R: 0, G: 0, B: 0, A: 255}
var red = sdl.Color{R: 255, G: 0, B: 0, A: 255}
var white = sdl.Color{R: 255, G: 255, B: 255, A: 255}
var lightGrey = sdl.Color{R: 222, G: 222, B: 222, A: 255}
var darkGrey = sdl.Color{R: 128, G: 128, B: 128, A: 255}

var angle = 0.0
var gameLayer *gogame.SpriteLayer
var cloudLayer *gogame.SpriteLayer

func initGameAssets() error {

	err := initFonts()
	if err != nil {
		return err
	}

	err = initImages()
	if err != nil {
		return err
	}

	err = initAudio()
	if err != nil {
		return err
	}

	err = initSprites()
	if err != nil {
		return err
	}

	return nil
}

func initAudio() error {

	return nil
}

func initFonts() error {

	for _, fontRes := range fonts {
		err := assetHandler.AddFontAsset(fontRes)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	return nil

}

func initImages() error {
	for _, imageRes := range images {
		err := assetHandler.AddImageAsset(imageRes)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding an image asset:%s", err)
		}

	}

	return nil
}

func initSprites() error {

	var err error

	gameLayer, err = renderController.CreateSpriteLayer(gameLayerId, sdl.Point{0, 0})
	if err != nil {
		return err
	}

	// init gophers
	id := 0
	for x := 0; x < 1024; x += 32 {
		spriteId := fmt.Sprintf("%s:%d", GOPHER_RUN_SPRITE, id)
		id += 1
		gopherSprite := gogame.Sprite{Id: spriteId, ImageAssetId: GOPHER_RUN, Pos: sdl.Point{int32(x), 100}, Width: 32, Height: 32, Rotation: 25.0, Visible: true}

		err := assetHandler.AddSprite(gopherSprite.Id, &gopherSprite)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a sprite asset:%s", err)
		}
		gopherSprite.EnablePhysics(10.0)

		// add them to layer 0

		gameLayer.AddSpriteToLayer(spriteId)
		if err != nil {
			return fmt.Errorf("Error adding sprite:%s to layer:%d", spriteId, gameLayerId)
		}

	}

	cloudLayer, err = renderController.CreateSpriteLayer(cloudLayerId, sdl.Point{0, 0})
	if err != nil {
		return err
	}

	// init clouds
	id = 0
	for x := 0; x < 10; x++ {
		// cloud1
		cloud1Id := fmt.Sprintf("%s:%d", CLOUD1_SPRITE, id)
		id += 1
		x := int32(rand.Int() % gameWidth)
		y := int32(rand.Int() % gameHeight)

		cloud1Sprite := gogame.Sprite{Id: cloud1Id, ImageAssetId: CLOUD1_IMAGE, Pos: sdl.Point{x, y}, Width: 200, Height: 68, Rotation: 0.0, Visible: true}

		err := assetHandler.AddSprite(cloud1Sprite.Id, &cloud1Sprite)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a sprite asset:%s", err)
		}

		// add them to cloud

		cloudLayer.AddSpriteToLayer(cloud1Id)
		if err != nil {
			return fmt.Errorf("Error adding sprite:%s to layer:%d", cloud1Id, cloudLayer)
		}
		// cloud 2
		cloud2Id := fmt.Sprintf("%s:%d", CLOUD2_SPRITE, id)
		id += 1
		x = int32(rand.Int() % gameWidth)
		y = int32(rand.Int() % gameHeight)

		cloud2Sprite := gogame.Sprite{Id: cloud2Id, ImageAssetId: CLOUD2_IMAGE, Pos: sdl.Point{x, y}, Width: 200, Height: 103, Rotation: 0.0, Visible: true}

		err = assetHandler.AddSprite(cloud2Sprite.Id, &cloud2Sprite)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a sprite asset:%s", err)
		}

		// add them to cloud

		cloudLayer.AddSpriteToLayer(cloud2Id)
		if err != nil {
			return fmt.Errorf("Error adding sprite:%s to layer:%d", cloud2Id, cloudLayer)
		}

	}

	// init duke
	dukeSprite := gogame.Sprite{Id: JAVA_DUKE_SPRITE, ImageAssetId: JAVA_DUKE, Pos: sdl.Point{512, 200}, Width: 100, Height: 100, Rotation: 5.0, Visible: true}

	err = assetHandler.AddSprite(dukeSprite.Id, &dukeSprite)
	if err != nil {
		return fmt.Errorf("Error occurred whilst adding a sprite asset:%s", err)
	}

	dukeSprite.EnablePhysics(100.0)

	// init floor
	floorSprite := gogame.Sprite{Id: FLOOR_SPRITE, ImageAssetId: FLOOR_IMAGE, Pos: sdl.Point{10, 600}, Width: 1004, Height: 32, Rotation: 0.0, Visible: true}

	err = assetHandler.AddSprite(floorSprite.Id, &floorSprite)
	if err != nil {
		return fmt.Errorf("Error occurred whilst adding a sprite asset:%s", err)
	}

	floorSprite.EnablePhysics(math.MaxFloat64)

	return nil
}
