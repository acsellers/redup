package main

import "flag"
import "fmt"
import "github.com/simonz05/godis/redis"
import "html"
import "log"
import "net/http"

var PortNum = flag.Int("port", 5790, "Port to view redup on")

func StartHttpServer() {
	http.HandleFunc("/assets/", AssetHandler)
	http.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *PortNum), nil))
}

func AssetHandler(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/assets/bootstrap.css":
		w.Header().Set("Content-Type", "text/css")
		fmt.Fprint(w, BOOTSTRAPCSS)
	case "/assets/bootstrap.js":
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Fprint(w, BOOTSTRAPJS)
	case "/assets/jquery.js":
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Fprint(w, JQUERYJS)
	case "/assets/application.js":
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Fprint(w, APPLICATIONJS)
	case "/assets/application.css":
		w.Header().Set("Content-Type", "text/css")
		fmt.Fprint(w, APPLICATIONCSS)
	default:
		w.WriteHeader(404)
	}
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	navbar := fmt.Sprintf(NAVBAR, NavbarItems("One", "Two"))
	content := fmt.Sprintf(CONTENT, "Title", LeftBar(), ContentArea())
	fmt.Fprintf(w, BASE, navbar, content)
}

func NavbarItems(items ...string) (markedUpItems string) {
	for _, item := range items {
		markedUpItems += fmt.Sprintf(NAVBARITEM, html.EscapeString(item))
	}
	return
}

func LeftBar() string {
	barHtml := `<div class="well">
  <table class="table flat-bottom">
  <thead><tr><th class="all_keys">All Keys</th></tr></thead>
  <tbody>`
	keys, err := Conn.Keys("*")
	if err == nil {
		for _, key := range keys {
			barHtml += fmt.Sprintf("<tr><td><a href=\"#\" class=\"key-link\" data-key=\"%v\">%v</a></td></tr>", key, key)
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
			contentArea += ContentFor(key) + "<hr>"
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
				return fmt.Sprintf(KEYCONTENT, key, key, v.String())
			}
			return fmt.Sprintf(ERRORCONTENT, key, key, e)
		case "list":
			list, e := Conn.Lrange(key, 0, -1)
			if e == nil {
				return ListContent(key, list)
			}
			return fmt.Sprintf(ERRORCONTENT, key, key, e)
		case "set":
			list, e := Conn.Smembers(key)
			if e == nil {
				return ListContent(key, list)
			}
			return fmt.Sprintf(ERRORCONTENT, key, key, e)
		default:
			return fmt.Sprintf(KEYCONTENT, key, key, info)
		}
	}
	return fmt.Sprintf(KEYCONTENT, key, key, err)
}

func ListContent(key string, list *redis.Reply) string {
	listContent := fmt.Sprintf(LISTCONTENTHEAD, key, key)
	for _, v := range list.StringArray() {
		listContent += fmt.Sprintf("<li><pre class=\"redis-item\">%s</pre></li>", v)
	}

	return listContent + LISTCONTENTTAIL
}
