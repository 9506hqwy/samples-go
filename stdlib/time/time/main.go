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

	_, err = fmt.Printf("Nanosecond() = %d\n", t.Nanosecond())
	if err != nil {
		return err
	}

	// UTC
	_, err = fmt.Printf("Location() = %s\n", t.Location())
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Weekday() = %s (%d)\n", t.Weekday(), t.Weekday())
	if err != nil {
		return err
	}

	return nil
}

func printDate(t time.Time) error {
	_, err := fmt.Printf("Year() = %d\n", t.Year())
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Month() = %d\n", t.Month())
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Day() = %d\n", t.Day())
	if err != nil {
		return err
	}

	return nil
}

func printTime(t time.Time) error {
	_, err := fmt.Printf("Day() = %d\n", t.Hour())
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Minute() = %d\n", t.Minute())
	if err != nil {
		return err
	}

	_, err = fmt.Printf("Second() = %d\n", t.Second())
	if err != nil {
		return err
	}

	return nil
}

func calcDateTime(t time.Time) error {
	//revive:disable:add-constant
	aday := t.AddDate(0, 0, 1)
	//revive:enable:add-constant
	_, err := fmt.Printf("1 day after = %v\n", aday)
	if err != nil {
		return err
	}

	//revive:disable:add-constant
	bday := t.AddDate(0, 0, -2)
	//revive:enable:add-constant
	_, err = fmt.Printf("2 day before = %v\n", bday)
	if err != nil {
		return err
	}

	//revive:disable:add-constant
	dayCount := (time.Saturday - t.Weekday() + 7) % 7
	if dayCount == 0 {
		dayCount = 7
	}
	nearSat := t.AddDate(0, 0, int(dayCount))
	//revive:enable:add-constant
	_, err = fmt.Printf(
		"Near Next Saturday = %v (%s)\n",
		nearSat,
		nearSat.Weekday(),
	)
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

	err = calcDateTime(then)
	if err != nil {
		log.Fatal(err)
	}
}
