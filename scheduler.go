package main

import (
	"time"
)

type Callback func() error

func TaskScheduler(callback Callback, day, hour, min, sec int) {
	interval := time.Duration(day*24*int(time.Hour) + hour*int(time.Hour) + min*int(time.Minute) + sec*int(time.Second))
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for {
		<-ticker.C

		err := callback()

		if err != nil {
			return
		}
	}
}
