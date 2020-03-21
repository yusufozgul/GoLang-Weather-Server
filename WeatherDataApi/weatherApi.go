package WeatherDataApi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData(location string) string {

	if !checkString(location) {
		return ""
	}

	url := "https://community-open-weather-map.p.rapidapi.com/find?type=link%252C%20accurate&units=mecric%252C%20metric&q=" + location

	req, error := http.NewRequest("GET", url, nil)
	checkError(error)

	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "****************")

	res, error := http.DefaultClient.Do(req)
	checkError(error)

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	checkError(error)

	return string(body)
}
func checkString(key string) bool {
	if len(key) > 0 {
		return true
	}
	return false
}

func checkError(error error) {
	if error != nil {
		fmt.Println(error)
	}
}
