// Makes something a bit like taskpaper output for your bookmarks...
package main

import (
	"fmt"
	"log"
	"strings"

	bookmarks "github.com/dgl/go-safari-bookmarks"
)

func main() {
	bookmarks, err := bookmarks.Read()
	if err != nil {
		log.Fatal(err)
	}
	dump(bookmarks.Bookmark, 0)
}

func dump(b bookmarks.Bookmark, level int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", level*2), b)

	for _, item := range b.Children {
		dump(item, level+1)
	}
}
