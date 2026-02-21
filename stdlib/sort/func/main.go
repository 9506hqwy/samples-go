package main

import (
	"fmt"
	"log"
	"sort"
)

type Length []string

func (s Length) Len() int {
	return len(s)
}

func (s Length) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Length) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	lengths := []string{"ccc", "a", "bb"}

	sort.Sort(Length(lengths))

	_, err := fmt.Println(lengths)
	if err != nil {
		log.Fatal(err)
	}
}
