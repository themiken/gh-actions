package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Hello, " + currentUser.Username + "!")
}
