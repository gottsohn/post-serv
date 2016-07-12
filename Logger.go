package main

import (
  "fmt"
  "net/http"
  "log"
  "time"
)

func LogMessage(w http.ResponseWriter, message string) {
  fmt.Fprintf(w, message)
}

func LogError(err error) {
  log.Fatal(err)
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
