package limit

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// maxWait is how long a test waits for RunLimited before it gives up.
const maxWait = 300 * time.Millisecond

// probe hands out tasks that measure how many of them run at once. Each task
// waits at a gate until full tasks are running together, so a RunLimited
// that respects its limit reaches it exactly, without sleeping.
type probe struct {
	full    int32
	running atomic.Int32
	peak    atomic.Int32
	ran     atomic.Int32
	gate    chan struct{}
	opening sync.Once
}

func newProbe(full int) *probe {
	return &probe{full: int32(full), gate: make(chan struct{})}
}

func (p *probe) task() {
	now := p.running.Add(1)
	p.raisePeak(now)
	if now >= p.full {
		p.opening.Do(func() { close(p.gate) })
	}
	<-p.gate
	p.running.Add(-1)
	p.ran.Add(1)
}

// raisePeak stores now as the peak if it is higher. Another task may change
// the peak between the load and the swap; then it tries again.
func (p *probe) raisePeak(now int32) {
	for {
		old := p.peak.Load()
		if now <= old || p.peak.CompareAndSwap(old, now) {
			return
		}
	}
}

func TestRunLimited(t *testing.T) {
	tests := []struct {
		name  string
		tasks int
		limit int
	}{
		{name: "one at a time", tasks: 5, limit: 1},
		{name: "two at a time", tasks: 6, limit: 2},
		{name: "three at a time", tasks: 20, limit: 3},
		{name: "a limit above the number of tasks", tasks: 3, limit: 10},
		{name: "no tasks", tasks: 0, limit: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newProbe(min(tt.tasks, tt.limit))
			tasks := make([]func(), tt.tasks)
			for i := range tasks {
				tasks[i] = p.task
			}

			finished := make(chan struct{})
			go func() {
				RunLimited(tasks, tt.limit)
				close(finished)
			}()
			select {
			case <-finished:
			case <-time.After(maxWait):
				t.Fatalf("RunLimited with %d tasks and limit %d does not return; does it let %d tasks run at once?",
					tt.tasks, tt.limit, min(tt.tasks, tt.limit))
			}

			if ran := int(p.ran.Load()); ran != tt.tasks {
				t.Errorf("RunLimited with %d tasks returned after %d finished, want it to wait for all of them",
					tt.tasks, ran)
			}
			if peak := int(p.peak.Load()); peak > tt.limit {
				t.Errorf("RunLimited with limit %d ran %d tasks at once", tt.limit, peak)
			}
		})
	}
}
