package weather

type Weather struct {
	Description string
	Temp float64
}

type Meteorologist interface {
	getWeather(city string) (Weather, error)
}
