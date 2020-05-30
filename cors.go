package cors

import (
  "net/http"
)

func CreateHandlerFunc(c Config) func(http.Handler) http.HandlerFunc {
  return func(next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      ed := NewEventDispatcher()

      ed.addListener(RequestEvent, CheckRequestIsCors)
      ed.addListener(RequestEvent, func(e *Event, ed *EventDispatcher) {
        next.ServeHTTP(w, r)
      })

      ed.addListener(CorsRequestEvent, HandleCorsRequest)

      ed.addListener(PreflightRequestEvent, ApplyAllowOrigin)
      ed.addListener(PreflightRequestEvent, ApplyAllowCredentials)
      ed.addListener(PreflightRequestEvent, ApplyAllowMethods)
      ed.addListener(PreflightRequestEvent, ApplyAllowHeaders)
      ed.addListener(PreflightRequestEvent, ApplyMaxAge)
      ed.addListener(PreflightRequestEvent, ApplyPreflightTermination)

      ed.addListener(SimpleRequestEvent, ApplyAllowOrigin)
      ed.addListener(SimpleRequestEvent, ApplyAllowCredentials)
      ed.addListener(SimpleRequestEvent, ApplyExposedHeaders)

      ed.dispatch(NewEvent(c, w, r), RequestEvent)
    }
  }
}
