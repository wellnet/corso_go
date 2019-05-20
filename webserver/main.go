package main

import (
	"flag"
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Text string
}

func main() {
	port := flag.Int("port", 80, "webserver port")

	flag.Parse()

	http.HandleFunc("/", home)

	fmt.Printf("Start server on port %d", *port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", *port), nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = tmpl.Execute(w, Page{
		Title: "Page title",
		Text: "Lorem ipsum",
	})
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
