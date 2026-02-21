package main

import (
	"fmt"
	"log"
	"time"
)

func parseTime() error {
	t1, err := time.Parse(
		time.RFC3339,
		"2026-02-21T23:01:23+09:00")
	if err != nil {
		return err
	}

	// 2026-02-21 23:01:41 +0900 +0900
	_, err = fmt.Println(t1)
	if err != nil {
		return err
	}

	// format: year hour second AM/PM
	t2, err := time.Parse(
		"2006 3 5 PM",
		"2026 2 5 AM")
	if err != nil {
		return err
	}

	// 2026-01-01 02:00:05 +0000 UTC
	_, err = fmt.Println(t2)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := parseTime()
	if err != nil {
		log.Fatal(err)
	}
}
