package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define the struct matching the JSON data
type WeatherData struct {
	Latitude             float64      `json:"latitude"`
	Longitude            float64      `json:"longitude"`
	GenerationTimeMs     float64      `json:"generationtime_ms"`
	UtcOffsetSeconds     int          `json:"utc_offset_seconds"`
	Timezone             string       `json:"timezone"`
	TimezoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            float64      `json:"elevation"`
	CurrentUnits         CurrentUnits `json:"current_units"`
	Current              Current      `json:"current"`
	// HourlyUnits          HourlyUnits  `json:"hourly_units"`
	// Hourly               Hourly       `json:"hourly"`
}

type CurrentUnits struct {
	Time                string `json:"time"`
	Interval            string `json:"interval"`
	Temperature2M       string `json:"temperature_2m"`
	RelativeHumidity2M  string `json:"relative_humidity_2m"`
	ApparentTemperature string `json:"apparent_temperature"`
	Precipitation       string `json:"precipitation"`
}

type Current struct {
	Time                string  `json:"time"`
	Interval            int     `json:"interval"`
	Temperature2M       float64 `json:"temperature_2m"`
	RelativeHumidity2M  int     `json:"relative_humidity_2m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	Precipitation       float64 `json:"precipitation"`
}

// type HourlyUnits struct {
// 	Time                     string `json:"time"`
// 	Temperature2M            string `json:"temperature_2m"`
// 	RelativeHumidity2M       string `json:"relative_humidity_2m"`
// 	ApparentTemperature      string `json:"apparent_temperature"`
// 	PrecipitationProbability string `json:"precipitation_probability"`
// 	Precipitation            string `json:"precipitation"`
// 	UvIndex                  string `json:"uv_index"`
// }

// type Hourly struct {
// 	Time                     []string  `json:"time"`
// 	Temperature2M            []float64 `json:"temperature_2m"`
// 	RelativeHumidity2M       []int     `json:"relative_humidity_2m"`
// 	ApparentTemperature      []float64 `json:"apparent_temperature"`
// 	PrecipitationProbability []int     `json:"precipitation_probability"`
// 	Precipitation            []float64 `json:"precipitation"`
// 	UvIndex                  []float64 `json:"uv_index"`
// }

func showCurrentData(data WeatherData) {
	fmt.Println("Location: Mar del Plata, Argentina ", data.Latitude, data.Longitude)
	fmt.Println("Current date and time: ", data.Current.Time)
	fmt.Println("Current temperature: ", data.Current.Temperature2M, data.CurrentUnits.Temperature2M)
	fmt.Println("Current apparent temperature: ", data.Current.ApparentTemperature, data.CurrentUnits.ApparentTemperature)
	fmt.Println("Current relative humidity: ", data.Current.RelativeHumidity2M, data.CurrentUnits.RelativeHumidity2M)
	fmt.Println("Current precipitation", data.Current.Precipitation, data.CurrentUnits.Precipitation)
}

func main() {
	var apiURL = "https://api.open-meteo.com/v1/forecast?latitude=-38.0004&longitude=-57.5562&current=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation&hourly=temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,precipitation,uv_index&timezone=auto"
	// Make HTTP request
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Process JSON response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatal(err)
	}

	showCurrentData(weatherData)
	fmt.Println("Print ANY key to exit")
	fmt.Scanln()
}
