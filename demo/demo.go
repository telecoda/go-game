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
		fmt.Printf("Error creating game:%s. Program exit.\n", err)
		return
	}
	defer gogame.Destroy()

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

	demoScreens[0] = &DemoScreen{Id: 0, Title: "Title screen", Color: sdl.Color{R: 255, G: 255, B: 255, A: 255}, InitAssets: initDemo0Assets, UnloadAssets: unloadDemo0Assets, RenderCallback: demo0RenderCallback}
	demoScreens[1] = &DemoScreen{Id: 1, Title: "Text demo", Color: sdl.Color{R: 255, G: 0, B: 0, A: 255}, InitAssets: initDemo1Assets, UnloadAssets: unloadDemo1Assets, RenderCallback: demo1RenderCallback}
	demoScreens[2] = &DemoScreen{Id: 2, Title: "Credits screen", Color: sdl.Color{R: 128, G: 128, B: 128, A: 255}, InitAssets: initDemo2Assets, UnloadAssets: unloadDemo2Assets, RenderCallback: demo2RenderCallback}

}
