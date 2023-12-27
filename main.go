package main

import (
	"fmt"
	"log"

	"github.com/onfleet/gonfleet/client"
)

const apiKey = ""

func main() {
	client, err := client.New(apiKey, nil)

	if err != nil {
		log.Fatal(err)
	}

	if client != nil {
		fmt.Println("hay cliente:")
		tasks, error := client.Tasks.Get("FOi8Em0EwqrvxyzVLP9MjGI7")
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(tasks)
	}
}
