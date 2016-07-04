package main
import (
  "log"
  "fmt"
  "strconv"
  "runtime"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

var posts = Posts{
    Post{Title: "Dross 1", Images:[]string{"http://remikuti.com/logo.png"},Text:"Some more Dross", Read:false},
    Post{Title: "Dross 2", Images:[]string{"http://remikuti.com/logo.png"},Text:"Some more Dross2", Read:false},
}

func PostIndex (w http.ResponseWriter, r *http.Request) {
  if err := json.NewEncoder(w).Encode(posts); err != nil {
    log.Fatal(err)
  }
}

func PostShow (w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  postId, err := strconv.ParseInt(vars["postId"], 10, 64)
  if err != nil {
      fmt.Fprintf(w, "Error parsing ID %s", vars["postId"])
  }

  post := posts[postId]
  // if post != nil {
  //   // LogError(err)
  //   LogMessage(w, "Index doesn's exist")
  // }

  if err1 := json.NewEncoder(w).Encode(post); err1 != nil {
      LogError(err1)
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to Schmetterling, running on ", runtime.Version())
}
