package greetings

import (
	"regexp"
	"testing"
)

func Test_Can_Display_Hello(t *testing.T) {
	value := "eliseekn"
	expected := regexp.MustCompile(`\b` + value + `\b`)
	message, err := Hello(value)
	if !expected.MatchString(message) || err != nil {
		t.Fatal(err)
	}
}
