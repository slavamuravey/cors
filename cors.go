package cors

import (
  "net/http"
  "regexp"
  "strconv"
  "strings"
)

// CreateMiddleware returns middleware for using in client's code based on configuration. Middleware is function that
// receives http.Handler interface and returns http.HandlerFunc function, that implements http.Handler interface.
// You can pass this function as http.HandlerFunc into http.HandleFunc. And you can pass it into http.ListenAndServe.
func CreateMiddleware(c *Config) func(http.Handler) http.HandlerFunc {
  return func(next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      // If your server makes a decision about what to return based on a what’s in a HTTP header,
      // you need to include that header name in your Vary, even if the request didn’t include that header.
      // (https://textslashplain.com/2018/08/02/cors-and-vary/)
      w.Header().Add(VaryHeader, OriginHeader)
      w.Header().Add(VaryHeader, RequestMethodHeader)
      w.Header().Add(VaryHeader, RequestHeadersHeader)

      if !isCorsRequest(r) {
        next.ServeHTTP(w, r)
      } else if isPreflightRequest(r) {
        handlePreflightRequest(c, w, r, next)
      } else {
        handleSimpleRequest(c, w, r, next)
      }
    }
  }
}

// handlePreflightRequest handles simple cross-origin request
func handleSimpleRequest(c *Config, w http.ResponseWriter, r *http.Request, next http.Handler) {
  if c.AllowAllOrigin {
    w.Header().Set(AllowOriginHeader, "*")
  } else if c.AllowOriginPattern != "" {
    origin := r.Header.Get(OriginHeader)
    match, err := regexp.MatchString(c.AllowOriginPattern, origin)

    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("Origin header validation error: " + err.Error()))

      return
    }

    if match {
      w.Header().Set(AllowOriginHeader, origin)
    }
  } else {
    next.ServeHTTP(w, r)

    return
  }

  if c.AllowCredentials {
    w.Header().Set(AllowCredentialsHeader, "true")
  }

  if len(c.ExposedHeaders) > 0 {
    w.Header().Set(ExposeHeadersHeader, strings.Join(c.ExposedHeaders, ","))
  }

  next.ServeHTTP(w, r)
}

// handlePreflightRequest handles preflight cross-origin request
func handlePreflightRequest(c *Config, w http.ResponseWriter, r *http.Request, next http.Handler) {
  if c.AllowAllOrigin {
    w.Header().Set(AllowOriginHeader, "*")
  } else if c.AllowOriginPattern != "" {
    origin := r.Header.Get(OriginHeader)
    match, err := regexp.MatchString(c.AllowOriginPattern, origin)

    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("Origin header validation error: " + err.Error()))

      return
    }

    if match {
      w.Header().Set(AllowOriginHeader, origin)
    }
  } else {
    next.ServeHTTP(w, r)

    return
  }

  if c.AllowCredentials {
    w.Header().Set(AllowCredentialsHeader, "true")
  }

  method := r.Header.Get(RequestMethodHeader)

  if !contains(strings.ToUpper(method), c.AllowMethods) {
    w.WriteHeader(http.StatusMethodNotAllowed)

    return
  }

  allowMethods := c.AllowMethods

  // If client sends method in not upper case we have to allow it.
  if !contains(method, c.AllowMethods) {
    allowMethods = append(allowMethods, method)
  }

  w.Header().Set(AllowMethodsHeader, strings.Join(allowMethods, ","))

  requestHeaders := r.Header.Get(RequestHeadersHeader)

  var headers string
  if c.AllowAllHeaders {
    headers = requestHeaders
  } else {
    headers = strings.Join(c.AllowHeaders, ",")
  }

  if headers != "" {
    w.Header().Set(AllowHeadersHeader, headers)
  }

  if requestHeaders != "" && !c.AllowAllHeaders {
    r := regexp.MustCompile(` *, *`)
    headers := r.Split(strings.TrimSpace(requestHeaders), -1)

    for _, h := range headers {
      h = http.CanonicalHeaderKey(h)

      if !contains(h, c.AllowHeaders) {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Unauthorized header " + h))

        return
      }
    }
  }

  if c.MaxAge > 0 {
    w.Header().Set(MaxAgeHeader, strconv.Itoa(c.MaxAge))
  }

  if !c.ContinuousPreflight {
    var status int
    if c.PreflightTerminationStatus == 0 {
      status = http.StatusOK
    } else {
      status = c.PreflightTerminationStatus
    }

    w.WriteHeader(status)

    return
  }

  next.ServeHTTP(w, r)
}

// isCorsRequest checks if request is CORS
func isCorsRequest(r *http.Request) bool {
  origin := r.Header.Get(OriginHeader)
  host := r.Host

  return !(origin == "" || origin == "http://"+host || origin == "https://"+host)
}

// isPreflightRequest checks if request is preflight
func isPreflightRequest(r *http.Request) bool {
  return r.Method == http.MethodOptions && r.Header.Get(RequestMethodHeader) != ""
}
