package main

import (
	"fmt"
	"log"
	"time"
)

func formatTime() error {
	now := time.Now()

	_, err := fmt.Println(now.Format(time.UnixDate))
	if err != nil {
		return err
	}

	_, err = fmt.Println(now.Format(time.RFC822Z))
	if err != nil {
		return err
	}

	_, err = fmt.Println(now.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func formatPartial() error {
	now := time.Now()

	_, err := fmt.Println(now.Format("2006/1/_2 Mon 3:04:5.000"))
	if err != nil {
		return err
	}

	err = formatPartialDate(now)
	if err != nil {
		return err
	}

	err = formatPartialTime(now)
	if err != nil {
		return err
	}

	_, err = fmt.Println("AM/PM:", now.Format("PM"))
	if err != nil {
		return err
	}

	return nil
}

func formatPartialDate(now time.Time) error {
	_, err := fmt.Println("Year:", now.Format("2006"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Month:", now.Format("01"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Day of the week:", now.Format("Mon"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Day of the month:", now.Format("2"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Day of the year:", now.Format("002"))
	if err != nil {
		return err
	}

	return nil
}

func formatPartialTime(now time.Time) error {
	_, err := fmt.Println("Hour:", now.Format("3"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Minute:", now.Format("4"))
	if err != nil {
		return err
	}

	_, err = fmt.Println("Second:", now.Format("5"))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := formatTime()
	if err != nil {
		log.Fatal(err)
	}

	err = formatPartial()
	if err != nil {
		log.Fatal(err)
	}
}
