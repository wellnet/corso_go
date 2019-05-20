package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/wellnet/corso_go/weather"
)

type Page struct {
	Title string
	Text  string
}

func (p Page) render(w http.ResponseWriter, filename string) error {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	port := flag.Int("port", 80, "webserver port")

	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/logged", logging(home))
	http.HandleFunc("/forecast", forecast)

	fmt.Printf("Start server on port %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", *port), nil)
	if err != nil {
		panic(err)
	}
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "Page title",
		Text:  "Lorem ipsum",
	}

	err := p.render(w, "templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func forecast(w http.ResponseWriter, r *http.Request) {
	city := "Alessandria"

	owm := weather.NewOWM("e48e7caf838a70d5cfa6fb5f6206013d")
	wt, err := owm.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	p := Page{
		Title: "Forecast",
		Text:  fmt.Sprintf("In %s weather is %s with a temperature of %.1f °C", city, wt.Description, wt.Temp),
	}

	err = p.render(w, "templates/forecast.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
