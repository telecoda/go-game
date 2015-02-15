package main

// init assets for demo 5
func initDemo5Assets() error {
	return nil
}

// render screen for demo 5
func demo5RenderCallback() {
	renderController.ClearScreen(demoScreen.Color)

	renderTitle()
}

func unloadDemo5Assets() error {
	return nil
}
