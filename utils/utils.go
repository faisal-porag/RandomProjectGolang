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

func FindCurrentTimeWithTimeZone() {
  timeZone := "Asia/Kolkata" // timezone value
  loc, _ := time.LoadLocation(timeZone)
  currentTime = time.Now().In(loc)
  fmt.Println("currentTime : ", currentTime)
}

//TODO EXAMPLE
  // define array of strings
  //fruits := []string{"Mango", "Grapes", "Kiwi", "Apple", "Grapes"}
  //fmt.Println("Array before removing duplicates : ", fruits)
  //// Array after duplicates removal
  //dulicatesRemovedArray := RemoveDuplicatesFromSlice(fruits)
  //fmt.Println("Array after removing duplicates : ",  dulicatesRemovedArray)

func RemoveDuplicatesFromSlice(intSlice []string) []string {
  keys := make(map[string]bool)
  list := []string{}
  for _, entry := range intSlice {
    if _, value := keys[entry]; !value {
      keys[entry] = true
      list = append(list, entry)
    }
  }
 return list
}





