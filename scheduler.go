package gogame

import "time"

type ScheduledFunction func()

type FunctionScheduler struct {
	frequency        time.Duration
	iterations       int
	currentCount     int
	paused           bool
	callbackFunction ScheduledFunction
	timer            *time.Ticker
	quitChannel      chan bool
}

type Scheduler interface {
	Start()
	Destroy()
	Pause()
	Unpause()
}

var INFINITY = -1

func NewFunctionScheduler(frequency time.Duration, iterations int, callbackFunc ScheduledFunction) *FunctionScheduler {
	fs := FunctionScheduler{
		frequency:        frequency,
		iterations:       iterations,
		currentCount:     0,
		paused:           false,
		callbackFunction: callbackFunc,
		quitChannel:      make(chan bool, 1),
	}

	return &fs
}

func (fs *FunctionScheduler) Start() {

	fs.timer = time.NewTicker(fs.frequency)

	go func() {
		for {

			// wait for tick
			select {
			case <-fs.quitChannel:
				return
			case <-fs.timer.C:
				fs.invokeFunction()
			}
		}

	}()

}

func (fs *FunctionScheduler) invokeFunction() {

	if fs.paused {
		return
	}

	fs.currentCount++

	if fs.iterations == INFINITY {
		fs.callbackFunction()
		return
	}

	if fs.currentCount > fs.iterations {
		fs.Destroy()
	} else {
		fs.callbackFunction()
	}
}

func (fs *FunctionScheduler) Pause() {
	fs.paused = true
}

func (fs *FunctionScheduler) Unpause() {
	fs.paused = false
}

func (fs *FunctionScheduler) Destroy() {

	fs.quitChannel <- true

}
