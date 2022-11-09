package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Username string `json:username"`
	Password string `json:password"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	var C Config
	err = json.Unmarshal([]byte(file), &C)

	var username, password string
	fmt.Printf("Enter Username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter Password: ")
	fmt.Scanln(&password)

	if username != C.Username || password != C.Password {
		log.Println("Incorrect Login Info!")
		os.Exit(0)
	} else {
		log.Println("Success!")
	}
}
