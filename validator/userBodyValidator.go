package validator

import (
	"errors"
	"regexp"
	"strings"
	"testTaskBitmediaLabs/entity"
	"time"
)

const (
	stringFieldPattern   = "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$"
	stringFieldMinLength = 2
	stringFieldMaxLength = 40
	emailPattern         = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	emailMinLength       = 10
	emailMaxLength       = 40
	genderMinLength      = 4
	genderMaxLength      = 10
)

const (
	emptyStringError       = "error: string field can't be empty"
	lengthError            = "error: string field has incorrect length"
	stringFieldFormatError = "error: string field hasn't a valid format"
	emailFormatError       = "error: email field should be a valid email address"
	genderError            = "error: gender field requires such values \"Male\" or \"Female\""
)

const birthDateFormat = "Monday, January 02, 2006 15:04 AM"

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
	err = birthDateValidation(birthDateFormat, user.BirthDate)
	return err
}

func stringFieldValidation(lastName string) error {
	if lastName == "" {
		return errors.New(emptyStringError)
	}
	lastName = strings.TrimSpace(lastName)
	if len(lastName) < stringFieldMinLength || len(lastName) > stringFieldMaxLength {
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

func birthDateValidation(layout, dateString string) error {
	_, err := time.Parse(layout, dateString)
	return err
}
