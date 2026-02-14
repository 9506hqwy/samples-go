package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

const ExitErr = 1

type InData struct {
	Number int
}

type Data struct {
	Number int
	String string
	Array  []int
	Object InData
}

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

	obj, err := decodeFile(file)
	if err != nil {
		log.Fatal(err)
	}

	err = encodeObj(obj)
	if err != nil {
		log.Fatal(err)
	}
}

//revive:disable:cognitive-complexity

func decodeFile(file *os.File) (*Data, error) {
	decoder := json.NewDecoder(file)

	var obj Data
	err := decoder.Decode(&obj)
	if err != nil {
		return nil, err
	}

	_, err = fmt.Printf("Number = %d\n", obj.Number)
	if err != nil {
		return nil, err
	}

	_, err = fmt.Printf("String = %s\n", obj.String)
	if err != nil {
		return nil, err
	}

	_, err = fmt.Print("Array =")
	if err != nil {
		return nil, err
	}

	for _, i := range obj.Array {
		_, err = fmt.Printf(" %d", i)
		if err != nil {
			return nil, err
		}
	}

	_, err = fmt.Print("\n")
	if err != nil {
		return nil, err
	}

	_, err = fmt.Printf("Object.Number = %d\n", obj.Object.Number)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

//revive:enable:cognitive-complexity

func encodeObj(obj *Data) error {
	encoder := json.NewEncoder(os.Stdout)

	err := encoder.Encode(obj)
	if err != nil {
		return err
	}

	return nil
}
