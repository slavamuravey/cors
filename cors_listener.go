package cors

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func applyNext(e *event, ed *eventDispatcher) {
	e.next.ServeHTTP(e.w, e.r)
}

func applyPreflightTermination(e *event, ed *eventDispatcher) {
	if e.c.ContinuousPreflight {
		return
	}

	e.stopPropagation()

	e.w.Header().Add("Content-Length", "0")
	e.w.WriteHeader(http.StatusOK)
}

func applyExposedHeaders(e *event, ed *eventDispatcher) {
	if len(e.c.ExposedHeaders) > 0 {
		e.w.Header().Set(ExposeHeadersHeader, strings.Join(e.c.ExposedHeaders, ", "))
	}
}

func applyMaxAge(e *event, ed *eventDispatcher) {
	if e.c.MaxAge > 0 {
		e.w.Header().Set(MaxAgeHeader, strconv.Itoa(e.c.MaxAge))
	}
}

func applyAllowCredentials(e *event, ed *eventDispatcher) {
	if e.c.AllowCredentials {
		e.w.Header().Set(AllowCredentialsHeader, "true")
	}
}

func applyAllowOrigin(e *event, ed *eventDispatcher) {
	if e.c.AllowOrigin == "" || e.c.AllowOrigin == "*" {
		e.w.Header().Add(VaryHeader, OriginHeader)
		e.w.Header().Set(AllowOriginHeader, "*")

		return
	}

	origin := e.r.Header.Get(OriginHeader)
	match, err := regexp.MatchString(e.c.AllowOrigin, origin)

	if err != nil {
		return
	}

	if match {
		e.w.Header().Add(VaryHeader, OriginHeader)
		e.w.Header().Set(AllowOriginHeader, origin)
	}
}

func applyAllowMethods(e *event, ed *eventDispatcher) {
	//Todo
	//method := e.r.Header.Get(RequestMethodHeader)
	//
	//e.w.Header().Add(VaryHeader, RequestMethodHeader)
	//e.w.Header().Set(AllowMethodsHeader, strings.ToUpper(method))
}

func applyAllowHeaders(e *event, ed *eventDispatcher) {
	//e.w.Header().Add(VaryHeader, RequestHeadersHeader)
	//Todo
}

func checkRequestIsCors(e *event, ed *eventDispatcher) {
	origin := e.r.Header.Get(OriginHeader)
	host := e.r.Host

	if origin == "" || origin == "http://"+host || origin == "https://"+host {
		e.forwardToNext()
	}
}

func nextForwarder(l listener) listener {
	return func(e *event, ed *eventDispatcher) {
		if e.isForwardedToNext() {
			return
		}

		l(e, ed)
	}
}
