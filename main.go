package main

import (
    "crypto/rand"
    "math/big"
    "os"
    "io"
	"log"
	"net/http"
	"fmt"
	"html/template"
	"strconv"
)

func createSessionId() int64 {
  n, err := rand.Int(rand.Reader, big.NewInt(100000000000000))
  if err != nil {
      panic(err)
  }
  return n.Int64()
}

func addAction(w http.ResponseWriter, r *http.Request) {
  e := r.ParseForm()
  data := r.Form.Get("data")
  store := newMemoryDb()
  store.add(Todo(data))
  fmt.Println(store.store)
  io.WriteString(w, data)
}

func indexAction(w http.ResponseWriter, r *http.Request) {
  createSessionId()
  t := template.Must(template.ParseFiles("public/index.html",))

  cookie, _ := r.Cookie("SESID")
  if cookie == nil {
      sessionId := strconv.FormatInt(createSessionId(), 10)
      cookie := &http.Cookie{
          Name: "SESID",
          Value: sessionId,
       }
      fp, err := os.Create("./tmp/" + sessionId + ".txt")
      if err != nil {
          log.Println(err)
          return
      }
      defer fp.Close()


   http.SetCookie(w, cookie)
  }

    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", indexAction)

    http.HandleFunc("/add", addAction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}