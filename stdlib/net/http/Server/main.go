package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const ExitErr = 1

const (
	PortMin     = 1025
	PortMax     = 65535
	PortDefault = 8000
)

type indexHandler struct{}

//revive:disable:unused-receiver

func (handler *indexHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, "Hello, World!")
	if err != nil {
		panic(err)
	}
}

//revive:enable:unused-receiver

func main() {
	port := flag.Int("port", PortDefault, "a listen port")
	flag.Parse()

	if *port < PortMin || PortMax < *port {
		flag.PrintDefaults()
		os.Exit(ExitErr)
	}

	http.Handle("/index", &indexHandler{})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
