package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	bookmarks "github.com/dgl/go-safari-bookmarks"
)

func main() {
	bookmark, err := bookmarks.Read()
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			log.Print("Full disk access is required. Please enable Full Disk Access for your terminal application in Settings.")
			bookmarks.RequestFullDiskAccess()
		}
		log.Fatal(err)
	}
	rl := bookmark.ReadingItems()
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
