package main

import (
  "fmt"
  "github.com/gorilla/websocket"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello AGAIN from inside Go Browser")
}
