package cors

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

type testCase struct {
  config     *Config
  method     string
  reqHeaders http.Header
  resHeaders http.Header
  statusCode int
}

func TestNotCorsRequest(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
    },
    statusCode: 200,
  })
}

func TestSimpleEmptyAllowOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: "",
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
    },
    statusCode: 200,
  })
}

func TestSimpleAllowAllOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin": []string{"*"},
    },
    statusCode: 200,
  })
}

func TestSimpleAllowConcreteOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: "^http://foo.example.com$",
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin": []string{"http://foo.example.com"},
    },
    statusCode: 200,
  })
}

func TestSimpleCantParseOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: string([]byte{227}),
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
    },
    statusCode: 500,
  })
}

func TestSimpleExposeHeadersAllowCredentials(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin:   true,
      AllowCredentials: true,
      ExposedHeaders:   []string{"Header1", "Header2"},
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":      []string{"*"},
      "Access-Control-Allow-Credentials": []string{"true"},
      "Access-Control-Expose-Headers":    []string{"Header1,Header2"},
    },
    statusCode: 200,
  })
}

func TestPreflightAllowConcreteOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: "^http://foo.example.com$",
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "DELETE",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin": []string{"http://foo.example.com"},
    },
    statusCode: 405,
  })
}

func TestPreflightAllowCredentials(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin:   true,
      AllowCredentials: true,
      AllowMethods:     []string{"DELETE"},
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "DELETE",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":      []string{"*"},
      "Access-Control-Allow-Credentials": []string{"true"},
      "Access-Control-Allow-Methods":     []string{"DELETE"},
    },
    statusCode: 200,
  })
}

func TestPreflightEmptyAllowOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: "",
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "DELETE",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
    },
    statusCode: 200,
  })
}

func TestPreflightCantParseOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: string([]byte{227}),
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "DELETE",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
    },
    statusCode: 500,
  })
}

func TestPreflightFail(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
      MaxAge:         1000,
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "DELETE",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin": []string{"*"},
    },
    statusCode: 405,
  })
}

func TestPreflightMethodNotInUpperCase(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
      MaxAge:         1000,
      AllowMethods:   []string{"LINK"},
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "Link",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":  []string{"*"},
      "Access-Control-Max-Age":       []string{"1000"},
      "Access-Control-Allow-Methods": []string{"LINK,Link"},
    },
    statusCode: 200,
  })
}

func TestPreflightTerminationStatus(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin:             true,
      AllowMethods:               []string{"PUT"},
      PreflightTerminationStatus: 204,
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "PUT",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":  []string{"*"},
      "Access-Control-Allow-Methods": []string{"PUT"},
    },
    statusCode: 204,
  })
}

func TestPreflightContinuous(t *testing.T) {
  var nextFuncInvoked bool
  nextFunc := func(w http.ResponseWriter, r *http.Request) {
    nextFuncInvoked = true
  }
  hf := CreateHandlerFunc(&Config{
    AllowAllOrigin:      true,
    AllowMethods:        []string{"PUT"},
    ContinuousPreflight: true,
  })(http.HandlerFunc(nextFunc))
  res := httptest.NewRecorder()
  req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
  req.Header = http.Header{
    "Origin": []string{
      "http://foo.example.com",
    },
    "Access-Control-Request-Method": []string{
      "PUT",
    },
  }
  hf(res, req)

  assertTrue(t, nextFuncInvoked, "Next function must be invoked")
}

func TestPreflightNotContinuous(t *testing.T) {
  var nextFuncInvoked bool
  nextFunc := func(w http.ResponseWriter, r *http.Request) {
    nextFuncInvoked = true
  }
  hf := CreateHandlerFunc(&Config{
    AllowAllOrigin:      true,
    AllowMethods:        []string{"PUT"},
    ContinuousPreflight: false,
  })(http.HandlerFunc(nextFunc))
  res := httptest.NewRecorder()
  req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
  req.Header = http.Header{
    "Origin": []string{
      "http://foo.example.com",
    },
    "Access-Control-Request-Method": []string{
      "PUT",
    },
  }
  hf(res, req)

  assertFalse(t, nextFuncInvoked, "Next function mustn't be invoked")
}

func TestPreflightAllowHeadersFailure(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
      AllowMethods:   []string{"PUT"},
      AllowHeaders:   []string{"Header1", "Header2"},
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "PUT",
      },
      "Access-Control-Request-Headers": []string{
        "Header3,Header4",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":  []string{"*"},
      "Access-Control-Allow-Methods": []string{"PUT"},
      "Access-Control-Allow-Headers": []string{"Header1,Header2"},
    },
    statusCode: 400,
  })
}

func TestPreflightAllowHeadersSuccess(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
      AllowMethods:   []string{"PUT"},
      AllowHeaders:   []string{"Header1", "Header2"},
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "PUT",
      },
      "Access-Control-Request-Headers": []string{
        "Header1,Header2",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":  []string{"*"},
      "Access-Control-Allow-Methods": []string{"PUT"},
      "Access-Control-Allow-Headers": []string{"Header1,Header2"},
    },
    statusCode: 200,
  })
}

func TestPreflightAllowAllHeaders(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin:  true,
      AllowMethods:    []string{"PUT"},
      AllowHeaders:    []string{"Header1", "Header2"},
      AllowAllHeaders: true,
    },
    method: "OPTIONS",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://foo.example.com",
      },
      "Access-Control-Request-Method": []string{
        "PUT",
      },
      "Access-Control-Request-Headers": []string{
        "Header3,Header4",
      },
    },
    resHeaders: http.Header{
      "Vary": []string{
        "Origin",
        "Access-Control-Request-Method",
        "Access-Control-Request-Headers",
      },
      "Access-Control-Allow-Origin":  []string{"*"},
      "Access-Control-Allow-Methods": []string{"PUT"},
      "Access-Control-Allow-Headers": []string{"Header3,Header4"},
    },
    statusCode: 200,
  })
}

func testTestCase(t *testing.T, tc *testCase) {
  hf := CreateHandlerFunc(tc.config)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
  res := httptest.NewRecorder()
  req, _ := http.NewRequest(tc.method, "http://example.com/foo", nil)
  req.Header = tc.reqHeaders
  hf(res, req)

  assertEquals(t, tc.resHeaders, res.Header(), "Headers must be equals")
  assertEquals(t, tc.statusCode, res.Code, "Status code must be equals")
}
