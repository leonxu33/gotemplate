package utils

import "strings"

func SplitRemoveEmpty(str string, token rune) []string {
	f := func(c rune) bool {
		return c == token
	}
	return strings.FieldsFunc(str, f)
}
