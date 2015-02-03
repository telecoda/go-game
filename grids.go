package gogame

import sdl "github.com/veandco/go-sdl2/sdl"

func (g *Game) RenderGrid(xSize, ySize int, color sdl.Color) {

	g.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	// draw surround
	g.Renderer.DrawLine(0, 0, g.width-1, 0)
	g.Renderer.DrawLine(g.width-1, 0, g.width-1, g.height-1)
	g.Renderer.DrawLine(g.width-1, g.height-1, 0, g.height-1)
	g.Renderer.DrawLine(0, g.height-1, 0, 0)

	// draw vertical lines
	for x := 0; x < g.width; x += xSize {

		g.Renderer.DrawLine(x, 0, x, g.height-1)

	}
	// draw horizontal lines
	for y := 0; y < g.height; y += ySize {

		g.Renderer.DrawLine(0, y, g.width-1, y)

	}
}
