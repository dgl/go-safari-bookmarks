// Makes something a bit like taskpaper output for your bookmarks...
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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
	dump(bookmark.Bookmark, 0)
}

func dump(b bookmarks.Bookmark, level int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", level*2), b)

	for _, item := range b.Children {
		dump(item, level+1)
	}
}
