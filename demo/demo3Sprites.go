package main

// init assets for demo 3
func initDemo3Assets() error {
	return nil
}

// render screen for demo 3
func demo3RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo3Assets() error {
	return nil
}
