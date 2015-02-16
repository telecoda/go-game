package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// fonts
	DEMO1_FONT_8  = "droidsans8"
	DEMO1_FONT_16 = "droidsans16"
	DEMO1_FONT_48 = "droidsans48"
)

var horizLeftTyped = ""
var horizLeftText = "Horizontal alignment: LEFT"
var typingSpeed = time.Duration(100 * time.Millisecond)
var currentLabel int

type textTyping struct {
	actualText string
	typedText  string
	hasCursor  bool
}

var textLabels []textTyping

var demo1QuitChan = make(chan bool, 0)
var demo1fonts = []gogame.FontAsset{}

// init assets for demo 1
func initDemo1Assets() error {

	fmt.Printf("Loading Demo1 assets\n")

	for _, fontAsset := range demo1fonts {
		err := assetHandler.AddFontAsset(fontAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a font asset:%s", err)
		}

	}

	startDemo1Animation()

	return nil
}

func unloadDemo1Assets() error {

	fmt.Printf("Unloading Demo1 assets\n")

	for _, fontRes := range demo1fonts {
		err := fontRes.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading a font asset:%s", err)
		}

	}

	stopDemo1Animation()
	return nil
}

func startDemo1Animation() {

	// init animation vars
	textLabels = make([]textTyping, 8)

	textLabels[0] = textTyping{actualText: "Horizontal alignment: LEFT", typedText: ""}
	textLabels[1] = textTyping{actualText: "Horizontal alignment: CENTER", typedText: ""}
	textLabels[2] = textTyping{actualText: "Horizontal alignment: ABS_CENTER", typedText: ""}
	textLabels[3] = textTyping{actualText: "Horizontal alignment: RIGHT", typedText: ""}
	textLabels[4] = textTyping{actualText: "Vertical alignment: TOP", typedText: ""}
	textLabels[5] = textTyping{actualText: "Vertical alignment: ABS_MIDDLE", typedText: ""}
	textLabels[6] = textTyping{actualText: "Vertical alignment: MIDDLE", typedText: ""}
	textLabels[7] = textTyping{actualText: "Vertical alignment: BOTTOM", typedText: ""}
	currentLabel = 0

	go demo1AnimationLoop()

}

func demo1AnimationLoop() {

	ticker := time.NewTicker(typingSpeed)

	for {

		// wait for ani tick
		select {
		case <-demo1QuitChan:
			ticker.Stop()
			return
		case <-ticker.C:
			demo1AnimationTick()
		}
	}

}

func stopDemo1Animation() {

	demo1QuitChan <- true
}

// this code is called for each tick of the timer
func demo1AnimationTick() {

	// update current label

	hasFinished := textLabels[currentLabel].update()
	if hasFinished {
		// move onto next label
		if currentLabel != len(textLabels)-1 {
			currentLabel++
		}
	}

}

func (t *textTyping) update() bool {
	if len(t.typedText) < len(t.actualText) {
		// append another character
		t.typedText = t.actualText[0 : len(t.typedText)+1]
		t.hasCursor = true
		return false
	}

	// update complete
	t.hasCursor = false
	return true
}

func (t *textTyping) String() string {

	if t.hasCursor {
		return t.typedText + "_"
	} else {
		return t.typedText
	}
}

// render screen for demo 1
func demo1RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()

	textX := int32(demoWidth / 2)
	//textY := int32(demoHeight / 2)

	// valign
	fontPos := sdl.Point{textX, 150}
	renderController.RenderText(SHARED_FONT_24, textLabels[0].String(), fontPos, 0.0, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 200}
	renderController.RenderText(SHARED_FONT_24, textLabels[1].String(), fontPos, 0.0, black, gogame.TOP, gogame.CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 250}
	renderController.RenderText(SHARED_FONT_24, textLabels[2].String(), fontPos, 0.0, black, gogame.TOP, gogame.ABS_CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 300}
	renderController.RenderText(SHARED_FONT_24, textLabels[3].String(), fontPos, 0.0, black, gogame.TOP, gogame.RIGHT)
	renderFontPos(fontPos)

	// halign
	fontPos = sdl.Point{textX, 350}
	renderController.RenderText(SHARED_FONT_24, textLabels[4].String(), fontPos, 0.0, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 400}
	renderController.RenderText(SHARED_FONT_24, textLabels[5].String(), fontPos, 0.0, black, gogame.ABS_MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 450}
	renderController.RenderText(SHARED_FONT_24, textLabels[6].String(), fontPos, 0.0, black, gogame.MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 550}
	renderController.RenderText(SHARED_FONT_24, textLabels[7].String(), fontPos, 0.0, black, gogame.BOTTOM, gogame.LEFT)
	renderFontPos(fontPos)

}

func renderFontPos(pos sdl.Point) {
	renderController.GetRenderer().SetDrawColor(255, 0, 0, 255)
	rect := sdl.Rect{pos.X, pos.Y, 5, 5}
	renderController.GetRenderer().FillRect(&rect)

}
