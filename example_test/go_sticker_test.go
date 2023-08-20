package example_test

import (
	"github.com/sinlov-go/go-sticker"
	"testing"
	"time"
)

func TestSTicker(t *testing.T) {
	// mock STicker
	tests := []struct {
		name     string
		duration time.Duration
		accuracy time.Duration
		maxCnt   int
		wantErr  error
	}{
		{
			name:     "sample",
			duration: time.Second * 3,
			accuracy: time.Second,
			maxCnt:   5,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do STicker
			ticker := go_sticker.New(tc.duration, tc.accuracy)
			cnt := 0
			for tick := range ticker.C {
				if cnt > tc.maxCnt {
					ticker.Stop()
					t.Logf("stop tc %s tick: %v", t.Name(), tick)
					break
				}
				// Process tick
				t.Logf("tc %s cnt %d tick: %v", t.Name(), cnt, tick)
				cnt++
			}
		})
	}
}
