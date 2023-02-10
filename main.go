package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var assets embed.FS

func serveFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
	buildFs, err := fs.Sub(assets, "static")
	if err != nil {
		panic(err)
	}

	fsH := http.FileServer(http.FS(buildFs))
	fsH.ServeHTTP(w, req)
}

func main() {

	http.HandleFunc("/", serveFiles)

	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
