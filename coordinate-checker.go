package main

import (
	"ai-forecast/structs"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type KilometerDeviation struct {
	Latitude  float64
	Longitude float64
}

const oneKmInDeg float64 = 0.00898
const earthRadiusAtEquator = 111.32

func CoordinateChecker(coords Coordinate) (KilometerDeviation, error) {
	fmt.Println("Входные координаты:", coords)

	// Формируем URL для запроса к API
	url := fmt.Sprintf("https://marine-api.open-meteo.com/v1/marine?latitude=%f&longitude=%f&hourly=wave_height", coords.Latitude, coords.Longitude)
	fmt.Println("URL запроса:", url)

	// Отправляем HTTP GET запрос
	resp, err := http.Get(url)
	if err != nil {
		return KilometerDeviation{}, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		return KilometerDeviation{}, fmt.Errorf("неожиданный статус код: %d", resp.StatusCode)
	}

	// Парсим JSON ответ
	var testCoordinate structs.MarineForecast
	if err := json.NewDecoder(resp.Body).Decode(&testCoordinate); err != nil {
		return KilometerDeviation{}, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	fmt.Printf("Ответные координаты: Широта: %f, Долгота: %f\n", testCoordinate.Latitude, testCoordinate.Longitude)

	// Вычисляем отклонение по широте
	deviationLatitude := coords.Latitude - testCoordinate.Latitude
	deviationLatitudeKilometer := deviationLatitude / oneKmInDeg
	fmt.Printf("Отклонение по широте: %f градусов, что соответствует %f км\n", deviationLatitude, deviationLatitudeKilometer)

	// Вычисляем длину одного градуса долготы в километрах на данной широте
	lengthOfDegreeLongitude := ConverterLongitudeToKilometer(testCoordinate.Latitude)

	// Вычисляем отклонение по долготе
	deviationLongitude := coords.Longitude - testCoordinate.Longitude
	deviationLongitudeKilometer := deviationLongitude * lengthOfDegreeLongitude
	fmt.Printf("Отклонение по долготе: %f градусов, что соответствует %f км\n", deviationLongitude, deviationLongitudeKilometer)

	result := KilometerDeviation{
		Latitude:  deviationLatitudeKilometer,
		Longitude: deviationLongitudeKilometer,
	}

	return result, nil
}

func ConverterLongitudeToKilometer(latitude float64) float64 {
	latitudeRad := latitude * math.Pi / 180
	cosLatitude := math.Cos(latitudeRad)

	lengthOfDegreeLongitude := earthRadiusAtEquator * cosLatitude
	fmt.Printf("Длина одного градуса долготы на широте %.2f°: %.2f км\n", latitude, lengthOfDegreeLongitude)
	return lengthOfDegreeLongitude
}
