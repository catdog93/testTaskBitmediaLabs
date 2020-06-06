package validator

import (
	"errors"
	"regexp"
	"strings"
)

const regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const (
	emptyStringError = "error: lastName can't be empty"
	lengthError      = "error: lastName must has length between 2 to 40"

	emailFormatError = "The email field should be a valid email address!"
)

func LastNameValidation(lastName string) error {
	err := stringValidation(lastName)
	if err != nil {
		return err
	}
	return nil
}

func stringValidation(inputString string) error {
	if inputString == "" {
		return errors.New(emptyStringError)
	}
	inputString = strings.TrimSpace(inputString)
	if len(inputString) > 1 && len(inputString) < 40 {
		return errors.New(lengthError)
	}
	if !regexpEmail.Match(a.Email) {
		return erros.New()
	}
	return nil
}
