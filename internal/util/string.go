package util

import "strings"

func splitDelimsFunc(delims []rune) func(rune) bool {
	return func(r rune) bool {
		for _, delim := range delims {
			if delim == r {
				return true
			}
		}
		return false
	}
}

func Split(str string, delims ...rune) []string {
	return strings.FieldsFunc(str, splitDelimsFunc(delims))
}
