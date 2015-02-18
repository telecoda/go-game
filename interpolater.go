package gogame

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type ColorInterpolator struct {
	scheduler    FunctionScheduler
	fromColor    sdl.Color
	toColor      sdl.Color
	CurrentColor sdl.Color
	floatR       float32
	floatG       float32
	floatB       float32
	floatA       float32
}

// creates an interpolater that will iterate the current color, from the fromColor to the toColor
// within the defined time duration, eg. black to white in 5 seconds
func NewColorInterpolator(fromColor sdl.Color, toColor sdl.Color, duration time.Duration) (*ColorInterpolator, error) {

	if duration.Seconds() == 0 {
		return nil, fmt.Errorf("Error: interpolation duration must be greater than 0 seconds")
	}

	// calculate incremental factor based on duration

	rDiff := float32(toColor.R) - float32(fromColor.R)
	gDiff := float32(toColor.G) - float32(fromColor.G)
	bDiff := float32(toColor.B) - float32(fromColor.B)
	aDiff := float32(toColor.A) - float32(fromColor.A)

	durMilliSeconds := duration.Seconds() * 1000

	iterFrequency := time.Duration(33 * time.Millisecond) // 30 fps
	iterMilliSeconds := iterFrequency.Seconds() * 1000

	iterations := float32(durMilliSeconds / iterMilliSeconds)

	rInc := float32(rDiff / iterations)
	gInc := float32(gDiff / iterations)
	bInc := float32(bDiff / iterations)
	aInc := float32(aDiff / iterations)

	cInt := ColorInterpolator{
		fromColor:    fromColor,
		toColor:      toColor,
		CurrentColor: fromColor,
		floatR:       float32(fromColor.R),
		floatG:       float32(fromColor.G),
		floatB:       float32(fromColor.B),
		floatA:       float32(fromColor.A),
	}

	interpFunc := func() {
		cInt.floatR += rInc
		cInt.floatG += gInc
		cInt.floatB += bInc
		cInt.floatA += aInc
		cInt.CurrentColor = sdl.Color{R: uint8(cInt.floatR), G: uint8(cInt.floatG), B: uint8(cInt.floatB), A: uint8(cInt.floatA)}
	}

	fSched := NewFunctionScheduler(iterFrequency, int(iterations), interpFunc)

	cInt.scheduler = *fSched

	return &cInt, nil
}

func (ci *ColorInterpolator) Start() {
	ci.scheduler.Start()
}
