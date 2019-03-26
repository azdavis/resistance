package main

import (
	"regexp"
)

var validNameRE = regexp.MustCompile(`[^\s]`)

func ValidName(s string) bool {
	const maxLen = 32
	return s != "" && len(s) <= maxLen && validNameRE.Match([]byte(s))
}
