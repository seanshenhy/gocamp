package restrict

import (
	"sync"
	"time"
)

type slideWin struct {
	width int
	queue []time.Time
	mu    sync.Mutex
}

func New(num int) (sw *slideWin) {
	return &slideWin{
		width: num,
		queue: make([]time.Time, 0, num),
		mu:    sync.Mutex{},
	}
}

func (q *slideWin) Add() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) < q.width {
		q.queue = append(q.queue, time.Now())
		return true
	}
	first := q.queue[0]
	if time.Now().Sub(first) < 1*time.Second {
		return false
	}
	q.queue = append(q.queue[1:], time.Now())
	return true
}
