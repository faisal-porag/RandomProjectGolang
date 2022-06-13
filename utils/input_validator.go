package utils

import (
	"regexp"
	"strings"
)

var validInputPattern *regexp.Regexp

func init(){
	validInputPattern = regexp.MustCompile("^\\s|^[\/a-zA-Z0-9*#_@().:\\s-]+$")
}

func IsValidInput(input string) bool {
	var trimmedInput string = strings.TrimSpace(input)
	return validInputPattern.MatchString(trimmedInput)
}
