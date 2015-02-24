package main

import (
	"fmt"
	"time"

	"github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

var typingSpeed = time.Duration(100 * time.Millisecond)
var rotationSpeed = time.Duration(20 * time.Millisecond)
var currentLabel int

type textTyping struct {
	actualText string
	typedText  string
	hasCursor  bool
}

var textLabels []textTyping

var demo1Fonts = []gogame.FontAsset{}

const (
	// fonts
	TYPEWRITER_KEY_SOUND  = "typewriter-key"
	TYPEWRITER_BELL_SOUND = "typewriter-bell"
)

var demo1Audio = []gogame.AudioAsset{
	{BaseAsset: gogame.BaseAsset{Id: TYPEWRITER_KEY_SOUND, FilePath: "./demo_assets/audio/typewriter-key.wav"}},
	{BaseAsset: gogame.BaseAsset{Id: TYPEWRITER_BELL_SOUND, FilePath: "./demo_assets/audio/typewriter-bell.wav"}},
}

// animation
var typingAnimSched *gogame.FunctionScheduler
var rotatingTextAnimSched *gogame.FunctionScheduler

// init assets for demo 1
func initDemo1Assets() error {

	fmt.Printf("Loading Demo1 assets\n")

	for _, audioAsset := range demo1Audio {
		err := assetHandler.AddAudioAsset(audioAsset, true)
		if err != nil {
			return fmt.Errorf("Error occurred whilst adding a audio asset:%s", err)
		}

	}

	startDemo1Animation()

	return nil
}

func playKey() error {

	audioPlayer.PlayAudio(TYPEWRITER_KEY_SOUND, 0)

	return nil
}

func playBell() error {

	audioPlayer.PlayAudio(TYPEWRITER_BELL_SOUND, 0)

	return nil
}

func unloadDemo1Assets() error {

	fmt.Printf("Unloading Demo1 assets\n")

	for _, fontAsset := range demo1Fonts {
		err := fontAsset.Unload()
		if err != nil {
			return fmt.Errorf("Error occurred whilst unloading a font asset:%s", err)
		}

	}

	typingAnimSched.Destroy()
	rotatingTextAnimSched.Destroy()
	return nil
}

func startDemo1Animation() {

	// init animation vars
	angle = 0.0
	textLabels = make([]textTyping, 8)

	textLabels[0] = textTyping{actualText: "Horizontal alignment: LEFT", typedText: ""}
	textLabels[1] = textTyping{actualText: "Horizontal alignment: CENTER", typedText: ""}
	textLabels[2] = textTyping{actualText: "Horizontal alignment: ABS_CENTER", typedText: ""}
	textLabels[3] = textTyping{actualText: "Horizontal alignment: RIGHT", typedText: ""}
	textLabels[4] = textTyping{actualText: "Vertical alignment: ABS_MIDDLE", typedText: ""}
	textLabels[5] = textTyping{actualText: "Vertical alignment: TOP", typedText: ""}
	textLabels[6] = textTyping{actualText: "Vertical alignment: MIDDLE", typedText: ""}
	textLabels[7] = textTyping{actualText: "Vertical alignment: BOTTOM", typedText: ""}
	currentLabel = 0

	typingAnimSched = gogame.NewFunctionScheduler(typingSpeed, -1, demo1AnimateTyping)
	rotatingTextAnimSched = gogame.NewFunctionScheduler(rotationSpeed, 360, demo1AnimateRotation)

	typingAnimSched.Start()

}

// this code is called for each tick of the timer
func demo1AnimateTyping() {

	// update current label
	hasFinished := textLabels[currentLabel].update()
	if hasFinished {
		// move onto next label
		if currentLabel != len(textLabels)-1 {
			currentLabel++
			playBell()

		} else {
			// move onto rotation
			typingAnimSched.Destroy()

			rotatingTextAnimSched.Start()
		}

	} else {

		playKey()

	}

}

func demo1AnimateRotation() {

	angle = angle + 10.0
	if angle > 360 {
		angle = angle - 360.0
	}

	fmt.Printf("demo1rotatetext\n")

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

	// valign
	fontPos := sdl.Point{textX, 200}
	renderController.RenderText(SHARED_FONT_24, textLabels[0].String(), fontPos, angle, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 250}
	renderController.RenderText(SHARED_FONT_24, textLabels[1].String(), fontPos, angle, black, gogame.TOP, gogame.CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 300}
	renderController.RenderText(SHARED_FONT_24, textLabels[2].String(), fontPos, angle, black, gogame.TOP, gogame.ABS_CENTER)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 350}
	renderController.RenderText(SHARED_FONT_24, textLabels[3].String(), fontPos, angle, black, gogame.TOP, gogame.RIGHT)
	renderFontPos(fontPos)

	// halign
	fontPos = sdl.Point{textX, 400}
	renderController.RenderText(SHARED_FONT_24, textLabels[4].String(), fontPos, angle, black, gogame.ABS_MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 450}
	renderController.RenderText(SHARED_FONT_24, textLabels[5].String(), fontPos, angle, black, gogame.TOP, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 500}
	renderController.RenderText(SHARED_FONT_24, textLabels[6].String(), fontPos, angle, black, gogame.MIDDLE, gogame.LEFT)
	renderFontPos(fontPos)

	fontPos = sdl.Point{textX, 550}
	renderController.RenderText(SHARED_FONT_24, textLabels[7].String(), fontPos, angle, black, gogame.BOTTOM, gogame.LEFT)
	renderFontPos(fontPos)

}

func renderFontPos(pos sdl.Point) {
	renderController.GetRenderer().SetDrawColor(255, 0, 0, 255)
	rect := sdl.Rect{pos.X, pos.Y, 5, 5}
	renderController.GetRenderer().FillRect(&rect)

}
