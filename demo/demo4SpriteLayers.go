package main

// init assets for demo 4
func initDemo4Assets() error {
	return nil
}

// render screen for demo 4
func demo4RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo4Assets() error {
	return nil
}
