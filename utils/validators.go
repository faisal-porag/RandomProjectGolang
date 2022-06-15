package utils

import (
	"fmt"
	"regexp"
	"mail"
)

var minPINLength = 4
var maxPINLength = 6

var PhoneNumberRe = regexp.MustCompile(`^(?:\+88)?01[3-9]\d{8}$`)
var PinNumberRe = regexp.MustCompile(fmt.Sprintf(`^\d{%d,%d}$`, minPINLength, maxPINLength))

func ValidatePhoneNumber(number string) bool {
	return PhoneNumberRe.MatchString(number)
}

func ValidatePinNumber(number string) bool {
	return PinNumberRe.MatchString(number)
}

func IsValidEmailAddress(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}
