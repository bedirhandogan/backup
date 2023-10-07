package main

import (
	"flag"
	"fmt"
	"github.com/bedirhandogan/color"
	"os"
)

func main() {
	source := flag.String("source", "", "Path to the folder to copy.")
	destination := flag.String("destination", "", "The path to the target folder.")
	duration := flag.String("duration", "", "Backup duration of the program.")
	flag.Parse()

	if *source == "" || *destination == "" || *duration == "" {
		fmt.Println(color.Colorize("%yellow Use: %reset  go run . %white40 -source %cyan [e.g: C:/Users/John/Project] %white40 -destination %cyan [e.g: C:/Users/John/Backup] %white40 -duration %cyan [format: d:0 h:0 m:0 s:0] %reset"))
		return
	}

	if _, err := os.Stat(*source); os.IsNotExist(err) {
		fmt.Printf(color.Colorize("%red Source: Could not find folder to copy. \n %reset"))
		return
	}

	if _, err := os.Stat(*destination); os.IsNotExist(err) {
		fmt.Println(color.Colorize("%red Destination: The target folder to copy to was not found. \n %reset"))
		return
	}

	var b Backup

	// initialization
	b.Source = *source
	b.Destination = *destination
	b.Duration = ParseTime(*duration)
	b.Handler()
}
