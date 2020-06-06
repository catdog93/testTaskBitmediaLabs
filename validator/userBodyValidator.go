package validator

import (
	"errors"
	"regexp"
	"strings"
	"testTaskBitmediaLabs/entity"
)

const (
	lastNameMinLength  = 2
	lastNameMaxLength  = 40
	genderMinLength    = 4
	genderMaxLength    = 10
	stringFieldPattern = "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$"
	emailPattern       = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	emailMinLength     = 10
	emailMaxLength     = 40
)

const (
	emptyStringError       = "error: string field can't be empty"
	lengthError            = "error: string field has incorrect length"
	stringFieldFormatError = "error: string field hasn't a valid format"
	emailFormatError       = "error: email field should be a valid email address"
	genderError            = "error: gender field requires such values \"Male\" or \"Female\""
)

func UserValidation(user entity.UserBody) error {
	err := stringFieldValidation(user.LastName)
	if err != nil {
		return err
	}
	err = genderValidation(user.Gender)
	if err != nil {
		return err
	}
	err = stringFieldValidation(user.Country)
	if err != nil {
		return err
	}
	err = stringFieldValidation(user.City)
	if err != nil {
		return err
	}
	err = emailValidation(user.Email)
	if err != nil {
		return err
	}
	return err
}

func stringFieldValidation(lastName string) error {
	if lastName == "" {
		return errors.New(emptyStringError)
	}
	lastName = strings.TrimSpace(lastName)
	if len(lastName) < lastNameMinLength || len(lastName) > lastNameMaxLength {
		return errors.New(lengthError)
	}
	regex := regexp.MustCompile(stringFieldPattern)
	if !regex.MatchString(lastName) {
		return errors.New(stringFieldFormatError)
	}
	return nil
}

func stringLengthValidation(inputString string, minLength, maxLength int) error {
	if inputString == "" {
		return errors.New(emptyStringError)
	}
	if len(inputString) < minLength || len(inputString) > maxLength {
		return errors.New(lengthError)
	}
	return nil
}

func emailValidation(email string) error {
	email = strings.TrimSpace(email)
	err := stringLengthValidation(email, emailMinLength, emailMaxLength)
	if err != nil {
		return err
	}
	regex := regexp.MustCompile(emailPattern)
	if !regex.MatchString(email) {
		return errors.New(emailFormatError)
	}
	return nil
}

func genderValidation(gender entity.Gender) error {
	genderString := string(gender)
	genderString = strings.TrimSpace(genderString)
	err := stringLengthValidation(genderString, genderMinLength, genderMaxLength)
	if err != nil {
		return err
	}
	genderLower := entity.Gender(strings.ToLower(genderString))
	if genderLower != entity.MaleLower && genderLower != entity.FemaleLower {
		return errors.New(genderError)
	}
	return nil
}
