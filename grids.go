package gogame

import sdl "github.com/veandco/go-sdl2/sdl"

func RenderGrid(xSize, ySize int, color sdl.Color) {

	game.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	// draw surround
	game.Renderer.DrawLine(0, 0, game.width-1, 0)
	game.Renderer.DrawLine(game.width-1, 0, game.width-1, game.height-1)
	game.Renderer.DrawLine(game.width-1, game.height-1, 0, game.height-1)
	game.Renderer.DrawLine(0, game.height-1, 0, 0)

	// draw vertical lines
	for x := 0; x < game.width; x += xSize {

		game.Renderer.DrawLine(x, 0, x, game.height-1)

	}
	// draw horizontal lines
	for y := 0; y < game.height; y += ySize {

		game.Renderer.DrawLine(0, y, game.width-1, y)

	}
}
