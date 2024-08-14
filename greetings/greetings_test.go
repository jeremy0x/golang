package greetings

import (
	"regexp"
	"testing"
)

// call greetings.Hello with a name and check for a valid return value
func TestHelloName(t *testing.T)  {
	name := "Jeremy"
	want := regexp.MustCompile(`\b`+name+`\b`)
	message, error := Hello("Jeremy")
	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Jeremy") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

// calls greetings.Hello with an empty string and checks for an error
func TestHelloEmpty(t *testing.T)  {
	message, error := Hello("")
	if message != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}
}