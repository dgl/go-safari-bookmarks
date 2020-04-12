package main

import (
  "fmt"
  "log"
  "time"

  bookmarks "github.com/dgl/go-safari-bookmarks"
)

func main() {
  bookmarks, err := bookmarks.Read()
  if err != nil {
    log.Fatal(err)
  }
  rl := bookmarks.ReadingItems()
  for _, item := range rl {
    date := item.ReadingList.DateAdded.Format("2006-01-02")
    done := ""
    var emptyTime time.Time
    if item.ReadingList.DateLastViewed.Equal(emptyTime) {
      done = " @done"
    }
    fmt.Printf("- %v @date(%v)%v\n", item, date, done)
  }
}
