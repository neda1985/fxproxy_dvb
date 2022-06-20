package validator

import (
	"regexp"
	"strings"
)

var GlobalMatcher Matcher
var AllowedList = []string{
	"/company/",
	"/company/{id}",
	"/company/account",
	"/account",
	"/account/{id}",
	"/{id}",
	"/account/{id}/user",
	"/tenant/account/blocked",
}

var idRegexp = regexp.MustCompile(`^[a-z\d]+\d+`)

type Matcher [][]string

func (m *Matcher) matchSingle(split []string, pattern []string) bool {
	if len(split) != len(pattern) {
		return false
	}
	for i := range pattern {
		if pattern[i] == "{id}" {
			if !idRegexp.MatchString(split[i]) {
				return false
			}
			continue
		}

		if pattern[i] != split[i] {
			return false
		}
	}

	return true
}

func (m Matcher) Validate(path string) bool {
	split := strings.Split(strings.Trim(path, "/"), "/")
	for i := range m {
		if m.matchSingle(split, m[i]) {
			return true
		}
	}
	return false
}

func NewMatcher(in []string) Matcher {
	m := Matcher{}
	for i := range in {
		m = append(m, strings.Split(strings.Trim(in[i], "/"), "/"))
	}
	return m
}

func ValidatePath(path string) bool {
	return GlobalMatcher.Validate(path)
}
