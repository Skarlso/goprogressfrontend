package main

import (
	"flag"
	"fmt"
)

var location string
var apiver string

func displayOptions() {
    fmt.Println("Please choose from the following options:")
	fmt.Println("1. Create Character")
	fmt.Println("2. Play")
	fmt.Println("3. Quit")
}

func init() {
    const (
        defaultLocation = "http://localhost:8989"
        defaultAPIVersion = "1"
    )
	flag.StringVar(&location, "-l", defaultLocation, "The location of GoProgressQuest.")
    flag.StringVar(&apiver, "-api", defaultAPIVersion, "The API version you are playing against.")
}

func main() {
	choice := 0
	fmt.Println("Welcome To GoFront.")
    fmt.Println("Server location:", location)
    fmt.Println("API version:", apiver)
	for choice != 3 {
		displayOptions()
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("Error scanning for choice: %v\n", err)
		}

		fmt.Println("Scanned choice:", choice)
	}
}
