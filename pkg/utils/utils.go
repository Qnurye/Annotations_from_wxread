package utils

import (
	"unicode/utf8"
)

func LastCharsOf(s string, n int) string {
	strLen := utf8.RuneCountInString(s)
	if strLen < n {
		return ""
	}
	runes := []rune(s)
	r := runes[strLen-n:]
	c := string(r)
	return c
}

func FirstCharsOf(s string, n int) string {
	runes := []rune(s)
	if len(runes) < n {
		return ""
	}
	r := runes[:n]
	c := string(r)
	return c
}
