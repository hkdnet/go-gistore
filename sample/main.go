package main

import (
	"fmt"
	"os"

	"github.com/hkdnet/go-gistore"
)

func main() {
	client := gistore.NewClient("9740d2de8850351f730c6ed851a550b0")
	var ret []Bookmark
	if err := client.SelectAll("bookmarks", &ret); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, v := range ret {
		fmt.Printf("%2d [%s : %s](%s)\n", v.ID, v.Title, v.Description, v.URL)
	}
}

type Bookmark struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
