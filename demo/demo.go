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
var renderController gogame.RenderController
var eventHandler gogame.EventHandler

func main() {

	initDemoScreens()

	var err error

	// init assetHandler and renderController
	assetHandler, renderController, eventHandler, err = gogame.NewGame("go-game demo", demoWidth, demoHeight, nil, demoEventReceiver)
	if err != nil {
		fmt.Printf("Error creating game. Program exit.\n", err)
		return
	}
	defer gogame.Destroy()

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
	demoScreen = demoScreens[currentDemoScreenId]
	err := demoScreen.activate()
	if err != nil {
		return err
	}

	err = previousDemoScreen.UnloadAssets()
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

	renderController.SetCallback(d.RenderCallback)

	return nil
}

func initDemoScreens() {

	demoScreens = make(map[int]*DemoScreen)

	demoScreens[0] = &DemoScreen{Id: 0, Title: "Title screen", Color: lightGrey, InitAssets: initDemo0Assets, UnloadAssets: unloadDemo0Assets, RenderCallback: demo0RenderCallback}
	demoScreens[1] = &DemoScreen{Id: 1, Title: "Fonts:", Description: "Text can be rendered from fonts loaded into the AssetManager", Color: white, InitAssets: initDemo1Assets, UnloadAssets: unloadDemo1Assets, RenderCallback: demo1RenderCallback}
	demoScreens[2] = &DemoScreen{Id: 2, Title: "Textures:", Description: "Textures can be rendered from images loaded into the AssetManager", Color: lightGrey, InitAssets: initDemo2Assets, UnloadAssets: unloadDemo2Assets, RenderCallback: demo2RenderCallback}
	demoScreens[3] = &DemoScreen{Id: 3, Title: "Sprites:", Color: white, InitAssets: initDemo3Assets, UnloadAssets: unloadDemo3Assets, RenderCallback: demo3RenderCallback}
	demoScreens[4] = &DemoScreen{Id: 4, Title: "Sprite Layers:", Color: lightGrey, InitAssets: initDemo4Assets, UnloadAssets: unloadDemo4Assets, RenderCallback: demo4RenderCallback}
	demoScreens[5] = &DemoScreen{Id: 5, Title: "Audio:", Color: white, InitAssets: initDemo5Assets, UnloadAssets: unloadDemo5Assets, RenderCallback: demo5RenderCallback}
	demoScreens[6] = &DemoScreen{Id: 6, Title: "Physics:", Color: lightGrey, InitAssets: initDemo6Assets, UnloadAssets: unloadDemo6Assets, RenderCallback: demo6RenderCallback}
	demoScreens[7] = &DemoScreen{Id: 7, Title: "Utilities:", Color: white, InitAssets: initDemo7Assets, UnloadAssets: unloadDemo7Assets, RenderCallback: demo7RenderCallback}
	demoScreens[8] = &DemoScreen{Id: 8, Title: "Game:", Color: white, InitAssets: initDemo8Assets, UnloadAssets: unloadDemo8Assets, RenderCallback: demo8RenderCallback}
	demoScreens[9] = &DemoScreen{Id: 9, Title: "Credits:", Color: lightGrey, InitAssets: initDemo9Assets, UnloadAssets: unloadDemo9Assets, RenderCallback: demo9RenderCallback}

}
