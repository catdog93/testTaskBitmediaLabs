package validator

import (
	"errors"
	"regexp"
	"strings"
	"testTaskBitmediaLabs/entity"
	"time"
)

const (
	lastNameMinLength  = 2
	lastNameMaxLength  = 40
	genderMinLength    = 4
	genderMaxLength    = 10
	stringFieldPattern = "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$"
	emailPattern       = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	emailMinLength     = 16
	emailMaxLength     = 60
	birthDateFormat    = "Tuesday, February 9, 0388 2:03 PM"
)

const (
	emptyStringError     = "error: string field can't be empty"
	lengthError          = "error: string field must has length between 2 to 40"
	lastNameFormatError  = "error: lastName hasn't a valid format"
	emailFormatError     = "error: email field should be a valid email address"
	genderError          = "error: gender field requires such values \"Male\" or \"Female\""
	birthDateFormatError = "error: birthDate hasn't a valid date format"
	yearOfBirthError     = "error: year of birth isn't valid"
)

func UserValidation(user entity.UserBody) error {
	err := lastFieldValidation(user.LastName)
	if err != nil {
		return err
	}
	err = genderValidation(user.Gender)
	if err != nil {
		return err
	}
	err = lastFieldValidation(user.Country)
	if err != nil {
		return err
	}
	err = lastFieldValidation(user.City)
	if err != nil {
		return err
	}
	err = emailValidation(user.Email)
	if err != nil {
		return err
	}
	err = birthdayValidation(user.BirthDate)
	return err
}

func lastFieldValidation(lastName string) error {
	if lastName == "" {
		return errors.New(emptyStringError)
	}
	lastName = strings.TrimSpace(lastName)
	if len(lastName) < lastNameMinLength || len(lastName) > lastNameMaxLength {
		return errors.New(lengthError)
	}
	regex := regexp.MustCompile(stringFieldPattern)
	if !regex.MatchString(lastName) {
		return errors.New(lastNameFormatError)
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

func birthdayValidation(birthday string) error {
	birthday = strings.TrimSpace(birthday)
	d, err := time.Parse(birthDateFormat, birthday)
	if err != nil {
		return errors.New(birthDateFormatError)
	}
	if d.Year() < 1900 {
		return errors.New(yearOfBirthError)
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
