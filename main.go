package main

import (
    "io"
	"log"
	"net/http"
)
　
func parseRequest(w http.ResponseWriter, r *http.Request) {
  e := r.ParseForm()
  log.Println(e)
  data := r.Form.Get("data")
  io.WriteString(w, data)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("assets/")))

    http.HandleFunc("/add", parseRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}