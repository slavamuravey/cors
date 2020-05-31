package cors

import (
  "net/http"
)

func CreateHandlerFunc(c Config) func(http.Handler) http.HandlerFunc {
  return func(next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      ed := newEventDispatcher()

      ed.addListener(requestEvent, handleRequest)
      ed.addListener(requestEvent, func(e *event, ed *eventDispatcher) {
        next.ServeHTTP(w, r)
      })

      ed.addListener(corsRequestEvent, handleCorsRequest)

      ed.addListener(preflightRequestEvent, handleAllowOrigin)
      ed.addListener(preflightRequestEvent, handleAllowCredentials)
      ed.addListener(preflightRequestEvent, handleAllowMethods)
      ed.addListener(preflightRequestEvent, handleAllowHeaders)
      ed.addListener(preflightRequestEvent, handleMaxAge)
      ed.addListener(preflightRequestEvent, handlePreflightTermination)

      ed.addListener(simpleRequestEvent, handleAllowOrigin)
      ed.addListener(simpleRequestEvent, handleAllowCredentials)
      ed.addListener(simpleRequestEvent, handleExposedHeaders)

      ed.dispatch(newEvent(c, w, r), requestEvent)
    }
  }
}
