package restrict

import (
	"sync"
	"time"
)

type slideWin struct {
	cap   int
	unit  time.Duration
	queue []time.Time
	mu    sync.Mutex
}

func New(cap int, unit time.Duration) (sw *slideWin) {
	return &slideWin{
		cap:   cap,
		unit:  unit,
		queue: make([]time.Time, 0, cap),
		mu:    sync.Mutex{},
	}
}

func (q *slideWin) IsPass() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) < q.cap {
		q.queue = append(q.queue, time.Now())
		return true
	}
	first := q.queue[0]
	if time.Now().Sub(first) < q.unit {
		return false
	}
	q.queue = append(q.queue[1:], time.Now())
	return true
}
