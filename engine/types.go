package engine

import "time"

type CtxKey struct{}

type Trace struct {
	ID      string
	Events  chan Event
	Started time.Time
}

type Event struct {
	Goroutine int64
	Parent    int64
	Name      string
	Action    string // start, done, panic
	Time      time.Time
}
