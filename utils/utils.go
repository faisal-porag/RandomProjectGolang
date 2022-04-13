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



func DateTimeFormatExample() {
	now := time.Now()
	fmt.Println(now) // 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
	fmt.Println(now.Format("20060102150405"))
	fmt.Println(now.Format("2006/01/02/15/04/05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05 (-0700)"))
	fmt.Println(now.Format("2006年01月02日 15:04"))
	fmt.Println(now.Format("2006-02-01"))
	fmt.Println(now.Format(time.Layout))   // 01/02 03:04:05PM '06 -0700
	fmt.Println(now.Format(time.ANSIC))    // Mon Jan _2 15:04:05 2006
	fmt.Println(now.Format(time.UnixDate)) // Mon Jan _2 15:04:05 MST 2006
	fmt.Println(now.Format(time.RubyDate)) // Mon Jan 02 15:04:05 -0700 2006
	fmt.Println(now.Format(time.RFC822))   // 02 Jan 06 15:04 MST
	fmt.Println(now.Format(time.RFC850))   // Monday, 02-Jan-06 15:04:05 MST
	fmt.Println(now.Format(time.Kitchen))  // 3:04PM
	fmt.Println(now.Format(time.Stamp))    // Jan _2 15:04:05
	
	currentTime := time.Now()
	fmt.Println("Current Time in String: ", currentTime.String())
    	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
    	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
    	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
    	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
    	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
    	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))
    	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))
    	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))
    	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))
    	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))
    	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))
    	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
    	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))
    	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))
    	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5")) 
    	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))  
    	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm")) 
}


//TODO EXAMPLE (binarySearch)
//items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
//fmt.Println(binarySearch(63, items))
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high{
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}


func EnglishTOBanglaNumber() {
	var finalEnlishToBanglaNumber = map[string]string{
		"0": "০",
		"1": "১",
		"2": "২",
		"3": "৩",
		"4": "৪",
		"5": "৫",
		"6": "৬",
		"7": "৭",
		"8": "৮",
		"9": "৯",
	}

	for _, v := range finalEnlishToBanglaNumber {
		fmt.Println(v)
	}
}


func ParseFlight(s string) (letters, numbers string) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	return string(l), string(n)
}

func GetParseFLightResult()  {
	flights := []string{"AB1234", "ABC1234", "ABC123", "AB12"}
	for _, flight := range flights {
		letters, numbers := ParseFlight(flight)
		fmt.Printf("%q: %q %q\n", flight, letters, numbers)
	}
}


func GetAllPossibleSubstringCombinationOfAString() {
	value := "abcdefghi"

	// Convert to rune slice for substrings.
	runes := []rune(value)

	// Loop over possible lengths, and possible start indexes.
	// ... Then take each possible substring from the source string.
	for length := 1; length < len(runes); length++ {
		for start := 0; start <= len(runes)-length; start++ {
			substring := string(runes[start : start+length])
			fmt.Println(substring)
		}
		//break
	}
}

func SplitingExample() {
    // Creating and initializing the strings
    str1 := "Welcome, to the, online portal, of GeeksforGeeks"
    str2 := "My dog name is Dollar"
    str3 := "I like to play Ludo"
 
    // Displaying strings
    fmt.Println("String 1: ", str1)
    fmt.Println("String 2: ", str2)
    fmt.Println("String 3: ", str3)
 
    // Splitting the given strings
    // Using SplitAfterN() function
    res1 := strings.SplitAfterN(str1, ",", 2)
    res2 := strings.SplitAfterN(str2, "", 4)
    res3 := strings.SplitAfterN(str3, "!", 1)
    res4 := strings.SplitAfterN("", "GeeksforGeeks, geeks", 3)
 
    // Displaying the result
    fmt.Println("\nResult 1: ", res1)
    fmt.Println("Result 2: ", res2)
    fmt.Println("Result 3: ", res3)
    fmt.Println("Result 4: ", res4)
 
}

func TrimingExample() {
    // Creating and initializing string
    // Using shorthand declaration
    str1 := "!!Welcome to GeeksforGeeks !!"
    str2 := "@@This is the tutorial of Golang$$"
  
    // Displaying strings
    fmt.Println("Strings before trimming:")
    fmt.Println("String 1: ", str1)
    fmt.Println("String 2:", str2)
  
    // Trimming the given strings
    // Using Trim() function
    res1 := strings.Trim(str1, "!")
    res2 := strings.Trim(str2, "@$")
  
    // Displaying the results
    fmt.Println("\nStrings after trimming:")
    fmt.Println("Result 1: ", res1)
    fmt.Println("Result 2:", res2)
}


func AddArrayDataWithCommaSeparatorString() {
    foodItems := []string{"pizza", "pasta", "sushi", "pho", "tikka masala"}
    // space after comma
    fmt.Println(strings.Join(foodItems, ", "))
}

func DateFormatV2() {
    current := time.Now().UTC()
    fmt.Println(current.Format("2006 Jan 02"))
}

func ValidDateChecker() {
	str1 := "31/07/2010"
	str2 := "1/13/2010"
	str3 := "29/2/2007"
	str4 := "31/08/2010"
	str5 := "29/02/200a"
	str6 := "29/02/200a"
	str7 := "55/02/200a"
	str8 := "2_/02/2009"

	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nDate: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Date: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Date: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Date: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Date: %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("Date: %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("Date: %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("Date: %v :%v\n", str8, re.MatchString(str8))
}

func LoopsInArray() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

func CheckStringHasSpecificSubString(source, data string) bool {
    //dataA := "Javascript"
    //ldataB := "Golang"

    //source := "this is a Golang"

    //if strings.Contains(source, dataA) {
        //fmt.Println("JavaScript")
   // }

    //if strings.Contains(source, dataB) {
        //fmt.Println("Golang")
    //}
	
     return strings.Contains(source, data)
}


