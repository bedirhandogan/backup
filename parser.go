package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ValidateFormatMatch(time string) {
	pattern := `^d:\d+ h:\d+ m:\d+ s:\d+$`

	re := regexp.MustCompile(pattern)

	if !re.MatchString(time) {
		fmt.Println("Specify time in wrong format.\n" + "Correct format: d:0 h:0 m:0 s:0")
		os.Exit(2)
	}
}

func ParseTime(duration string) Duration {
	ValidateFormatMatch(duration)

	times := make(map[string]int)

	parts := strings.Split(duration, " ") // [d:0 h:0 m:0 s:0]
	for _, part := range parts {
		keyValue := strings.Split(part, ":") // [d: 0] [h: 0] [m: 0] [s: 0]

		if len(keyValue) == 2 {
			key := keyValue[0]
			value := keyValue[1]

			valInt, err := strconv.Atoi(value) // convert int
			if err == nil {
				times[key] = valInt
			}
		}
	}

	return Duration{
		Day:  times["d"],
		Hour: times["h"],
		Min:  times["m"],
		Sec:  times["s"],
	}
}
