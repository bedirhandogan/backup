package main

import (
	"regexp"
)

var colors = map[string]string{
	// regular
	"black":  "\033[0;30m",
	"red":    "\033[0;31m",
	"green":  "\033[0;32m",
	"yellow": "\033[0;33m",
	"blue":   "\033[0;34m",
	"purple": "\033[0;35m",
	"cyan":   "\033[0;36m",
	"white":  "\033[0;37m",
	"gray":   "\033[0;90m",

	// background
	"bgblack":  "\033[0;40m",
	"bgred":    "\033[0;41m",
	"bggreen":  "\033[0;42m",
	"bgyellow": "\033[0;43m",
	"bgblue":   "\033[0;44m",
	"bgpurple": "\033[0;45m",
	"bgcyan":   "\033[0;46m",
	"bgwhite":  "\033[0;47m",
	"bggray":   "\033[0;100m",

	// bold
	"bblack":  "\033[0;1;30m",
	"bred":    "\033[0;1;31m",
	"bgreen":  "\033[0;1;32m",
	"byellow": "\033[0;1;33m",
	"bblue":   "\033[0;1;34m",
	"bpurple": "\033[0;1;35m",
	"bcyan":   "\033[0;1;36m",
	"bwhite":  "\033[0;1;37m",
	"bgray":   "\033[0;1;90m",

	// underline
	"ublack":  "\033[4;30m",
	"ured":    "\033[4;31m",
	"ugreen":  "\033[4;32m",
	"uyellow": "\033[4;33m",
	"ublue":   "\033[4;34m",
	"upurple": "\033[4;35m",
	"ucyan":   "\033[4;36m",
	"uwhite":  "\033[4;37m",
	"ugray":   "\033[4;90m",

	// reset
	"reset": "\033[0m",
}

func Color(text string) string {
	for key, value := range colors {
		pattern := `(?i)(%` + key + `)\s*` // yellow
		text = regexp.MustCompile(pattern).ReplaceAllString(text, value)
	}

	return text + colors["reset"]
}
