package utils

import (
	"fmt"
	"github.com/Faisal-CSE/RandomProjectGolang/config"
	"regexp"
)

//PinMinLength = 4
//PinMinLength = 6

var PhoneNumberRe = regexp.MustCompile(`^(?:\+88)?01[1-9]\d{8}$`)
var PinNumberRe = regexp.MustCompile(fmt.Sprintf(`^\d{%d,%d}$`, config.Config.PinMinLength, config.Config.PinMaxLength))

func ValidatePhoneNumber(number string) bool {
	return PhoneNumberRe.MatchString(number)
}

func ValidatePinNumber(number string) bool {
	return PinNumberRe.MatchString(number)
}
