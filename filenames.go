package main

import "regexp"

var regexps = []*regexp.Regexp{
	regexp.MustCompile(`(?i)user`),
	regexp.MustCompile(`(?i)password`),
	regexp.MustCompile(`(?i)kbd`),
	regexp.MustCompile(`(?i)login`),
}
