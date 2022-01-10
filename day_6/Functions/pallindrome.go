package functions

import (
	"strings"
)

func reverseString(s string) string {
	var out string
	for _, x := range s {
		out = string(x) + out
	}
	return out
}

func cleanInput(s string) string {
	var out string
	for _, x := range s {
		if (x > 47 && x < 58) || (x > 96 && x < 123) {
			out = out + string(x)
		}
	}
	return out
}

func CheckPallindrome(s string) bool {
	s = strings.ToLower(s)
	s = cleanInput(s)

	rev := reverseString(s)
	// fmt.Println(s, rev)
	if strings.Compare(s, rev) == 0 {
		return true
	}
	return false
}
