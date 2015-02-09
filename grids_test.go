package gogame

import (
	"reflect"
	"testing"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func TestRenderGrid(t *testing.T) {

	title := "RenderGrid Test"
	width := 100
	height := 100

	_, renderController, err := NewGame(title, width, height, nil)
	if err != nil {
		t.Fatalf("Error was not expected and got: %s", err)
	}

	renderColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	backgroundColor := sdl.Color{R: 1, G: 2, B: 3, A: 4}

	clearScreen(backgroundColor)
	renderController.RenderGrid(5, 5, renderColor)

	pixelColor, err := getPixelColorAt(sdl.Point{0, 0})

	if err != nil {
		t.Fatalf("Error was not expected and got: %s", err)
	}

	if !reflect.DeepEqual(renderColor, *pixelColor) {
		t.Fatalf("Error: expected colour:%v but got :%v", renderColor, pixelColor)
	}

}