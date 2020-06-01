package cors

import (
  "fmt"
  "path/filepath"
  "reflect"
  "runtime"
  "testing"
)

// assertEquals fails the test if got result isn't equals expected result
func assertEquals(t *testing.T, want, got interface{}, msg string) {
  if reflect.DeepEqual(want, got) {
    return
  }

  fail(t, msg, want, got)
}

// assertTrue fails the test if the condition is false
func assertTrue(t *testing.T, condition bool, msg string) {
  if condition {
    return
  }

  fail(t, msg, true, condition)
}

// assertFalse fails the test if the condition is true
func assertFalse(t *testing.T, condition bool, msg string) {
  if !condition {
    return
  }

  fail(t, msg, false, condition)
}

// fail fails the test and prints place of error and error message
func fail(t *testing.T, msg string, want, got interface{}) {
  _, file, line, _ := runtime.Caller(2)
  fmt.Printf(
    "%s:%d %s\n want: %#v\n got: %#v\n",
    filepath.Base(file), line, msg, want, got,
  )
  t.FailNow()
}
