package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const ExitErr = 1

func main() {
	filepath := flag.String("file", "", "a zip file path.")
	flag.Parse()

	if *filepath == "" {
		flag.PrintDefaults()
		os.Exit(ExitErr)
	}

	reader, err := zip.OpenReader(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	err = listFile(reader)
	if err != nil {
		log.Fatal(err)
	}
}

func listFile(reader *zip.ReadCloser) error {
	for _, file := range reader.File {
		if !strings.HasSuffix(file.Name, "/") {
			_, err := fmt.Printf("Name: %s(%d/%d)\n",
				file.Name,
				file.CompressedSize64,
				file.UncompressedSize64)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
