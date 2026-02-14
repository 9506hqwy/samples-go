package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
)

const ExitErr = 1

const (
	PortMin     = 1025
	PortMax     = 65535
	PortDefault = 8000
)

func main() {
	port := flag.Int("port", PortDefault, "a listen port")
	flag.Parse()

	if *port < PortMin || PortMax < *port {
		flag.PrintDefaults()
		os.Exit(ExitErr)
	}

	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Fatal(err)
	}

	err = mime.AddExtensionType(".wasm", "application/wasm")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("."))

	ret := http.ListenAndServe(
		fmt.Sprintf(":%d", *port),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			file := r.URL.Path[1:]
			ext := path.Ext(file)
			mtype := mime.TypeByExtension(ext)
			w.Header().Add("Content-Type", mtype)
			fs.ServeHTTP(w, r)
		}))

	log.Fatal(ret)
}
