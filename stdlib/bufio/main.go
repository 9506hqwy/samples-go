package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const ExitErr = 1
const LineInit = 1
const BufSize = 32

func main() {
	filePath := flag.String("file", "", "a file path")
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

	// default size = 4k
	reader := bufio.NewReaderSize(file, BufSize)

	err = showFile(reader)
	if err != nil {
		log.Fatal(err)
	}
}

//revive:disable:cognitive-complexity

func showFile(reader *bufio.Reader) error {
	lineCount := LineInit
	prefix := true
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if prefix {
			_, err = fmt.Printf("%03d: ", lineCount)
			if err != nil {
				return err
			}

			prefix = false
		}

		_, err = fmt.Printf("%s", line)
		if err != nil {
			return err
		}

		if isPrefix {
			continue
		}

		_, err = fmt.Print("\n")
		if err != nil {
			return err
		}

		prefix = true
		lineCount++
	}

	return nil
}

//revive:enable:cognitive-complexity
