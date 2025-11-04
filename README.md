# 🔭 Goroscope

Visualize and understand your Go concurrency.

Goroscope traces goroutine lifecycles, fan-out/fan-in patterns, and blocking
behavior in real time, giving you a clear timeline of concurrency within each request.

Perfect for debugging, teaching, and understanding how Go's concurrency really works.

## Installation

<code>go get github.com/dinson/goroscope</code>

## Usage

1. Initialize as gin middleware

Example:

```go
import (
	goroscopeGin "github.com/dinson/goroscope/pkg/gin"
)

func RegisterHandler(r *gin.RouterGroup) {
    r.Use(goroscopeGin.Middleware)

	r.POST("/my-route", h.MyHandler)
}

----- or -----

router.POST("/my-route", goroscopeGin.Middleware(), h.MyHandler)
```

2. Wrap goroutines in goroscope

Example:

```go
import "github.com/dinson/goroscope"

func fn() {
	// your logic
	
        goroscope.Go(ctx, "<goroutine-name>", func() { myBackgroundFunction(ctx) })
	
	// rest of your logic
}
```

3. Each request will output the goroutine start and end times to a `trace-<uuid>.json` file.

Example:

```json lines
{"Goroutine":1762273314917752000,"Parent":169,"Name":"v1planner","Action":"start","Time":"2025-11-04T20:21:54.917752+04:00"}
{"Goroutine":1762273314917810000,"Parent":169,"Name":"v2planner","Action":"start","Time":"2025-11-04T20:21:54.91781+04:00"}
{"Goroutine":1762273314917865000,"Parent":169,"Name":"v3planner","Action":"start","Time":"2025-11-04T20:21:54.917866+04:00"}
{"Goroutine":1762273314917810000,"Parent":0,"Name":"v2planner","Action":"done","Time":"2025-11-04T20:22:29.185576+04:00"}
{"Goroutine":1762273314917752000,"Parent":0,"Name":"v1planner","Action":"done","Time":"2025-11-04T20:22:43.293273+04:00"}
{"Goroutine":1762273314917865000,"Parent":0,"Name":"v3planner","Action":"done","Time":"2025-11-04T20:22:47.23716+04:00"}
```