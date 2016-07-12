package main

import (
  "net/http"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc  http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "PostIndex",
        "GET",
        "/posts",
        PostIndex,
    },
    Route{
        "PostShow",
        "GET",
        "/posts/{postId}",
        PostShow,
    },
    Route{
      "PostCreate",
      "POST",
      "/posts",
      PostCreate,
    },
}
