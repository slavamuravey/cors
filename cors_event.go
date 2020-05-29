package cors

import "net/http"

const (
	SimpleRequestEvent    string = "simple"
	PreflightRequestEvent string = "preflight"
)

type event struct {
	c                  Config
	w                  http.ResponseWriter
	r                  *http.Request
	next               http.Handler
	propagationStopped bool
	forwardedToNext    bool
}

func (e *event) stopPropagation() {
	e.propagationStopped = true
}

func (e *event) isPropagationStopped() bool {
	return e.propagationStopped
}

func (e *event) forwardToNext() {
	e.forwardedToNext = true
}

func (e *event) isForwardedToNext() bool {
	return e.forwardedToNext
}
