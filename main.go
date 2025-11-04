package goroscope

import (
	"context"
	"fmt"
	"github.com/dinson/goroscope/engine"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Goroscope active!")
}

// Go wraps goroutine function "fn" and accepts a "name".
// Giving "name" helps in identifying goroutines easily.
func Go(ctx context.Context, name string, fn func()) {
	t := FromContext(ctx)
	if t == nil {
		go fn()
		return
	}

	parent := curGID()
	gid := newGID()

	t.Events <- engine.Event{
		Goroutine: gid,
		Parent:    parent,
		Name:      name,
		Action:    "start",
		Time:      time.Now(),
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Events <- engine.Event{
					Goroutine: gid,
					Action:    "panic",
					Name:      name,
					Time:      time.Now(),
				}
				panic(r)
			}
			t.Events <- engine.Event{
				Goroutine: gid,
				Action:    "done",
				Name:      name,
				Time:      time.Now(),
			}
		}()
		fn()
	}()
}

func FromContext(ctx context.Context) *engine.Trace {
	if v := ctx.Value(engine.CtxKey{}); v != nil {
		return v.(*engine.Trace)
	}
	return nil
}

func curGID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	s := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.ParseInt(s, 10, 64)
	return id
}

func newGID() int64 {
	return time.Now().UnixNano() // pseudo-ID for child
}
