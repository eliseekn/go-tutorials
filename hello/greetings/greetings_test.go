package greetings_test

import (
	"hello/greetings"
	"regexp"
	"testing"
)

func Test_Can_Display_Hello(t *testing.T) {
	value := "eliseekn"
	expected := regexp.MustCompile(`\b` + value + `\b`)

	message, err := greetings.Hello(value)
	if err != nil {
		t.Fatal(err)
	}

	if !expected.MatchString(message) || err != nil {
		t.FailNow()
	}
}

func Test_Can_Display_Multiple_Hello(t *testing.T) {
	values := []string{"eliseekn", "wrh1d3"}

	messages, err := greetings.MultipleHello(values)
	if err != nil {
		t.Fatal(err)
	}

	for i := range values {
		expected := regexp.MustCompile(`\b` + values[i] + `\b`)

		if !expected.MatchString(messages[values[i]]) {
			t.FailNow()
		}
	}
}
