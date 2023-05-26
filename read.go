package bookmarks

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"

	"howett.net/plist"
)

const BookmarksPlist = "Library/Safari/Bookmarks.plist"

var defaultFile string

// Bookmarks is the top level of the tree. It also is a `Bookmark`.
type Bookmarks struct {
	Bookmark
}

type Bookmark struct {
	Children        []Bookmark
	WebBookmarkType string
	Title           string
	URLString       string
	// Mostly has 'title' in it...
	URIDictionary map[string]interface{}
	ReadingList   *ReadingList
}

type ReadingList struct {
	DateAdded      time.Time
	DateLastViewed time.Time
	PreviewText    string
}

func (b Bookmark) String() string {
	title := b.Title
	if len(title) == 0 && b.URIDictionary["title"] != nil {
		title = b.URIDictionary["title"].(string)
	}
	return fmt.Sprintf("%v %v", title, b.URLString)
}

// ReadingItems searches the tree for reading list items and returns them in a list.
func (b Bookmark) ReadingItems() (result []Bookmark) {
	for _, item := range b.Children {
		result = append(result, item.ReadingItems()...)
	}

	if b.ReadingList != nil {
		result = append(result, b)
	}
	return
}

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err) // init, can't do much better...
	}
	defaultFile = u.HomeDir + "/" + BookmarksPlist
}

// Read the default bookmarks plist
func Read() (*Bookmarks, error) {
	return Readfile(defaultFile)
}

// Readfile reads the given bookmarks plist
func Readfile(filename string) (*Bookmarks, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	d := plist.NewDecoder(file)
	var bookmarks Bookmarks
	err = d.Decode(&bookmarks)
	if err != nil {
		return nil, err
	}
	return &bookmarks, nil
}

// RequestFullDiskAccess opens Settings to the Full Disk Access pane.
// Generally you should present something to your user suggesting how they
// should proceed before doing this.
func RequestFullDiskAccess() error {
	return exec.Command("open", "x-apple.systempreferences:com.apple.preference.security?Privacy_AllFiles").Run()
}
