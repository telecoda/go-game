package main

// init assets for demo 8
func initDemo8Assets() error {
	return nil
}

// render screen for demo 8
func demo8RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo8Assets() error {
	return nil
}
