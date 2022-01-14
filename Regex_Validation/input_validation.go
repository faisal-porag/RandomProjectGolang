package Regex_Validation

import (
	env "RandomProjectGolang/environment_settings"
	"regexp"
	"strings"
)

var validInputPattern *regexp.Regexp

func init() {
	validInputPattern = regexp.MustCompile(env.GoDotEnvVariable("EXPECTED_INPUT_PATTERN"))
}

func IsValidInput(input string) bool {
	var trimmedInput = strings.TrimSpace(input)
	return validInputPattern.MatchString(trimmedInput)
}
