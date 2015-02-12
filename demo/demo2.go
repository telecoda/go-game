package main

// init assets for demo 2
func initDemo2Assets() error {
	return nil
}

// render screen for demo 2
func demo2RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)
	//renderController.RenderText(DROID_SANS_16, "Test text", sdl.Point{20, 60}, black)
	//renderController.RenderText(DROID_SANS_48, "Test text", sdl.Point{20, 80}, red)

}

func unloadDemo2Assets() error {
	return nil
}
