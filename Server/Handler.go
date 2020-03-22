package Server

import (
	"encoding/json"
	"net/http"

	Api "../WeatherDataApi"
	ServerModel "./Models"
)

func RunServer() {
	http.HandleFunc("/getWeather/", HandleGetWeather)
	http.ListenAndServe(":8080", nil)
}

func HandleGetWeather(w http.ResponseWriter, r *http.Request) {

	location := r.FormValue("location")
	result := Api.GetWeather(location)

	var formattedWeatherData ServerModel.WeatherServerModel
	var weatherData ServerModel.WeatherServerData

	weatherResult := result.List[0]

	weatherData.Temp = weatherResult.Main.Temp
	weatherData.Pressure = weatherResult.Main.Pressure
	weatherData.Description = weatherResult.Weather[0].Description
	weatherData.WindSpeed = weatherResult.Wind.Speed
	weatherData.WindDegree = weatherResult.Wind.Degree

	formattedWeatherData.Message = result.Message
	formattedWeatherData.Data = weatherData

	data, _ := json.Marshal(formattedWeatherData)
	w.Write(data)
}
