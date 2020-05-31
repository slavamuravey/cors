package cors

import (
  "net/http"
  "regexp"
  "strconv"
  "strings"
)

func ApplyPreflightTermination(e *Event, ed *EventDispatcher) {
  if e.C.ContinuousPreflight {
    return
  }

  e.TerminateRequest()
  e.StopPropagation()

  e.W.WriteHeader(http.StatusOK)
}

func ApplyExposedHeaders(e *Event, ed *EventDispatcher) {
  if len(e.C.ExposedHeaders) > 0 {
    e.W.Header().Set(ExposeHeadersHeader, strings.Join(e.C.ExposedHeaders, ", "))
  }
}

func ApplyMaxAge(e *Event, ed *EventDispatcher) {
  if e.C.MaxAge > 0 {
    e.W.Header().Set(MaxAgeHeader, strconv.Itoa(e.C.MaxAge))
  }
}

func ApplyAllowCredentials(e *Event, ed *EventDispatcher) {
  if e.C.AllowCredentials {
    e.W.Header().Set(AllowCredentialsHeader, "true")
  }
}

func ApplyAllowOrigin(e *Event, ed *EventDispatcher) {
  if e.C.AllowAllOrigin {
    e.W.Header().Add(VaryHeader, OriginHeader)
    e.W.Header().Set(AllowOriginHeader, "*")

    return
  }

  if e.C.AllowOriginPattern == "" {
    e.StopPropagation()

    return
  }

  origin := e.R.Header.Get(OriginHeader)
  match, err := regexp.MatchString(e.C.AllowOriginPattern, origin)

  if err != nil {
    e.W.WriteHeader(http.StatusInternalServerError)
    e.StopPropagation()
    e.TerminateRequest()

    return
  }

  if match {
    e.W.Header().Add(VaryHeader, OriginHeader)
    e.W.Header().Set(AllowOriginHeader, origin)
  }
}

func ApplyAllowMethods(e *Event, ed *EventDispatcher) {
  method := e.R.Header.Get(RequestMethodHeader)

  if !contains(strings.ToUpper(method), e.C.AllowMethods) {
    e.W.WriteHeader(http.StatusMethodNotAllowed)
    e.TerminateRequest()
    e.StopPropagation()

    return
  }

  allowMethods := e.C.AllowMethods

  // If client sends method in not upper case we have to allow it.
  if !contains(method, e.C.AllowMethods) {
    allowMethods = append(allowMethods, method)
  }

  e.W.Header().Add(VaryHeader, RequestMethodHeader)
  e.W.Header().Set(AllowMethodsHeader, strings.Join(allowMethods, ", "))
}

func ApplyAllowHeaders(e *Event, ed *EventDispatcher) {
  e.W.Header().Add(VaryHeader, RequestHeadersHeader)
  requestHeaders := e.W.Header().Get(RequestHeadersHeader)

  if len(e.C.AllowHeaders) > 0 {
    var headers string
    if e.C.AllowAllHeaders {
      headers = requestHeaders
    } else {
      headers = strings.Join(e.C.AllowHeaders, ", ")
    }

    if headers != "" {
      e.W.Header().Set(AllowHeadersHeader, headers)
    }
  }

  if requestHeaders != "" && !e.C.AllowAllHeaders {
    r := regexp.MustCompile(` *, *`)
    headers := r.Split(strings.TrimSpace(requestHeaders), -1)

    for _, h := range headers {
      h = http.CanonicalHeaderKey(h)

      if !contains(h, e.C.AllowHeaders) {
        e.W.WriteHeader(http.StatusBadRequest)
        e.W.Write([]byte("Unauthorized header " + h))
        e.TerminateRequest()
        e.StopPropagation()

        return
      }
    }
  }
}

func CheckRequestIsCors(e *Event, ed *EventDispatcher) {
  origin := e.R.Header.Get(OriginHeader)
  host := e.R.Host

  if origin == "" || origin == "http://"+host || origin == "https://"+host {
    return
  }

  corsRequestEvent := NewEvent(e.C, e.W, e.R)
  ed.dispatch(corsRequestEvent, CorsRequestEvent)

  if corsRequestEvent.IsRequestTerminated() {
    e.StopPropagation()
  }
}

func HandleCorsRequest(e *Event, ed *EventDispatcher) {
  if isPreflightRequest(e.R) {
    ed.dispatch(e, PreflightRequestEvent)
  } else {
    ed.dispatch(e, SimpleRequestEvent)
  }
}

func isPreflightRequest(r *http.Request) bool {
  return r.Method == http.MethodOptions && r.Header.Get(RequestMethodHeader) != ""
}
