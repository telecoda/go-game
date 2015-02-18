package gogame

import sdl "github.com/veandco/go-sdl2/sdl"

func (r renderController) RenderGrid(xSize, ySize int, color sdl.Color) {

	rect := sdl.Rect{X: 0, Y: 0, W: int32(r.width - 1), H: int32(r.height - 1)}

	r.RenderGridInRect(rect, xSize, ySize, color)

}

func (r renderController) RenderGridInRect(rect sdl.Rect, xSize, ySize int, color sdl.Color) {

	r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	x1 := int(rect.X)
	x2 := int(rect.X + rect.W - 1)
	y1 := int(rect.Y)
	y2 := int(rect.Y + rect.H - 1)
	// draw surround
	r.Renderer.DrawLine(x1, y1, x2, y1)
	r.Renderer.DrawLine(x2, y1, x2, y2)
	r.Renderer.DrawLine(x2, y2, x1, y2)
	r.Renderer.DrawLine(x1, y2, x1, y1)

	// draw vertical lines
	for x := x1; x < x2+1; x += xSize {

		r.Renderer.DrawLine(x, y1, x, y2)

	}
	// draw horizontal lines
	for y := y1; y < y2+1; y += ySize {

		r.Renderer.DrawLine(x1, y, x2, y)

	}

}
