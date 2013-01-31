package main

import (
	"fmt"
	"net/http"
)

func KeyHandler(w http.ResponseWriter, req *http.Request) {
	navbar := fmt.Sprintf(NAVBAR, NavbarItems("One", "Two"))
	key := FindKey(req.URL.Path)
	if key != "" {
		content := fmt.Sprintf(FULLCONTENT, key, KeyContent(key))
		fmt.Fprintf(w, BASE, navbar, content)
	} else {
		fmt.Fprintf(w, BASE, navbar, fmt.Sprintf(FULLCONTENT, "Key not found", ""))
	}
}
func KeyContent(key string) string {
	return ContentFor(key)
}
