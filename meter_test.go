package chrono

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeter_Start(t *testing.T) {
	type args struct {
		tag interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{tag: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMeter()
			m.Start(tt.args.tag)
			_, ok1 := m.timers.Load(tt.args.tag)
			_, ok2 := m.lapTimers.Load(tt.args.tag)
			assert.True(t, ok1)
			assert.True(t, ok2)
		})
	}
}

func TestMeter_Stop(t *testing.T) {
	type args struct {
		tag     interface{}
		noPrint []bool
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "ok",
			args: args{tag: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMeter()
			m.Start(tt.args.tag)
			target := time.Now().Add(time.Second * 3)
			for time.Now().Before(target) {
			}
			elapsed := m.Stop(tt.args.tag, tt.args.noPrint...)
			// Check that meter returns elapsed time
			assert.GreaterOrEqual(t, elapsed.Nanoseconds(), (time.Second * 3).Nanoseconds())
			// Check that meter deletes unused timers after stop
			_, ok1 := m.timers.Load(tt.args.tag)
			_, ok2 := m.lapTimers.Load(tt.args.tag)
			assert.False(t, ok1 || ok2)

		})
	}
}

func TestMeter_Lap(t *testing.T) {
	type args struct {
		tag     interface{}
		msg     string
		noPrint []bool
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "ok",
			args: args{tag: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMeter()
			firstLap := time.Second * 3
			secondLap := time.Second * 2
			m.Start(tt.args.tag)
			target := time.Now().Add(firstLap)
			for time.Now().Before(target) {
			}
			// Start a new lap
			m.Lap(tt.args.tag, "")
			target = time.Now().Add(secondLap)
			for time.Now().Before(target) {
			}
			secondLapElapsed := m.Lap(tt.args.tag, "")
			assert.True(t, secondLapElapsed >= secondLap && secondLapElapsed < firstLap)
		})
	}
}

func TestMeter_Cumulative(t *testing.T) {
	type args struct {
		tag        interface{}
		cumulative bool
		f          func()
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "cumulative true",
			args: args{
				tag:        "test",
				cumulative: true,
				f: func() {
					target := time.Now().Add(time.Second * 2)
					for time.Now().Before(target) {
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMeter()
			// Capture for 3 times
			m.Capture(tt.args.tag, tt.args.cumulative, tt.args.f)
			m.Capture(tt.args.tag, tt.args.cumulative, tt.args.f)
			m.Capture(tt.args.tag, tt.args.cumulative, tt.args.f)
			elapsed := m.StopCumulativeCapture(tt.args.tag)
			// Check elapsed time
			assert.True(t, elapsed >= time.Second*2*3)
			// Check garbage collection
			m.timers.Range(func(key, value interface{}) bool {
				return assert.Fail(t, "Timers not deleted")
			})
			m.lapTimers.Range(func(key, value interface{}) bool {
				return assert.Fail(t, "Timers not deleted")
			})
			m.accumulatedTime.Range(func(key, value interface{}) bool {
				return assert.Fail(t, "Timers not deleted")
			})
		})
	}
}
