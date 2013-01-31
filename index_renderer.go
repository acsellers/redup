package main

import (
	"fmt"
	"github.com/simonz05/godis/redis"
	"net/http"
)

const (
	LEFTBARHEAD = `<div class="well">
  <table class="table flat-bottom">
  <thead><tr><th class="all_keys">All Keys</th></tr></thead>
  <tbody>`
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	navbar := fmt.Sprintf(NAVBAR, NavbarItems("One", "Two"))
	content := fmt.Sprintf(CONTENT, "All Keys", LeftBar(), ContentArea())
	fmt.Fprintf(w, BASE, navbar, content)
}

func LeftBar() string {
	barHtml := LEFTBARHEAD
	keys, err := AllKeys()
	if err == nil {
		for _, key := range keys {
			barHtml += KeyListItem(key)
		}
		barHtml += "</tbody></table></div>"

		return barHtml
	}
	return fmt.Sprintf("<div class=\"well\">%v</div", err)
}

func ContentArea() string {
	contentArea := ""
	keys, err := Conn.Keys("*")
	if err == nil {
		for _, key := range keys {
			contentArea += ContentFor(key) + "<hr class='seperator'>"
		}
	}

	return contentArea
}

func ContentFor(key string) string {
	info, err := Conn.Type(key)
	if err == nil {
		switch info {
		case "string":
			v, e := Conn.Get(key)
			if e == nil {
				return fmt.Sprintf(KEYCONTENT, KeyToHtmlId(key), KeyToHtmlId(key), key, v.String())
			}
			return fmt.Sprintf(ERRORCONTENT, KeyToHtmlId(key), key, e)
		case "list":
			list, e := Conn.Lrange(key, 0, -1)
			if e == nil {
				return ListContent(key, list)
			}
			return fmt.Sprintf(ERRORCONTENT, KeyToHtmlId(key), key, e)
		case "set":
			list, e := Conn.Smembers(key)
			if e == nil {
				return ListContent(key, list)
			}
			return fmt.Sprintf(ERRORCONTENT, KeyToHtmlId(key), key, e)
		default:
			return fmt.Sprintf(KEYCONTENT, KeyToHtmlId(key), KeyToHtmlId(key), key, info)
		}
	}
	return fmt.Sprintf(KEYCONTENT, KeyToHtmlId(key), KeyToHtmlId(key), key, err)
}

func ListContent(key string, list *redis.Reply) string {
	listContent := fmt.Sprintf(LISTCONTENTHEAD, KeyToHtmlId(key), KeyToHtmlId(key), key)
	for _, v := range list.StringArray() {
		listContent += fmt.Sprintf("<li><pre class=\"redis-item\">%s</pre></li>", v)
	}

	return listContent + LISTCONTENTTAIL
}
