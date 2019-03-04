package main

import (
	"regexp"
)

var validNameRE = regexp.MustCompile(`^\w+$`)

// validName returns whether the name is valid.
func validName(s string) bool {
	const maxLen = 16
	return s != "" && len(s) < maxLen && validNameRE.Match([]byte(s))
}
