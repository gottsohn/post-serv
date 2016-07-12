package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
)

func main() {
  router := NewRouter()
  port := os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  
  bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), port)
  fmt.Printf("Serving on port:%s...", bind)
  err := http.ListenAndServe(bind, router)
  if err != nil {
      log.Fatal(err)
  }
}
