package main

import (
	"strings"
)

var colors = map[string]string{
	// regular
	"%black":  "\033[30m",
	"%red":    "\033[31m",
	"%green":  "\033[32m",
	"%yellow": "\033[33m",
	"%blue":   "\033[34m",
	"%purple": "\033[35m",
	"%cyan":   "\033[36m",
	"%white":  "\033[37m",
	"%gray":   "\033[90m",

	// background
	"%bgblack":  "\033[40m",
	"%bgred":    "\033[41m",
	"%bggreen":  "\033[42m",
	"%bgyellow": "\033[43m",
	"%bgblue":   "\033[44m",
	"%bgpurple": "\033[45m",
	"%bgcyan":   "\033[46m",
	"%bgwhite":  "\033[47m",
	"%bggray":   "\033[100m",

	// bold
	"%bblack":  "\033[1;30m",
	"%bred":    "\033[1;31m",
	"%bgreen":  "\033[1;32m",
	"%byellow": "\033[1;33m",
	"%bblue":   "\033[1;34m",
	"%bpurple": "\033[1;35m",
	"%bcyan":   "\033[1;36m",
	"%bwhite":  "\033[1;37m",
	"%bgray":   "\033[1;90m",

	// underline
	"%ublack":  "\033[4;30m",
	"%ured":    "\033[4;31m",
	"%ugreen":  "\033[4;32m",
	"%uyellow": "\033[4;33m",
	"%ublue":   "\033[4;34m",
	"%upurple": "\033[4;35m",
	"%ucyan":   "\033[4;36m",
	"%uwhite":  "\033[4;37m",
	"%ugray":   "\033[4;90m",

	// reset
	"%reset": "\033[0m",
}

func Color(text string) string {
	for key, value := range colors {
		tk := strings.ToLower(key)
		text = strings.ReplaceAll(text, tk, value)
	}

	return text
}