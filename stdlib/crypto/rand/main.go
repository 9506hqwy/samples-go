package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

const ExitErr = 1

func main() {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Printf("%v\n", bytes)
	if err != nil {
		log.Fatal(err)
	}
}
