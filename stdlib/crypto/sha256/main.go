package main

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	s := "Hello, World!"
	h := sha256.New()

	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	bin := h.Sum(nil)

	// [
	//  223 253 96 33 187 43 213 176 175 103 98 144 128 158 195 165
	//  49 145 221 129 199 247 10 75 40 104 138 54 33 130 152 111
	// ]
	_, err = fmt.Println(bin)
	if err != nil {
		log.Fatal(err)
	}

	// dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f
	_, err = fmt.Printf("%x\n", bin)
	if err != nil {
		log.Fatal(err)
	}
}
