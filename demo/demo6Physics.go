package main

// init assets for demo 6
func initDemo6Assets() error {
	return nil
}

// render screen for demo 6
func demo6RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo6Assets() error {
	return nil
}
