package gin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/goroscope/engine"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := &engine.Trace{
			ID:      uuid.NewString(),
			Events:  make(chan engine.Event, 1000),
			Started: time.Now(),
		}
		ctx := context.WithValue(c.Request.Context(), engine.CtxKey{}, t)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		close(t.Events)
		go persistTrace(t)
	}
}

func persistTrace(t *engine.Trace) {
	traceFile := fmt.Sprintf("trace-%s.json", t.ID)
	f, _ := os.Create(traceFile)
	enc := json.NewEncoder(f)
	for e := range t.Events {
		_ = enc.Encode(e)
	}

	err := f.Close()
	if err != nil {
		fmt.Println("error closing ", traceFile, " | Err: ", err)
	}
}
