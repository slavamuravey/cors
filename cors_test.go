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

func TestEmptyAllowOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowOriginPattern: "",
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://example.com/",
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

func TestAllowAllOrigin(t *testing.T) {
  testTestCase(t, &testCase{
    config: &Config{
      AllowAllOrigin: true,
    },
    method: "GET",
    reqHeaders: http.Header{
      "Origin": []string{
        "http://example.com/",
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

func testTestCase(t *testing.T, tc *testCase) {
  hf := CreateHandlerFunc(tc.config)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
  res := httptest.NewRecorder()
  req, _ := http.NewRequest(tc.method, "http://example.com/foo", nil)
  req.Header = tc.reqHeaders
  hf(res, req)

  assertEquals(t, tc.resHeaders, res.Header(), "Headers must be equals")
  assertEquals(t, tc.statusCode, 200, "Status code must be equals")
}
