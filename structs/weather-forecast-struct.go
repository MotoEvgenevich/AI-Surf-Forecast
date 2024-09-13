package structs

type WeatherForecast struct {
	Latitude         float64          `json:"latitude"`              //42.0625
	Longitude        float64          `json:"longitude"`             //41.75
	UTCOffsetSeconds int              `json:"utc_offset_seconds"`    //10800
	Timezone         string           `json:"timezone"`              //"Europe/Moscow"
	TimezoneAbbrev   string           `json:"timezone_abbreviation"` //"MSK"
	CurrentUnits     CurrentUnitsWind `json:"current_units"`         // attach struct
	Current          CurrentWind      `json:"current"`               // attach struct
	HourlyUnits      HourlyUnitsWind  `json:"hourly_units"`          // attach struct
	Hourly           HourlyWind       `json:"hourly"`                // attach struct
}

// strict tipization of units for check metric and imperial units system
type CurrentUnitsWind struct {
	Time             string `json:"time"`               // "iso8601"
	Interval         string `json:"interval"`           // "seconds"
	WindSpeed10m     string `json:"wind_speed_10m"`     // km/h speed of wind on latitude 10 meters
	WindDirection10m string `json:"wind_direction_10m"` // degrees
	WindGusts10m     string `json:"wind_gusts_10m"`     // km/h
}

type CurrentWind struct {
	Time             string  `json:"time"`               // "2024-09-13T12:15" format of time
	Interval         int     `json:"interval"`           // interval of gusts per Second
	WindSpeed10m     float64 `json:"wind_speed_10m"`     // km/h
	WindDirection10m int     `json:"wind_direction_10m"` // degrees
	WindGusts10m     float64 `json:"wind_gusts_10m"`     // km/h
}

// strict tipization of units for check metric and imperial units system
type HourlyUnitsWind struct {
	Time             string `json:"time"`               // standart of time "iso8601"
	WindSpeed10m     string `json:"wind_speed_10m"`     // km/h
	WindDirection10m string `json:"wind_direction_10m"` // degrees '°'
	WindGusts10m     string `json:"wind_gusts_10m"`     // km/h
}

type HourlyWind struct {
	Time             []string  `json:"time"`               // []massive of "2024-09-13T00:00"
	WindSpeed10m     []float64 `json:"wind_speed_10m"`     // []massive of 5.2
	WindDirection10m []int     `json:"wind_direction_10m"` // []massive of degrees
	WindGusts10m     []float64 `json:"wind_gusts_10m"`     // []massive of km/h
}
