package main

import "net/http"
import "log"

import "fmt"
import "time"
import "html"

func main() {
  s := &http.Server{
  	Addr:           ":8080",
  	ReadTimeout:    10 * time.Second,
  	WriteTimeout:   10 * time.Second,
  	MaxHeaderBytes: 1 << 20,
  }
  http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })
  http.ListenAndServe(s)
  log.Printf("starting server on port :8080")
}
