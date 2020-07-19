package str

import (
	"github.com/google/uuid"
	"regexp"
	"strings"
)

// Camel convert a given string to camelCase
func Camel(value string) string {

	link := regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	return link.ReplaceAllStringFunc(value, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// Snake convert a given string to snake_case
func Snake(value string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(value, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Trim remove all space at beginner and end of the given string
func Trim(value string) string {
	return strings.TrimSpace(value)
}

// Equals compare the two given string value if are equals
func Equals(first, second string) bool {
	return first == second
}

// Equals compare the two given string value if are equals
func EqualsIgnoreCase(first, second string) bool {
	return Trim(strings.ToLower(first)) == Trim(strings.ToLower(second))
}

// Addslashes add slash to a given string value if required.
//It returns the escaped string with backslashes in front of the pre-defined characters which is passed in the parameter.
func Addslashes(str string) string {
	var tmpRune []rune
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// UUID returns a string random generated UUID.
func UUID() string {

	return uuid.New().String()
}

//StartsWith method determines if the given string begins with the given value:
func StartsWith(haystack, needles string) bool {
	return strings.HasPrefix(haystack, needles)
}