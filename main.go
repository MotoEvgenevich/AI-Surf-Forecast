package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02T")
	fmt.Println(currentDate)
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

	for _, v := range currentDateSlice {
		fmt.Println(v)
	}
	fmt.Println(len(currentDateSlice))
}
