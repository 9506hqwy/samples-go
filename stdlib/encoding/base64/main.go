package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const ExitErr = 1

func main() {
	filePath := flag.String("file", "", "a json file path.")
	flag.Parse()

	if *filePath == "" {
		flag.PrintDefaults()
		os.Exit(ExitErr)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	encoded, err := encodeFile(file)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Printf("%s\n", *encoded)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := decodeText(encoded)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Printf("%s\n", *contents)
	if err != nil {
		log.Fatal(err)
	}
}

func encodeFile(file *os.File) (*string, error) {
	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(contents)
	return &encoded, nil
}

func decodeText(encoded *string) (*[]byte, error) {
	contents, err := base64.StdEncoding.DecodeString(*encoded)
	if err != nil {
		return nil, err
	}

	return &contents, nil
}
