package main

import (
	"flag"
	"fmt"
	"github.com/wellnet/corso_go/config"
	"github.com/wellnet/corso_go/weather"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Page struct {
	Title string
	Text  template.HTML
}

type User struct {
	cities []string
}

var (
	u User
)

func init() {
	u = User{
		cities: []string{"L'Aquila", "Potenza", "Catanzaro", "Napoli", "Bologna", "Trieste", "Roma", "Genova", "Milano", "Ancona", "Campobasso", "Torino", "Bari", "Cagliari", "Palermo", "Firenze", "Trento", "Perugia", "Aosta", "Venezia"},
	}
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
	config.Read()

	http.HandleFunc("/", home)
	http.HandleFunc("/logged", logging(home))
	http.HandleFunc("/forecast", forecast)
	http.HandleFunc("/concurrent-forecast", concurrentForecast)
	http.HandleFunc("/config-forecast", configForecast)

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

	err := p.render(w, "templates/page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func forecast(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	forecasts := weather.ForecastDownloader(u.cities)
	log.Printf("Time serial: %v", time.Now().Sub(t))

	var description string
	for _, forecast := range forecasts {
		description += fmt.Sprintf("%s</br>", forecast)
	}

	p := Page{
		Title: "Forecast",
		Text:  template.HTML(description),
	}

	err := p.render(w, "templates/page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func concurrentForecast(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	forecasts := weather.ForecastDownloaderConcurrent(u.cities)
	log.Printf("Time concurrent: %v", time.Now().Sub(t))

	var description string
	for _, forecast := range forecasts {
		description += fmt.Sprintf("%s</br>", forecast)
	}

	p := Page{
		Title: "Forecast",
		Text:  template.HTML(description),
	}

	err := p.render(w, "templates/page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func configForecast(w http.ResponseWriter, r *http.Request) {
	var forecasts []string

	t := time.Now()
	if config.Get().ForecastDownloader == "serial" {
		forecasts = weather.ForecastDownloader(u.cities)
		log.Printf("Time serial: %v", time.Now().Sub(t))
	} else {
		forecasts = weather.ForecastDownloaderConcurrent(u.cities)
		log.Printf("Time concurrent: %v", time.Now().Sub(t))
	}

	var description string
	for _, forecast := range forecasts {
		description += fmt.Sprintf("%s</br>", forecast)
	}

	p := Page{
		Title: "Forecast",
		Text:  template.HTML(description),
	}

	err := p.render(w, "templates/page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
