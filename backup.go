package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TimePeriod struct {
	Day, Hour, Min, Sec int
}

type Backup struct {
	Source      string
	Destination string
	Duration    TimePeriod
}

func CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func CopyDirectory(source, destination string) error {
	return filepath.Walk(source, func(sourcePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(destination, strings.TrimPrefix(sourcePath, source))

		if info.IsDir() {
			return os.MkdirAll(destinationPath, os.ModePerm)
		}

		return CopyFile(sourcePath, destinationPath)
	})
}

func (b *Backup) Handler() {
	// When the program runs
	fmt.Println(Color("%bcyanThe program was run."))

	// Source and destination path
	fmt.Printf(Color("%bcyan Source: %white %s \n"), b.Source)
	fmt.Printf(Color("%bcyan Destination: %white %s \n"), b.Destination)

	// Backup Duration
	fmt.Printf(Color("%bcyan Duration: %white Day: %d, Hour: %d, Min: %d, Sec: %d \n"), b.Duration.Day, b.Duration.Hour, b.Duration.Min, b.Duration.Sec)

	TaskScheduler(func() error {
		now := time.Now()
		d := now.Day() + b.Duration.Day
		h := now.Hour() + b.Duration.Hour
		m := now.Minute() + b.Duration.Min
		s := now.Second() + b.Duration.Sec

		if err := CopyDirectory(b.Source, b.Destination); err != nil {
			fmt.Println(Color("%red An error occurred while copying. \n"))
			return err
		}

		fmt.Printf(Color("%green Your folder has been backed up successfully. Next Backup: %white %02d:%02d:%02d - %d.%d.%d \n"),
			h, m, s, d, time.Now().Month(), time.Now().Year())

		return nil
	}, b.Duration.Day, b.Duration.Hour, b.Duration.Min, b.Duration.Sec)
}
