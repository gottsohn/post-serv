package main

import (
  "fmt"
  "net/http"
  "log"
)

func LogMessage(w http.ResponseWriter, message string) {
  fmt.Fprintf(w, message)
}

func LogError(err error) {
  log.Fatal(err)
}
