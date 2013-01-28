package main

import "flag"
import "fmt"
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
		fmt.Fprint(w, "")
	case "/assets/application.css":
		w.Header().Set("Content-Type", "text/css")
		fmt.Fprint(w, "")
	default:
		w.WriteHeader(404)
	}
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	navbar := fmt.Sprintf(NAVBAR, NavbarItems("One", "Two"))
	content := fmt.Sprintf(CONTENT, "Contect Here")
	fmt.Fprintf(w, BASE, navbar, content)
}

func NavbarItems(items ...string) (markedUpItems string) {
	for _, item := range items {
		markedUpItems += fmt.Sprintf(NAVBARITEM, html.EscapeString(item))
	}
	return
}
