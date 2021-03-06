package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
		defaultLocation   = "http://localhost:8989"
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
		switch choice {
		case 1:
			registerCharacter()
		case 2:
			play()
		}
	}
}

func registerCharacter() {
	var name string
	fmt.Println("Please enter name of the character:")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error reading character name: ", err)
		return
	}
	sendMessageToServer("create", name)
}

func play() {
	choice := 0
	name := ""
    stop := make(chan bool, 1)
	for choice != 3 {
		fmt.Println("1. Start adventure")
		fmt.Println("2. Stop adventure")
		fmt.Println("3. Back")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Print("Name of the character:")
			fmt.Scan(&name)
			sendMessageToServer("start", name)
			go displayInformation(name, stop)
		} else if choice == 2 {
			fmt.Print("Name of the character:")
			fmt.Scan(&name)
			sendMessageToServer("stop", name)
			select {
			case stop <- true:
			default:
			}
		}
	}
}

func displayInformation(name string, signal chan bool) {
    var stop bool
	for {
        select {
        case stop = <-signal:
        default:
        }

        if stop {
            break
        }

		resp, err := http.Get(location + "/api/" + apiver + "/load/" + name)
		if err != nil {
			fmt.Println("There was an error reading from the server:", err)
			return
		}

		defer resp.Body.Close()
		var p Character
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&p)
		fmt.Printf("Name: %s; Level: %d; Money: %d; Hp:%d;\n", p.Name, p.Level, p.Gold, p.Hp)
		// body,_ := json.MarshalIndent(p.Body, "", " ")
		// fmt.Printf("Body: %s;\n\n", body)
		stats,_ := json.MarshalIndent(p.Stats, "", " ")
		fmt.Printf("Stats: %s;\n\n", stats)
		time.Sleep(time.Millisecond * 1000)
	}
}

func sendMessageToServer(uri, name string) {
	type Post struct {
		Name string `json:"name"`
	}
	post := Post{Name: name}
	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error while marshalling post:", err)
	}

	resp, err := http.Post(location+"/api/"+apiver+"/"+uri, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("There was an error reading from the server:", err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response from the server: %s\n", body)
}
