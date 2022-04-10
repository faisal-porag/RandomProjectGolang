package utils

import (
	"fmt"
	"github.com/Faisal-CSE/RandomProjectGolang/config"
	"regexp"
	"time"
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

func IsTimeBetween2Times() bool {
  currentTime := time.Now() 
  // Time after 18 hours of currentTime
  futureTime := time.Now().Add(time.Hour * 18) 
  // Time after 10 hours of currentTime
  intermediateTime := time.Now().Add(time.Hour * 10) 
  if intermediateTime.After(currentTime) &&    intermediateTime.Before(futureTime) {
    return true
  } else {
    return false
  }
}
