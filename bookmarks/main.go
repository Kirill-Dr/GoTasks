package main

import (
	"fmt"
)

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("--- Bookmarks ---")
Menu:
	for {
		action := getMenu()
		switch action {
		case 1:
			getAllBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var choice int
	fmt.Println("Choose action:")
	fmt.Println("1. Get all bookmarks")
	fmt.Println("2. Add bookmark")
	fmt.Println("3. Delete bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return choice
}

func getAllBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks found")
		return
	}
	for key, value := range bookmarks {
		fmt.Println(key, "->", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Println("Enter bookmark key:")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Enter bookmark value:")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkKey string
	fmt.Println("Enter bookmark key to delete:")
	fmt.Scan(&bookmarkKey)
	delete(bookmarks, bookmarkKey)
}
