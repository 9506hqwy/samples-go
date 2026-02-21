package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	_, err := fmt.Println(time.Now(), "Timer start")
	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Second * 2)
	<-timer.C

	_, err = fmt.Println(time.Now(), "Timer expired")
	if err != nil {
		log.Fatal(err)
	}
}
