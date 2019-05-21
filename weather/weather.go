package weather

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Weather struct {
	Description string
	Temp        float64
}

type Meteorologist interface {
	getWeather(city string) (Weather, error)
}

func ForecastDownloader(cities []string) []string {
	owm := NewOWM(os.Getenv("API_KEY"))

	var wg sync.WaitGroup
	var description []string

	for _, c := range cities {
		wg.Add(1)
		go func(c string, wg *sync.WaitGroup, description *[]string) {
			defer wg.Done()
			wt, err := owm.GetWeather(c)
			if err != nil {
				log.Println(err.Error())
				return
			}

			*description = append(*description, fmt.Sprintf("In %s weather is %s with a temperature of %.1f °C", c, wt.Description, wt.Temp))
			fmt.Printf("%v", wt)
		}(c, &wg, &description)
	}

	wg.Wait()

	return description
}
