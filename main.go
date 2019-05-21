package main

import (
	"flag"
	"fmt"
	"github.com/wellnet/corso_go/weather"
	"html/template"
	"log"
	"net/http"
	"os"
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
		cities: []string{"Alessandria", "Torino", "Milano"},
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

	http.HandleFunc("/", home)
	http.HandleFunc("/logged", logging(home))
	http.HandleFunc("/forecast", forecast)
	http.HandleFunc("/concurrent-forecast", concurrentForecast)

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
	city := "Alessandria"

	// e48e7caf838a70d5cfa6fb5f6206013d
	owm := weather.NewOWM(os.Getenv("API_KEY"))
	wt, err := owm.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	p := Page{
		Title: "Forecast",
		Text:  template.HTML(fmt.Sprintf("In %s weather is %s with a temperature of %.1f °C", city, wt.Description, wt.Temp)),
	}

	err = p.render(w, "templates/page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func concurrentForecast(w http.ResponseWriter, r *http.Request) {
	forecasts := weather.ForecastDownloader(u.cities)

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
