package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("docroot"))
  http.Handle("/", fs)

  log.Println("Listening at port 3000...")
  http.ListenAndServe(":3000", nil)
}
