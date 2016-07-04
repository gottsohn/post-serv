package main
import (
  "time"
)

type Post struct {
  Title string      `json:"title"`
  Images []string   `json:"images"`
  Text string       `json:"text"`
  Read bool         `json:"read"`
  Date time.Time    `json:"date"`
}

type Posts []Post
