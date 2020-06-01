package cors

import (
  "fmt"
  "path/filepath"
  "runtime"
  "testing"
)

// True fails the test if the condition is false
func assertTrue(t *testing.T, ok bool, msg string) {
  if ok {
    return
  }

  fail(t, msg)
}

// False fails the test if the condition is true
func assertFalse(t *testing.T, ok bool, msg string) {
  if !ok {
    return
  }

  fail(t, msg)
}

// fail fails the test and prints place of error and error message
func fail(t *testing.T, msg string) {
  _, file, line, _ := runtime.Caller(2)
  fmt.Printf(
    "%s:%d %s\n",
    filepath.Base(file), line, msg,
  )
  t.FailNow()
}

