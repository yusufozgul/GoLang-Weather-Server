package models

type WeatherBase struct {
	Message string        `json:"message"`
	Cod     string        `json:"cod"`
	Count   int           `json:"count"`
	List    []WeatherData `json:"list"`
}
type WeatherData struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Coord   Coordinates      `json:"coord"`
	Main    WeatherMain      `json:"main"`
	Wind    WeatherWind      `json:"wind"`
	Weather []WeatherDetails `json:"weather"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type WeatherMain struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type WeatherWind struct {
	Speed  float32 `json:"speed"`
	Degree int     `json:"deg"`
}

type WeatherDetails struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
