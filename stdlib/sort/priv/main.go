package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	//revive:disable:add-constant
	array := []int{3, 1, 2}
	//revive:enable:add-constant
	sort.Ints(array)
	_, err := fmt.Println("Array:", array)
	if err != nil {
		log.Fatal(err)
	}

	times := []time.Time{
		time.Now().Add(2 * time.Second),
		time.Now(),
		time.Now().Add(time.Second),
	}
	sort.Slice(times, func(i, j int) bool {
		return times[j].After(times[i])
	})
	_, err = fmt.Println("Time:", times)
	if err != nil {
		log.Fatal(err)
	}
}
