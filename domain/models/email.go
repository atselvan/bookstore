package models

import (
	"errors"
	"fmt"
	"regexp"
)

type Email string

var (
	ERREmailRequired = errors.New("email is required")
	ERREmailInvalid  = func(email Email) error {
		return fmt.Errorf("email (%s) is not valid", email)
	}
)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", ERREmailRequired
	}

	if !Email(email).IsValid() {
		return "", ERREmailInvalid(Email(email))
	}

	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) IsValid() bool {
	emailRegex := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(e.String())
}
