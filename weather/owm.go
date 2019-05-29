package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OWMResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64     `json:"pressure"`
		Humidity float64     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type OWM struct {
	baseURL string
	key     string
}

func NewOWM(key string) *OWM {
	return &OWM{
		baseURL: "https://api.openweathermap.org/data/2.5/weather",
		key:     key,
	}
}

func (owm *OWM) GetWeather(city string) (Weather, error) {
	r, err := http.Get(fmt.Sprintf("%s?q=%s,it&appid=%s&units=metric", owm.baseURL, city, owm.key))
	if err != nil {
		return Weather{}, err
	}

	defer r.Body.Close()

	v := &OWMResponse{}
	err = json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return Weather{}, err
	}

	return Weather{
		Description: v.Weather[0].Description,
		Temp:        v.Main.Temp,
	}, nil
}
