package main
import (
  "time"
)

type Post struct {
  Id int64            `json:"id"`
  Title string      `json:"title"`
  Images []string   `json:"images"`
  Text string       `json:"text"`
  Read bool         `json:"read"`
  Date time.Time    `json:"date"`
}

type Posts []Post
