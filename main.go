package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type apiConfig struct {
	OpenWeatherApiKey string `json:"OpenWeatherApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius    float64 `json:"temp"`
		Fahrenheit float64 `json:"fahrenheit"`
	}
}

func (w *weatherData) convertToFahrenheit() {
	w.Main.Fahrenheit = w.Main.Celsius*1.8 + 32
}

func readAPIConfig(filename string) (apiConfig, error) {
	var config apiConfig

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func getCity(c echo.Context) error {
	city := c.Param("city")
	data, err := query(city)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to Process Request")
	}
	return c.JSON(http.StatusOK, data)
}

func query(city string) (weatherData, error) {
	var data weatherData

	config, err := readAPIConfig(".api-conf")
	if err != nil {
		return data, errors.New("unable to read config file")
	}

	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + config.OpenWeatherApiKey + "&units=metric"
	resp, err := http.Get(url)
	if err != nil {
		return data, errors.New("unable to get weather data")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, errors.New("unable to decode JSON data")
	}

	data.convertToFahrenheit()

	return data, nil
}

func main() {
	e := echo.New()

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/weather/:city", getCity)

	e.Logger.Fatal(e.Start("localhost:3000"))
}
