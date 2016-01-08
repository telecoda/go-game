package main

// author: @telecoda

/* demo - to show the various features of the go-game library
 */

import (
	"fmt"

	gogame "github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

var demoWidth = 1024
var demoHeight = 800
var cx = int32(demoWidth / 2)
var cy = int32(demoHeight / 2)

var currentDemoScreenId int

var demoScreens map[int]*DemoScreen
var demoScreen *DemoScreen

var assetHandler gogame.AssetManager
var audioPlayer gogame.AudioPlayer

var engine *gogame.Engine
var renderer gogame.Renderer

func main() {

	initDemoScreens()

	var err error

	//assetHandler, renderController, audioPlayer, eventHandler, err = gogame.NewGame("go-game demo", demoWidth, demoHeight, nil, demoEventReceiver)
	engine, err = gogame.NewGame("go-game demo", demoWidth, demoHeight, nil, demoEventReceiver)
	if err != nil {
		fmt.Printf("Error creating game. Program exit.\n", err)
		return
	}
	defer gogame.Destroy()

	renderer = engine.GetRenderer()
	assetHandler = engine.AssetManager
	audioPlayer = engine.GetAudioPlayer()

	err = initSharedAssets()
	if err != nil {
		fmt.Printf("Error initialsing assets. Program exit.\n", err)
		return
	}

	// start with demo 1
	currentDemoScreenId = 0
	demoScreen = demoScreens[currentDemoScreenId]
	err = demoScreen.activate()

	if err != nil {
		fmt.Printf("Error activating screen:%s. Program exit.\n", err)
		return
	}
	// start main event loop
	gogame.EventLoop()

}

// receive events from game
func demoEventReceiver(e interface{}) {
	switch t := e.(type) {
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_SPACE:
			err := nextDemo()
			if err != nil {
				fmt.Printf("Error switching demos :%s", err)
			}
		}
	}
}

func nextDemo() error {
	nextDemoScreenId := currentDemoScreenId + 1

	if nextDemoScreenId > len(demoScreens)-1 {
		currentDemoScreenId = 0
	} else {
		currentDemoScreenId = nextDemoScreenId
	}

	previousDemoScreen := demoScreen

	err := previousDemoScreen.UnloadAssets()
	if err != nil {
		return err
	}

	demoScreen = demoScreens[currentDemoScreenId]
	err = demoScreen.activate()
	if err != nil {
		return err
	}

	gogame.ReportMemoryUsage()

	return nil

}

// make this screen the current one being rendered by the renderer
func (d DemoScreen) activate() error {
	err := demoScreen.InitAssets()
	if err != nil {
		return err
	}

	fmt.Printf("Now showing demo:%s\n", demoScreen.Title)

	renderer.SetCallback(d.RenderCallback)

	return nil
}

func initDemoScreens() {

	demoScreens = make(map[int]*DemoScreen)

	demoScreens[0] = &DemoScreen{Id: 0, Title: "Title screen", Color: lightGrey, InitAssets: initDemo0Assets, UnloadAssets: unloadDemo0Assets, RenderCallback: demo0RenderCallback}
	demoScreens[1] = &DemoScreen{Id: 1, Title: "Fonts:", Description: "Text can be rendered from fonts  loaded into the AssetManager",
		CodeSample: []string{
			"// add font",
			"fontAsset :=gogame.FontAsset{{BaseAsset: gogame.BaseAsset{Id: \"my_id\", FilePath: \"path_to_font.ttf\"}, Size: 8}",
			"assetHandler.AddFontAsset(fontAsset, true)",
			"",
			"// render text",
			"renderController.RenderText(\"my_font_id\", \"my text to render\", sdl.Point{X: 25, Y: 25}, 0.0, color, gogame.TOP, gogame.LEFT)",
		},
		Color: white, InitAssets: initDemo1Assets, UnloadAssets: unloadDemo1Assets, RenderCallback: demo1RenderCallback}

	demoScreens[2] = &DemoScreen{Id: 2, Title: "Textures:", Description: "Textures can be rendered from images loaded into the AssetManager. They are best for drawing static images such as backgrounds",
		CodeSample: []string{
			"// add image",
			"imageAsset :=gogame.ImageAsset{{BaseAsset: gogame.BaseAsset{Id: \"my_image_id\", FilePath: \"my_image.png\"}}",
			"assetHandler.AddFontAsset(fontAsset, true)",
			"",
			"// render textures",
			"renderController.RenderTexture(\"my_image_id\", sdl.Point{x, y}, sizeX, sizeY)",
			"renderController.RenderRotatedTexture(\"my_image_id\", sdl.Point{x, y}, degrees, sizeX, sizeY)",
		},
		Color: lightGrey, InitAssets: initDemo2Assets, UnloadAssets: unloadDemo2Assets, RenderCallback: demo2RenderCallback}

	demoScreens[3] = &DemoScreen{Id: 3, Title: "Sprites:", Description: "Sprites allow you greater control over rendering and can be used for in game objects.",
		CodeSample: []string{
			"// add sprite",
			"sprite = &gogame.Sprite{Id: \"my_sprite_id\", ImageAssetId: \"my_image_id\", Pos: sdl.Point{100, 200}, ",
			"                                                 Width: 32, Height: 32, Rotation: 0.0, Visible: true}",
			"assetHandler.AddSprite(sprite)",
			"",
			"// render sprites",
			"renderController.RenderSprite(\"my_sprite_id\")",
		},
		Color: white, InitAssets: initDemo3Assets, UnloadAssets: unloadDemo3Assets, RenderCallback: demo3RenderCallback}
	demoScreens[4] = &DemoScreen{Id: 4, Title: "Sprite Layers:", Description: "Sprites can be managed as a group by adding them to a SpriteLayer",
		CodeSample: []string{
			"layer, err := renderController.CreateSpriteLayer(layerId, sdl.Point{0, 0})",
			"// add sprite to layer",
			"sprite = &gogame.Sprite{Id: \"my_sprite_id\", ImageAssetId: \"my_image_id\", Pos: sdl.Point{100, 200}, ",
			"                                                 Width: 32, Height: 32, Rotation: 0.0, Visible: true}",
			"assetHandler.AddSprite(sprite)",
			"layer.AddSpriteToLayer(\"my_sprite_id\")",
			"// render ALL layers 0=nearest",
			"renderController.RenderLayers()",
		},
		Color: lightGrey, InitAssets: initDemo4Assets, UnloadAssets: unloadDemo4Assets, RenderCallback: demo4RenderCallback}
	demoScreens[5] = &DemoScreen{Id: 5, Title: "Audio:", Description: "Audio can be controlled by the AudioPlayer",
		CodeSample: []string{
			"// add some audio (use true to load the audio at the same time",
			"audioAsset := {BaseAsset: gogame.BaseAsset{Id: \"my_audio_id\", FilePath: \"./demo_assets/audio/typewriter-key.wav\"}}",
			"assetHandler.AddAudioAsset(audioAsset, true)",
			"",
			"// play it",
			"audioPlayer.PlayAudio(\"my_audio_id\", 0)",
		},
		Color: white, InitAssets: initDemo5Assets, UnloadAssets: unloadDemo5Assets, RenderCallback: demo5RenderCallback}
	demoScreens[6] = &DemoScreen{Id: 6, Title: "Physics:", Description: "By enabling physics on a sprite means the sprite is controlled by the physics engine",
		CodeSample: []string{
			"// enable physics (provide a mass for the sprite eg. 0.2kg's",
			"gopherSprite.EnablePhysics(0.2)",
			"",
			"// add an immovable physics sprite (stuff will bounce off it)",
			"floorSprite.EnablePhysics(gogame.ImmovableMass)",
		},
		Color: lightGrey, InitAssets: initDemo6Assets, UnloadAssets: unloadDemo6Assets, RenderCallback: demo6RenderCallback}
	demoScreens[7] = &DemoScreen{Id: 7, Title: "Physics Joints:", Description: "Physics enabled sprites can be linked together easily using joints",
		CodeSample: []string{
			"// enable physics ",
			"stillSprite.EnablePhysics(gogame.ImmovableMass)",
			"swingingSprite.EnablePhysics(0.2)",
			"",
			"// joints are positioned relative to the sprites they join",
			"xOffset := gopherWidth / 2",
			"yOffset := gopherHeight / 2",
			"stillSprite.JoinTo(swingingSprite, sdl.Point{int32(xOffset), int32(yOffset)})",
		},

		Color: white, InitAssets: initDemo7Assets, UnloadAssets: unloadDemo7Assets, RenderCallback: demo7RenderCallback}
	demoScreens[8] = &DemoScreen{Id: 8, Title: "Game:", Color: white, InitAssets: initDemo8Assets, UnloadAssets: unloadDemo8Assets, RenderCallback: demo8RenderCallback}
	demoScreens[9] = &DemoScreen{Id: 9, Title: "Credits:", Color: lightGrey, InitAssets: initDemo9Assets, UnloadAssets: unloadDemo9Assets, RenderCallback: demo9RenderCallback}

}
