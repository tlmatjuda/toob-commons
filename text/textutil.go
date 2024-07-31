package text

import (
	"bytes"
	"strings"
	"unicode"
)

const (
	CHAR_FORWARD_SLASH = "/"
	CHAR_BACK_SLASH    = "\\"
	CHAR_FULL_STOP     = "."
	CHAR_ASTERIX       = "*"
	EMPTY              = ""
	WHITE_SPACE        = " "
	COLON              = ":"
)

// StringBlank
// Strempty checks whether string contains only whitespace or not
func StringBlank(s string) bool {
	if len(s) == 0 {
		return true
	}

	r := []rune(s)
	l := len(r)

	for l > 0 {
		l--
		if !unicode.IsSpace(r[l]) {
			return false
		}
	}

	return true
}

// StringNotBlank
// Used to negate the IsBankLogic
func StringNotBlank(s string) bool {
	return !StringBlank(s)
}

// ListContains
// Used to check if the String Slice contains a specific set of characters
func ListContains(arguments []string, arg string) bool {
	var contains bool
	for _, argItem := range arguments {
		if strings.Contains(argItem, arg) {
			contains = true
		}
	}

	return contains
}

// GetArg
// Used to get a CLI flag arg by supplying the collection of args and stating which flag value you want.
func GetArg(arguments []string, arg string) string {
	var responseArg string
	for _, argItem := range arguments {
		if strings.Contains(argItem, arg) {
			responseArg = strings.Split(argItem, "=")[1]
		}
	}

	return responseArg
}

// EqualsIgnoreCase
// Used to compare two strings while ignoring the case of the text.
func EqualsIgnoreCase(textArg string, anotherTextArg string) bool {
	return bytes.EqualFold([]byte(textArg), []byte(anotherTextArg))
}

// Equals
// Checks string Equality including case sensitivity
func Equals(textArg string, anotherTextArg string) bool {
	return textArg == anotherTextArg
}

// NotEquals
// Negates the NotEquals
func NotEquals(textArg string, anotherTextArg string) bool {
	return !Equals(textArg, anotherTextArg)
}

// Equals
// Negates the EqualsIgnoreCase
func NotEqualsIgnoreCase(textArg string, anotherTextArg string) bool {
	return !EqualsIgnoreCase(textArg, anotherTextArg)
}

func Trim(textValue string) string {
	return strings.TrimSpace(textValue)
}
