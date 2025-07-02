package automatedTesting

import (
	"fmt"
	"regexp"
)

// Goes back and makes sure your programm still works, while changing code.
func IsEmail(s string) (string, error) {
	validator, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if validator.MatchString(s) {
		return "Valid Email.", nil
	} else {
		// Errorf text must be in lower case and should not end with punctuation or new lines.
		return "", fmt.Errorf("not a valid email")
	}
}
