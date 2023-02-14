package main

import (
    "io"
	"log"
	"net/http"
	"fmt"
	"html/template"
)

func addAction(w http.ResponseWriter, r *http.Request) {
  e := r.ParseForm()
  log.Println(e)
  data := r.Form.Get("data")
  store := newMemoryDb()
  store.add(Todo(data))
  fmt.Println(store.store)
  io.WriteString(w, data)
}

func indexAction(w http.ResponseWriter, r *http.Request) {
  t := template.Must(template.ParseFiles("public/index.html",))

  cookie := &http.Cookie{
            Name:    "SESID",
            Value:   "384279583",
            }

    http.SetCookie(w, cookie)
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", indexAction)

    http.HandleFunc("/add", addAction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}