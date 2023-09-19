package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	source := flag.String("source", "", "Path to the folder to copy.")
	destination := flag.String("destination", "", "The path to the target folder.")
	duration := flag.String("duration", "", "Backup duration of the program.")
	flag.Parse()

	if *source == "" || *destination == "" || *duration == "" {
		fmt.Println(Color("%yellow Use: %reset go run . %gray -source %cyan [e.g: C:/Users/John/Project] %gray -destination %cyan [e.g: C:/Users/John/Backup] %gray -time %cyan [format: d:0 h:0 m:0 s:0]"))
		return
	}

	if _, err := os.Stat(*source); os.IsNotExist(err) {
		fmt.Printf(Color("%red Source: Could not find folder to copy. \n"))
		return
	}

	if _, err := os.Stat(*destination); os.IsNotExist(err) {
		fmt.Println(Color("%red Destination: The target folder to copy to was not found. \n"))
		return
	}

	var b Backup

	// initialization
	b.Source = *source
	b.Destination = *destination
	b.Duration = ParseTime(*duration)
	b.Handler()
}
