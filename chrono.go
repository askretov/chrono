// Package chrono provides basic chronometric features to measure time elapsed in various cases
// Package usage is as simple as native clock timer in your phone but also has some useful features like cumulative measurements
package chrono

import (
	"sync"
	"time"
)

var (
	instance *Meter
	o        sync.Once
)

// NewMeter returns a new instance of Meter
func NewMeter() *Meter {
	return &Meter{}
}

// getInstance returns a global singleton Meter instance used for package's direct calls
func getInstance() *Meter {
	o.Do(func() {
		instance = NewMeter()
	})
	return instance
}

// Start starts the timer with given tag
// Timer will be reset and restarted if another measurement with the same tag is ongoing
func Start(tag interface{}) {
	getInstance().Start(tag)
}

// Stop returns total time elapsed since a timer started
// It also prints time elapsed by default
func Stop(tag interface{}, noPrint ...bool) time.Duration {
	return getInstance().Stop(tag, noPrint...)
}

// Lap returns time elapsed since the last Lap call or Start for given tag
// It also prints time elapsed by default
func Lap(tag interface{}, msg string, noPrint ...bool) time.Duration {
	return getInstance().Lap(tag, msg, noPrint...)
}

// Capture Prints a time f run took
// If cumulative is true then nothing prints at the end, you have to call StopCumulativeCapture to print accumulated time
func Capture(tag interface{}, cumulative bool, f func()) {
	getInstance().Capture(tag, cumulative, f)
}

// StopCumulativeCapture stops cumulative capture and returns accumulated elapsed time value
func StopCumulativeCapture(tag interface{}) time.Duration {
	return getInstance().StopCumulativeCapture(tag)
}
