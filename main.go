package main

import (
    "io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo/", func(w http.ResponseWriter, _ *http.Request) {
                             		io.WriteString(w, "Hello-1\n")
    })
	http.HandleFunc("/bar/", func(w http.ResponseWriter, _ *http.Request) {
                             		io.WriteString(w, "Hello-2\n")
    })

	log.Fatal(http.ListenAndServe(":8080", nil))
}