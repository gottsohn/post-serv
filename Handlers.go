package main
import (
  "log"
  "fmt"
  "strconv"
  "runtime"
  "net/http"
  "encoding/json"
  "io"
  "io/ioutil"
  "github.com/gorilla/mux"
)

var posts = Posts{
    Post{
      Id: 1,
      Title: "Dross 1",
      Images:[]string{"http://remikuti.com/logo.png"},
      Text:"Some more Dross",
      Read:false,
    },

    Post{
      Id: 2,
      Title: "Dross 2",
      Images:[]string{"http://remikuti.com/logo.png"},
      Text:"Some more Dross2",
      Read:false,
    },
}

func PostIndex (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(posts); err != nil {
    log.Fatal(err)
  }
}

func FindPost(id int64) Post {
  for _, post := range posts {
    if post.Id == id {
      return post
    }
  }

  return Post{
    Id: -1,
  }
}

func PostShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  postId, err := strconv.ParseInt(vars["postId"], 10, 64)
  if err != nil {
      fmt.Fprintf(w, "Error parsing ID %s", vars["postId"])
  }

  post := FindPost(postId)
  if post.Id != -1 {
    if err1 := json.NewEncoder(w).Encode(post); err1 != nil {
        LogError(err1)
    }
  } else {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Post with id: %i was not found", postId)
  }
}

func RepoCreatePost(post Post) Post {
  nweId := len(posts) + 1
  post.Id = int64(nweId)
  posts = append(posts, post)
  return post
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
  var post Post
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    LogError(err)
  }

  if err := r.Body.Close(); err != nil {
       LogError(err)
  }

  if err = json.Unmarshal(body, &post); err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      LogError(err)
    }

    t := RepoCreatePost(post)
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
      LogError(err)
    }
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Willkommen zu Schmetterling, rene aus ", runtime.Version())
}
