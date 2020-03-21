package WeatherDataApi

import (
	"encoding/json"
	"fmt"

	models "./models"
)

func GetWeather(location string) models.WeatherBase {

	var weatherData models.WeatherBase
	responseString := GetData(location)

	error := json.Unmarshal([]byte(responseString), &weatherData)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(weatherData)
	return weatherData
}
