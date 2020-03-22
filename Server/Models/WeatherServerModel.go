package Models

type WeatherServerModel struct {
	Message string            `json:"message"`
	Data    WeatherServerData `json:"data"`
}

type WeatherServerData struct {
	Temp        float64 `json:"temp"`
	Pressure    float64 `json:"pressure"`
	Description string  `json:"description"`
	WindSpeed   float32 `json:"windSpeed"`
	WindDegree  int     `json:"windDeg"`
}
