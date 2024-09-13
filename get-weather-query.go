package main

import (
	"ai-forecast/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var forecastWind structs.WeatherForecast
var forecastMarine structs.MarineForecast

func main() {
	forecastWind, err := getWindWeatherForecast()
	if err != nil {
		log.Fatalf("Ошибка получения прогноза: %v", err)
	}
	forecastMarine, err := getMarineForecast()
	if err != nil {
		log.Fatalf("Ошибка получения прогноза: %v", err)
	}
	// print wind example
	fmt.Printf("Время: %v\n", forecastWind.Current.Time)
	fmt.Printf("Скорость ветра: %v %v\n", forecastWind.Current.WindSpeed10m, forecastWind.CurrentUnits.WindSpeed10m)
	fmt.Printf("Порывы ветра: %v %v\n", forecastWind.Current.WindGusts10m, forecastWind.CurrentUnits.WindGusts10m)
	// print marine example
	fmt.Printf("Время: %v\n", forecastMarine.Current.Time)
	fmt.Printf("Высота свеволовой волны: %v %v\n", forecastMarine.Current.SwellWaveHeight, forecastMarine.CurrentUnits.SwellWaveHeight)
}

func getWindWeatherForecast() (structs.WeatherForecast, error) {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=42.13076&longitude=41.65917&current=wind_speed_10m,wind_direction_10m,wind_gusts_10m&hourly=wind_speed_10m,wind_direction_10m,wind_gusts_10m&timezone=Europe%2FMoscow")
	if err != nil {
		return structs.WeatherForecast{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return structs.WeatherForecast{}, fmt.Errorf("неожиданный статус код: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&forecastWind); err != nil {
		return structs.WeatherForecast{}, fmt.Errorf("ошибка парсинга JSON: %d", err)
	}
	return forecastWind, nil
}

func getMarineForecast() (structs.MarineForecast, error) {
	resp, err := http.Get("https://marine-api.open-meteo.com/v1/marine?latitude=42.0625&longitude=41.75&current=wave_height,wave_direction,wave_period,wind_wave_height,wind_wave_direction,wind_wave_period,wind_wave_peak_period,swell_wave_height,swell_wave_direction,swell_wave_period,swell_wave_peak_period&hourly=wave_height,wave_direction,wave_period,wind_wave_height,wind_wave_direction,wind_wave_period,wind_wave_peak_period,swell_wave_height,swell_wave_direction,swell_wave_period,swell_wave_peak_period&daily=wave_height_max,wave_period_max,wind_wave_height_max,wind_wave_period_max,wind_wave_peak_period_max,swell_wave_height_max,swell_wave_direction_dominant,swell_wave_period_max,swell_wave_peak_period_max&timezone=Europe%2FMoscow")
	if err != nil {
		return structs.MarineForecast{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return structs.MarineForecast{}, fmt.Errorf("неожиданный статус код: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&forecastMarine); err != nil {
		return structs.MarineForecast{}, fmt.Errorf("ошибка парсинга JSON: %d", err)
	}
	return forecastMarine, nil
}
