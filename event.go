package cors

import "net/http"

const (
  corsRequestEvent = string(iota)
  preflightRequestEvent
  requestEvent
  simpleRequestEvent
)

type event struct {
  c                  Config
  w                  http.ResponseWriter
  r                  *http.Request
  propagationStopped bool
  requestTerminated  bool
}

func newEvent(c Config, w http.ResponseWriter, r *http.Request) *event {
  e := new(event)
  e.c = c
  e.w = w
  e.r = r

  return e
}

func (e *event) stopPropagation() {
  e.propagationStopped = true
}

func (e *event) isPropagationStopped() bool {
  return e.propagationStopped
}

func (e *event) terminateRequest() {
  e.requestTerminated = true
}

func (e *event) isRequestTerminated() bool {
  return e.requestTerminated
}
