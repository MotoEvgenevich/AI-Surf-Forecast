package main

import (
	"ai-forecast/structs"
	"fmt"
	"time"
)

type WeatherConditions struct {
	Sea  structs.MarineForecast
	Wind structs.WeatherForecast
}

var DataWeather []WeatherConditions

func main() {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02T")
	currentDateSlice := make([]string, 0)
	h := 0
	hour := "00-00"
	for i := 0; i < 41; i++ {
		date := currentDate + hour
		currentDateSlice = append(currentDateSlice, date)

		h += 3
		if h == 24 { // Когда h достигает 24, происходит смена даты
			h = 0
			currentTime = currentTime.Add(time.Hour * 24) // Увеличиваем дату на 1 день
			currentDate = currentTime.Format("2006-01-02T")
		}

		hour = fmt.Sprintf("%02d-00", h)
	}

	// Печать всех дат и времени
	for _, v := range currentDateSlice {
		fmt.Println(v)
	}

	// Добавление данных о погоде в слайс структур
	for _, v := range currentDateSlice {
		// Заполняем поле времени в MarineForecast и WeatherForecast
		sea := structs.MarineForecast{
			Current: structs.CurrentDataMarine{
				Time: v, // Добавляем дату и время в поле Time
			},
		}
		wind := structs.WeatherForecast{
			Current: structs.CurrentWind{
				Time: v, // Добавляем дату и время в поле Time
			},
		}

		// Добавляем структуру в слайс DataWeather
		DataWeather = append(DataWeather, WeatherConditions{
			Sea:  sea,
			Wind: wind,
		})
	}

	// Печать содержимого DataWeather для проверки
	for _, data := range DataWeather {
		fmt.Printf("Sea Time: %s, Wind Time: %s\n", data.Sea.Current.Time, data.Wind.Current.Time)
	}
}
