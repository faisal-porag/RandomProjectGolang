package utils

import (
	"fmt"
	"github.com/Faisal-CSE/RandomProjectGolang/config"
	"regexp"
	"time"
	"strings"
   	"strconv"
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


//Reverse an array
func ReverseSlice(a []int) []int {
  for i := len(a)/2 - 1; i >= 0; i-- {
   pos := len(a) - 1 - i
   a[i], a[pos] = a[pos], a[i]
}
 return a
}


//TODO EXAMPLE (Convert a slice to a comma-separated string)
//result := ConvertSliceToString([]int{10, 20, 30, 40})
//fmt.Println("Slice converted string : ", result)
func ConvertSliceToString(input []int) string {
   var output []string
   for _, i := range input {
      output = append(output, strconv.Itoa(i))
   }
   return strings.Join(output, ",")
}
//output:
//Slice converted string :  10,20,30,40


//Convert given string to snake_case
func ConvertToSnakeCase(input string) string {
  var matchChars = regexp.MustCompile("(.)([A-Z][a-z]+)")
  var matchAlpha = regexp.MustCompile("([a-z0-9])([A-Z])")
  
  snake := matchChars.ReplaceAllString(input, "${1}_${2}")
  snake = matchAlpha.ReplaceAllString(snake, "${1}_${2}")
  return strings.ToLower(snake)
}
//input: ILikeProgrammingINGo123
//output:
//String in snake case :  i_like_programming_in_go123


