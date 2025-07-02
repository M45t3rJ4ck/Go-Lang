package automatedTesting

import (
	"testing"
)

// Goes back and makes sure your programm still works, while changing code.

func TestIsEmail(t *testing.T) {
	// To run test -> in containing folder in terminal enter -> go test -v
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email.")
	}

	_, err = IsEmail("wavoges@gmail.com")
	if err != nil {
		t.Error("wavoges@gmail.com is an email.")
	}

	_, err = IsEmail("wavoges@gmai")
	if err != nil {
		t.Error("wavoges@gmai is not an email.")
	}
}
