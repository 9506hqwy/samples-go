package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public/*
var public embed.FS

func main() {
	file, err := fs.Sub(public, "public")
	if err != nil {
		panic(err)
	}

	root := http.FS(file)
	http.Handle("/", http.FileServer(root))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
