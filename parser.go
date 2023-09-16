package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ValidateFormatMatch(duration string) {
	pattern := `^d:\d+ h:\d+ m:\d+ s:\d+$`

	re := regexp.MustCompile(pattern)

	if !re.MatchString(duration) {
		fmt.Println(Color("%red Error: Specify time in wrong format. " + "%cyan Correct format: %gray d:0 h:0 m:0 s:0"))
		os.Exit(2)
	}
}

func ParseTime(duration string) TimePeriod {
	ValidateFormatMatch(duration)

	tp := make(map[string]int)

	parts := strings.Split(duration, " ") // [d:0 h:0 m:0 s:0]
	for _, part := range parts {
		keyValue := strings.Split(part, ":") // [d: 0] [h: 0] [m: 0] [s: 0]

		if len(keyValue) == 2 {
			key := keyValue[0]
			value := keyValue[1]

			valInt, err := strconv.Atoi(value) // convert int
			if err == nil {
				tp[key] = valInt
			}
		}
	}

	if tt := tp["d"] + tp["h"] + tp["m"] + tp["s"]; tt == 0 {
		fmt.Println(Color("%red Error: You did not specify a duration, please specify a duration."))
		os.Exit(1)
	}

	return TimePeriod{
		Day:  tp["d"],
		Hour: tp["h"],
		Min:  tp["m"],
		Sec:  tp["s"],
	}
}
