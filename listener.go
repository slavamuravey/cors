package cors

import (
  "net/http"
  "regexp"
  "strconv"
  "strings"
)

func handlePreflightTermination(e *event, ed *eventDispatcher) {
  if e.c.ContinuousPreflight {
    return
  }

  e.terminateRequest()
  e.stopPropagation()

  var status int
  if e.c.PreflightTerminationStatus == 0 {
    status = http.StatusOK
  } else {
    status = e.c.PreflightTerminationStatus
  }

  e.w.WriteHeader(status)
}

func handleExposedHeaders(e *event, ed *eventDispatcher) {
  if len(e.c.ExposedHeaders) > 0 {
    e.w.Header().Set(ExposeHeadersHeader, strings.Join(e.c.ExposedHeaders, ", "))
  }
}

func handleMaxAge(e *event, ed *eventDispatcher) {
  if e.c.MaxAge > 0 {
    e.w.Header().Set(MaxAgeHeader, strconv.Itoa(e.c.MaxAge))
  }
}

func handleAllowCredentials(e *event, ed *eventDispatcher) {
  if e.c.AllowCredentials {
    e.w.Header().Set(AllowCredentialsHeader, "true")
  }
}

func handleAllowOrigin(e *event, ed *eventDispatcher) {
  if e.c.AllowAllOrigin {
    e.w.Header().Set(AllowOriginHeader, "*")

    return
  }

  if e.c.AllowOriginPattern == "" {
    e.stopPropagation()

    return
  }

  origin := e.r.Header.Get(OriginHeader)
  match, err := regexp.MatchString(e.c.AllowOriginPattern, origin)

  if err != nil {
    e.w.WriteHeader(http.StatusInternalServerError)
    e.stopPropagation()
    e.terminateRequest()

    return
  }

  if match {
    e.w.Header().Set(AllowOriginHeader, origin)
  }
}

func handleAllowMethods(e *event, ed *eventDispatcher) {
  method := e.r.Header.Get(RequestMethodHeader)

  if !contains(strings.ToUpper(method), e.c.AllowMethods) {
    e.w.WriteHeader(http.StatusMethodNotAllowed)
    e.terminateRequest()
    e.stopPropagation()

    return
  }

  allowMethods := e.c.AllowMethods

  // If client sends method in not upper case we have to allow it.
  if !contains(method, e.c.AllowMethods) {
    allowMethods = append(allowMethods, method)
  }

  e.w.Header().Set(AllowMethodsHeader, strings.Join(allowMethods, ", "))
}

func handleAllowHeaders(e *event, ed *eventDispatcher) {
  requestHeaders := e.w.Header().Get(RequestHeadersHeader)

  if len(e.c.AllowHeaders) > 0 {
    var headers string
    if e.c.AllowAllHeaders {
      headers = requestHeaders
    } else {
      headers = strings.Join(e.c.AllowHeaders, ", ")
    }

    if headers != "" {
      e.w.Header().Set(AllowHeadersHeader, headers)
    }
  }

  if requestHeaders != "" && !e.c.AllowAllHeaders {
    r := regexp.MustCompile(` *, *`)
    headers := r.Split(strings.TrimSpace(requestHeaders), -1)

    for _, h := range headers {
      h = http.CanonicalHeaderKey(h)

      if !contains(h, e.c.AllowHeaders) {
        e.w.WriteHeader(http.StatusBadRequest)
        e.w.Write([]byte("Unauthorized header " + h))
        e.terminateRequest()
        e.stopPropagation()

        return
      }
    }
  }
}

func handleRequest(e *event, ed *eventDispatcher) {
  //If your server makes a decision about what to return based on a what’s in a HTTP header,
  //you need to include that header name in your Vary, even if the request didn’t include that header.
  //(https://textslashplain.com/2018/08/02/cors-and-vary/)
  e.w.Header().Add(VaryHeader, OriginHeader)
  e.w.Header().Add(VaryHeader, RequestMethodHeader)
  e.w.Header().Add(VaryHeader, RequestHeadersHeader)

  if !isCorsRequest(e.r) {
    return
  }

  corsEvent := newEvent(e.c, e.w, e.r)
  ed.dispatch(corsEvent, corsRequestEvent)

  if corsEvent.isRequestTerminated() {
    e.stopPropagation()
  }
}

func isCorsRequest(r *http.Request) bool {
  origin := r.Header.Get(OriginHeader)
  host := r.Host

  return !(origin == "" || origin == "http://"+host || origin == "https://"+host)
}

func handleCorsRequest(e *event, ed *eventDispatcher) {
  if isPreflightRequest(e.r) {
    ed.dispatch(e, preflightRequestEvent)
  } else {
    ed.dispatch(e, simpleRequestEvent)
  }
}

func isPreflightRequest(r *http.Request) bool {
  return r.Method == http.MethodOptions && r.Header.Get(RequestMethodHeader) != ""
}
