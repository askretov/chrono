package chrono

import (
	"testing"
)

func Test_Chrono(t *testing.T) {
	// Just a usage example
	// Basic
	Start("basic")
	for i := 0; i < 1000000000; i++ {
	}
	Stop("basic")

	// Laps
	Start("laps")
	for i := 0; i < 1000000000; i++ {
	}
	Lap("laps", "lap 1")
	for i := 0; i < 1000000000; i++ {
	}
	Lap("laps", "lap 2")
	for i := 0; i < 1000000000; i++ {
	}
	Lap("laps", "lap 3")
	Stop("laps")

	// Capture
	Capture("capture", false, func() {
		for i := 0; i < 1000000000; i++ {
		}
	})
}
