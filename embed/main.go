package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/*
var dist embed.FS

func main() {
	file, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	root := http.FS(file)
	http.Handle("/", http.FileServer(root))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
