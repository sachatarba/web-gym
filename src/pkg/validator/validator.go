package validator

import (
	"net/mail"
	"regexp"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	correctNumber, err := regexp.MatchString("[(\\+{0,1}7)8]-\\d{3}-\\d{3}-\\d{2}-\\d{2}", phoneNumber)
	if err != nil {
		return false
	}

	return correctNumber
}
