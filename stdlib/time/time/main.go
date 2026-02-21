package main

import (
	"fmt"
	"log"
	"time"
)

func printDateTime(t time.Time) error {
	_, err := fmt.Println(t)
	if err != nil {
		return err
	}

	err = printDate(t)
	if err != nil {
		return err
	}

	err = printTime(t)
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Nanosecond())
	if err != nil {
		return err
	}

	// UTC
	_, err = fmt.Println(t.Location())
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Weekday())
	if err != nil {
		return err
	}

	return nil
}

func printDate(t time.Time) error {
	_, err := fmt.Println(t.Year())
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Month())
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Day())
	if err != nil {
		return err
	}

	return nil
}

func printTime(t time.Time) error {
	_, err := fmt.Println(t.Hour())
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Minute())
	if err != nil {
		return err
	}

	_, err = fmt.Println(t.Second())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	now := time.Now()

	_, err := fmt.Println(now)
	if err != nil {
		log.Fatal(err)
	}

	//revive:disable:add-constant
	then := time.Date(2026, 2, 21, 01, 21, 58, 123456789, time.UTC)
	//revive:enable:add-constant

	err = printDateTime(then)
	if err != nil {
		log.Fatal(err)
	}
}
