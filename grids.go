package gogame

import sdl "github.com/veandco/go-sdl2/sdl"

func (r renderController) RenderGrid(xSize, ySize int, color sdl.Color) {

	r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	// draw surround
	r.Renderer.DrawLine(0, 0, r.width-1, 0)
	r.Renderer.DrawLine(r.width-1, 0, r.width-1, r.height-1)
	r.Renderer.DrawLine(r.width-1, r.height-1, 0, r.height-1)
	r.Renderer.DrawLine(0, r.height-1, 0, 0)

	// draw vertical lines
	for x := 0; x < r.width; x += xSize {

		r.Renderer.DrawLine(x, 0, x, r.height-1)

	}
	// draw horizontal lines
	for y := 0; y < r.height; y += ySize {

		r.Renderer.DrawLine(0, y, r.width-1, y)

	}
}
