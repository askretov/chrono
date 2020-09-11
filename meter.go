package chrono

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Meter struct {
	// Regular timers' start time values
	timers sync.Map
	// Lap start time values (time.Time)
	lapTimers sync.Map
	// Accumulated time within cumulative capture sessions (time.Duration)
	accumulatedTime sync.Map
}

// Start starts the timer with given tag
// Timer will be reset and restarted if another measurement with the same tag is ongoing
func (m *Meter) Start(tag interface{}) {
	m.timers.Store(tag, time.Now())
	m.lapTimers.Store(tag, time.Now())
}

// Stop returns total time elapsed since a timer started
// It also prints time elapsed by default
func (m *Meter) Stop(tag interface{}, noPrint ...bool) time.Duration {
	timeStop := time.Now()
	timeStart, ok := m.timers.Load(tag)
	if !ok {
		fmt.Printf("chrono > no running timer found with tag %v\n", tag)
		return 0
	}
	elapsed := timeStop.Sub(timeStart.(time.Time))
	if len(noPrint) == 0 || !noPrint[0] {
		fmt.Printf("chrono > tag: %v, time elapsed: %s\n", tag, elapsed)
	}
	// Delete timers
	m.timers.Delete(tag)
	m.lapTimers.Delete(tag)
	return elapsed
}

// Lap returns time elapsed since the last Lap call or Start for given tag
// It also prints time elapsed by default
func (m *Meter) Lap(tag interface{}, msg string, noPrint ...bool) time.Duration {
	timeStop := time.Now()
	timeStart, ok := m.lapTimers.Load(tag)
	if !ok {
		fmt.Printf("chrono > no running lap timer found with tag %v\n", tag)
		return 0
	}
	elapsed := timeStop.Sub(timeStart.(time.Time))
	if len(noPrint) == 0 || !noPrint[0] {
		fmt.Printf("chrono > tag: %v, msg: %s, time elapsed: %s\n", tag, msg, elapsed)
	}
	// Reset a lap timer
	m.lapTimers.Store(tag, time.Now())
	return elapsed
}

// Capture Prints a time f run took
// If cumulative is true then nothing prints at the end, you have to call StopCumulativeCapture to print accumulated time
func (m *Meter) Capture(tag interface{}, cumulative bool, f func()) {
	// Generate a random timer key
	timerTag := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	Start(timerTag)
	f()
	elapsed := Stop(timerTag, true)
	// Check if cumulative
	if cumulative {
		var accumulated time.Duration
		// Get current value
		if cv, ok := m.accumulatedTime.Load(tag); ok {
			accumulated += cv.(time.Duration)
		}
		accumulated += elapsed
		m.accumulatedTime.Store(tag, accumulated)
	} else {
		fmt.Printf("chrono > tag: %v, time elapsed: %s\n", tag, elapsed)
	}
}

// StopCumulativeCapture stops cumulative capture and returns accumulated elapsed time value
func (m *Meter) StopCumulativeCapture(tag interface{}) time.Duration {
	cv, ok := m.accumulatedTime.Load(tag)
	if !ok {
		fmt.Printf("chrono > no running cumulative capture found with tag %v\n", tag)
		return 0
	}
	fmt.Printf("chrono > tag: %v, accumulated time elapsed: %s\n", tag, cv.(time.Duration))
	m.accumulatedTime.Delete(tag)
	return cv.(time.Duration)
}
