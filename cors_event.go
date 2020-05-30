package cors

import "net/http"

const (
  SimpleRequestEvent = string(iota)
  PreflightRequestEvent
  CorsRequestEvent
  RequestEvent
)

type Event struct {
  C                  Config
  W                  http.ResponseWriter
  R                  *http.Request
  propagationStopped bool
  requestTerminated  bool
}

func NewEvent(c Config, w http.ResponseWriter, r *http.Request) *Event {
  e := new(Event)
  e.C = c
  e.W = w
  e.R = r

  return e
}

func (e *Event) stopPropagation() {
  e.propagationStopped = true
}

func (e *Event) isPropagationStopped() bool {
  return e.propagationStopped
}

func (e *Event) terminateRequest() {
  e.requestTerminated = true
}

func (e *Event) isRequestTerminated() bool {
  return e.requestTerminated
}
