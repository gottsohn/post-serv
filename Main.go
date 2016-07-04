package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
)

func main() {
  router := NewRouter()
  bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
  fmt.Printf("Serving on port %s", bind)
  err := http.ListenAndServe(bind, router)
  if err != nil {
      log.Fatal(err)
  }
}
