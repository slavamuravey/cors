package cors

import (
  "net/http"
)

type Config struct {
  ContinuousPreflight bool
  AllowOrigin         string
  AllowMethods        []string
  AllowHeaders        []string
  AllowCredentials    bool
  ExposedHeaders      []string
  MaxAge              int
}

func CreateHandlerFunc(c Config) func(http.Handler) http.HandlerFunc {
  return func(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  ed := new(eventDispatcher)
	  ed.listeners = map[string][]listener{}
	  ed.addListener(PreflightRequestEvent, nextForwarder(checkRequestIsCors))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyAllowOrigin))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyAllowCredentials))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyAllowMethods))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyAllowHeaders))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyMaxAge))
	  ed.addListener(PreflightRequestEvent, nextForwarder(applyPreflightTermination))
	  ed.addListener(PreflightRequestEvent, applyNext)

	  ed.addListener(SimpleRequestEvent, nextForwarder(checkRequestIsCors))
	  ed.addListener(SimpleRequestEvent, nextForwarder(applyAllowOrigin))
	  ed.addListener(SimpleRequestEvent, nextForwarder(applyAllowCredentials))
	  ed.addListener(SimpleRequestEvent, nextForwarder(applyExposedHeaders))
	  ed.addListener(SimpleRequestEvent, applyNext)

	  e := new(event)
	  e.w = w
	  e.r = r
	  e.c = c
	  e.next = next

	  if isPreflightRequest(r) {
		ed.dispatch(e, PreflightRequestEvent)
	  } else {
		ed.dispatch(e, SimpleRequestEvent)
	  }
	}
  }
}

func isPreflightRequest(r *http.Request) bool {
  return r.Method == http.MethodOptions && r.Header.Get(RequestMethodHeader) != ""
}
