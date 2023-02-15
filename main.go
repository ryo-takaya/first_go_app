package main

import (
    "crypto/rand"
    "math/big"
    "os"
    "io"
	"log"
	"net/http"
	"html/template"
	"strconv"
	"io/ioutil"
	"strings"
	"fmt"
)

func createSessionId() int64 {
      n, err := rand.Int(rand.Reader, big.NewInt(100000000000000))
      if err != nil {
          panic(err)
      }
      return n.Int64()
}

func indexAction(w http.ResponseWriter, r *http.Request) {
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
        t.Execute(w, nil)
    } else {
        data, _ := ioutil.ReadFile("./tmp/" + cookie.Value + ".txt")
        if data == nil {
            t.Execute(w, nil)
            return
        }
        todoList := strings.Split(string(data), "\n")
        t.Execute(w, todoList)
  }
}

func addAction(w http.ResponseWriter, r *http.Request) {
      r.ParseForm()
      requestData := r.Form.Get("data")
      sessionId, _ := r.Cookie("SESID")
      fp, _ := os.OpenFile("./tmp/" + sessionId.Value + ".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
      defer fp.Close()
      fmt.Fprintln(fp, requestData)
      io.WriteString(w, requestData)
}

func main() {
    http.HandleFunc("/", indexAction)

    http.HandleFunc("/add", addAction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}