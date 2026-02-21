package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func output() error {
	echoCmd := exec.Command("echo", "Hello, world!")

	echoOut, err := echoCmd.Output()
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(echoOut))
	if err != nil {
		return err
	}

	return nil
}

func pipe() error {
	catCmd := exec.Command("cat")

	catIn, err := catCmd.StdinPipe()
	if err != nil {
		return err
	}

	catOut, err := catCmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = catCmd.Start()
	if err != nil {
		return err
	}

	err = writeStdin(catIn, []byte("Hello, World!"))
	if err != nil {
		return err
	}

	stdout, err := io.ReadAll(catOut)
	if err != nil {
		return err
	}

	err = catCmd.Wait()
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(stdout))
	if err != nil {
		return err
	}

	return nil
}

func writeStdin(stdin io.WriteCloser, input []byte) error {
	_, err := stdin.Write(input)
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := output()
	if err != nil {
		log.Fatal(err)
	}

	err = pipe()
	if err != nil {
		log.Fatal(err)
	}
}
