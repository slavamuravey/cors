package cors

import (
  "net/http"
)

func CreateHandlerFunc(c Config) func(http.Handler) http.HandlerFunc {
  return func(next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      ed := newEventDispatcher()

      ed.addListener(requestEvent, checkRequestIsCors)
      ed.addListener(requestEvent, func(e *event, ed *eventDispatcher) {
        next.ServeHTTP(w, r)
      })

      ed.addListener(corsRequestEvent, handleCorsRequest)

      ed.addListener(preflightRequestEvent, applyAllowOrigin)
      ed.addListener(preflightRequestEvent, applyAllowCredentials)
      ed.addListener(preflightRequestEvent, applyAllowMethods)
      ed.addListener(preflightRequestEvent, applyAllowHeaders)
      ed.addListener(preflightRequestEvent, applyMaxAge)
      ed.addListener(preflightRequestEvent, applyPreflightTermination)

      ed.addListener(simpleRequestEvent, applyAllowOrigin)
      ed.addListener(simpleRequestEvent, applyAllowCredentials)
      ed.addListener(simpleRequestEvent, applyExposedHeaders)

      ed.dispatch(newEvent(c, w, r), requestEvent)
    }
  }
}
