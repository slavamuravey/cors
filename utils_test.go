package cors

import (
  "testing"
)

func TestContainsTrue(t *testing.T) {
  needle := "needle"
  haystack := []string{
    "needle",
    "something else",
  }
  assertTrue(t, contains(needle, haystack), "Needle must be presented in haystack")
}

func TestContainsFalse(t *testing.T) {
  needle := "needle"
  haystack := []string{
    "something else",
  }
  assertFalse(t, contains(needle, haystack), "Needle mustn't be presented in haystack")
}

func TestContainsEmptyHaystack(t *testing.T) {
  needle := "needle"
  var haystack []string
  assertFalse(t, contains(needle, haystack), "Needle mustn't be presented in haystack")
}
