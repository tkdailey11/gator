package server

import (
	"regexp"
	"testing"

	"github.com/tkdailey11/gator/pkg/server"
)

// TestHelloWorld calls server.Hello with a name, checking
// for a valid return value.
func TestHelloWorld(t *testing.T) {
  name := "World"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := server.Hello(name)
  if !want.MatchString(msg) || err != nil {
      t.Fatalf(`Hello("World") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
  msg, err := server.Hello("")
  if msg != "" || err == nil {
      t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
  }
}