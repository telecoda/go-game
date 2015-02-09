package gogame

import (
	"unsafe"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func clearScreen(color sdl.Color) {
	rendCont.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	rendCont.Renderer.Clear()
}

func readPixels(rect *sdl.Rect, pixels unsafe.Pointer) error {

	format := rendCont.Window.GetPixelFormat()
	//surface := rendCont.Window.GetSurface()

	/*if surface != nil {
		bpp := surface.BytesPerPixel()
		fmt.Printf("Bytes per pixel:%d\n", bpp)
	} else {
		fmt.Printf("No surface\n")
	}*/
	//fmt.Printf("Pixel format:%s\n", sdl.GetPixelFormatName(uint(format)))

	pitch := int(rect.W)
	return rendCont.Renderer.ReadPixels(rect, format, pixels, pitch)
}

func getPixelColorAt(point sdl.Point) (*sdl.Color, error) {

	rect := sdl.Rect{0, 0, int32(rendCont.width), int32(rendCont.height)}

	buf := make([]uint32, rendCont.width*rendCont.height)
	pixelPtr := unsafe.Pointer(&buf[0])

	err := readPixels(&rect, pixelPtr)
	if err != nil {
		return nil, err
	}

	index := point.Y*rect.W + point.X

	pixel := buf[index]

	a := uint8(pixel >> 24)
	r := uint8(pixel >> 16 & 0xFF)
	g := uint8(pixel >> 8 & 0xFF)
	b := uint8(pixel & 0xFF)

	return &sdl.Color{A: a, R: r, G: g, B: b}, nil
}
