package cors

import "net/http"

const (
  CorsRequestEvent = string(iota)
  PreflightRequestEvent
  RequestEvent
  SimpleRequestEvent
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

func (e *Event) StopPropagation() {
  e.propagationStopped = true
}

func (e *Event) IsPropagationStopped() bool {
  return e.propagationStopped
}

func (e *Event) TerminateRequest() {
  e.requestTerminated = true
}

func (e *Event) IsRequestTerminated() bool {
  return e.requestTerminated
}
