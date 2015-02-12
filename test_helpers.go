package gogame

import (
	"fmt"
	"runtime"
	"unsafe"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func readPixels(rect *sdl.Rect, pixels unsafe.Pointer) error {

	format := rendCont.Window.GetPixelFormat()

	fmt.Printf("Pixel format:%s\n", sdl.GetPixelFormatName(uint(format)))

	pitch := int(rect.W)
	return rendCont.Renderer.ReadPixels(rect, format, pixels, pitch)
}

func getSurfaceFromFont() (*sdl.Surface, error) {

	font, err := gameAssets.getFont(SYSTEM_FONT_ID)
	if err != nil {
		return nil, err
	}

	textColor := sdl.Color{R: 255, B: 0, G: 0, A: 255}

	textSurface := font.RenderText_Solid("test", textColor)

	return textSurface, err

}

func getPixelColorAt(point sdl.Point) (*sdl.Color, error) {

	format := rendCont.Window.GetPixelFormat()

	pixelFormat, err := sdl.AllocFormat(uint(format))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Pixel format:%v\n", pixelFormat)

	rect := sdl.Rect{0, 0, int32(rendCont.width), int32(rendCont.height)}

	buf := make([]uint32, rendCont.width*rendCont.height)
	pixelPtr := unsafe.Pointer(&buf[0])

	err = readPixels(&rect, pixelPtr)
	if err != nil {
		return nil, err
	}

	index := point.Y*rect.W*int32(pixelFormat.BytesPerPixel) + point.X*int32(pixelFormat.BytesPerPixel)

	pixel := buf[index]

	r, g, b, a := sdl.GetRGBA(pixel, pixelFormat)

	return &sdl.Color{A: a, R: r, G: g, B: b}, nil
}

// prints memory usage to the stdout
func ReportMemoryUsage() {

	fmt.Printf("Memory Usage\n")
	fmt.Printf("------------\n")
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	fmt.Printf("Memory allocated in use:%d\n", m.Alloc)
	fmt.Printf("Memory total allocated :%d\n", m.TotalAlloc)
	fmt.Printf("Memory heap            :%d\n", m.HeapAlloc)
	fmt.Printf("Memory next GC         :%d\n", m.NextGC)
}
