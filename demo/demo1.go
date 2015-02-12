package main

// init assets for demo 1
func initDemo1Assets() error {
	return nil
}

// render screen for demo 1
func demo1RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)
	//renderController.RenderText(DROID_SANS_16, "Test text", sdl.Point{20, 60}, black)
	//renderController.RenderText(DROID_SANS_48, "Test text", sdl.Point{20, 80}, red)

}

func unloadDemo1Assets() error {
	return nil
}
