package main

// init assets for demo 2
func initDemo2Assets() error {
	return nil
}

// render screen for demo 2
func demo2RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo2Assets() error {
	return nil
}
