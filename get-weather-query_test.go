package main

import (
	"fmt"
	"net/url"
	"testing"
)

type WeatherForecastParams struct {
	Latitude  float64
	Longitude float64
	Current   string
	Hourly    string
	Timezone  string
}

type MarineForecastParams struct {
	Latitude  float64
	Longitude float64
	Current   string
	Hourly    string
	Daily     string
	Timezone  string
}

func GenerateWeatherForecastURL(params WeatherForecastParams) string {
	baseURL := "https://api.open-meteo.com/v1/forecast"
	u, _ := url.Parse(baseURL)

	query := url.Values{}
	query.Add("latitude", fmt.Sprintf("%.5f", params.Latitude))
	query.Add("longitude", fmt.Sprintf("%.5f", params.Longitude))
	query.Add("current", params.Current)
	query.Add("hourly", params.Hourly)
	query.Add("timezone", params.Timezone)

	u.RawQuery = query.Encode()
	return u.String()
}

func GenerateMarineForecastURL(params MarineForecastParams) string {
	baseURL := "https://marine-api.open-meteo.com/v1/marine"
	u, _ := url.Parse(baseURL)

	query := url.Values{}
	query.Set("latitude", fmt.Sprintf("%.2f", params.Latitude))
	query.Set("longitude", fmt.Sprintf("%.2f", params.Longitude))

	query.Set("current", params.Current)
	query.Set("hourly", params.Hourly)
	query.Set("daily", params.Daily)
	query.Set("timezone", params.Timezone)

	u.RawQuery = query.Encode()
	return u.String()
}

func TestQuery(t *testing.T) {
	weatherForecast := WeatherForecastParams{
		Latitude:  42.13076,
		Longitude: 41.65917,
		Current:   "wind_speed_10m,wind_direction_10m,wind_gusts_10m",
		Hourly:    "wind_speed_10m,wind_direction_10m,wind_gusts_10m",
		Timezone:  "Europe/Moscow",
	}

	marineForecast := MarineForecastParams{
		Latitude:  42.0625,
		Longitude: 41.75,
		Current:   "wave_height,wave_direction,wave_period,wind_wave_height,wind_wave_direction,wind_wave_period,wind_wave_peak_period,swell_wave_height,swell_wave_direction,swell_wave_period,swell_wave_peak_period",
		Hourly:    "wave_height,wave_direction,wave_period,wind_wave_height,wind_wave_direction,wind_wave_period,wind_wave_peak_period,swell_wave_height,swell_wave_direction,swell_wave_period,swell_wave_peak_period",
		Daily:     "wave_height_max,wave_period_max,wind_wave_height_max,wind_wave_period_max,wind_wave_peak_period_max,swell_wave_height_max,swell_wave_direction_dominant,swell_wave_period_max,swell_wave_peak_period_max",
		Timezone:  "Europe/Moscow",
	}
	weatherURL := GenerateWeatherForecastURL(weatherForecast)
	forecastWind, err := getWindWeatherForecast(weatherURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Weather Forecast URL:", weatherURL)
	fmt.Printf("Время: %v\n", forecastWind.Current.Time)
	fmt.Printf("Скорость ветра: %v %v\n", forecastWind.Current.WindSpeed10m, forecastWind.CurrentUnits.WindSpeed10m)
	fmt.Printf("Порывы ветра: %v %v\n", forecastWind.Current.WindGusts10m, forecastWind.CurrentUnits.WindGusts10m)

	marineURL := GenerateMarineForecastURL(marineForecast)
	forecastMarine, err := getMarineForecast(marineURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marine Forecast URL:", marineURL)
	fmt.Printf("Время: %v\n", forecastMarine.Current.Time)
	fmt.Printf("Широта: %v\n Долгота: %v\n", forecastMarine.Latitude, forecastMarine.Longitude)
	fmt.Printf("Высота свеволовой волны: %v %v\n", forecastMarine.Current.SwellWaveHeight, forecastMarine.CurrentUnits.SwellWaveHeight)
}
