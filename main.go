package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
	fs := http.FileServer(http.Dir("./static"))
	fs.ServeHTTP(w, req)
}

func main() {

	http.HandleFunc("/", serveFiles)

	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
