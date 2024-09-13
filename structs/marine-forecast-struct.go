package structs

type MarineForecast struct {
	Latitude         float64            `json:"latitude"`              // 42,
	Longitude        float64            `json:"longitude"`             // 41,5
	UTCOffsetSeconds int                `json:"utc_offset_seconds"`    //10800 second
	Timezone         string             `json:"timezone"`              //"Europe/Moscow"
	TimezoneAbbrev   string             `json:"timezone_abbreviation"` //"MSK"
	CurrentUnits     CurrentUnitsMarine `json:"current_units"`
	HourlyUnits      HourlyUnitsMarine  `json:"hourly_units"`
	Current          CurrentDataMarine  `json:"current"`
}
type HourlyDataMarine struct {
	Time                string  `json:"time"`
	WaveHeight          float64 `json:"wave_height"`
	WaveDirection       int     `json:"wave_direction"`
	WavePeriod          float64 `json:"wave_period"`
	WindWaveHeight      float64 `json:"wind_wave_height"`
	WindWaveDirection   int     `json:"wind_wave_direction"`
	WindWavePeriod      float64 `json:"wind_wave_period"`
	WindWavePeakPeriod  float64 `json:"wind_wave_peak_period"`
	SwellWaveHeight     float64 `json:"swell_wave_height"`
	SwellWaveDirection  int     `json:"swell_wave_direction"`
	SwellWavePeriod     float64 `json:"swell_wave_period"`
	SwellWavePeakPeriod float64 `json:"swell_wave_peak_period"`
}

type CurrentDataMarine struct {
	Time                string  `json:"time"`
	WaveHeight          float64 `json:"wave_height"`
	WaveDirection       int     `json:"wave_direction"`
	WavePeriod          float64 `json:"wave_period"`
	WindWaveHeight      float64 `json:"wind_wave_height"`
	WindWaveDirection   int     `json:"wind_wave_direction"`
	WindWavePeriod      float64 `json:"wind_wave_period"`
	WindWavePeakPeriod  float64 `json:"wind_wave_peak_period"`
	SwellWaveHeight     float64 `json:"swell_wave_height"`
	SwellWaveDirection  int     `json:"swell_wave_direction"`
	SwellWavePeriod     float64 `json:"swell_wave_period"`
	SwellWavePeakPeriod float64 `json:"swell_wave_peak_period"`
}

type CurrentUnitsMarine struct {
	Time                string `json:"time"`
	Interval            string `json:"interval"`
	WaveHeight          string `json:"wave_height"`
	WaveDirection       string `json:"wave_direction"`
	WavePeriod          string `json:"wave_period"`
	WindWaveHeight      string `json:"wind_wave_height"`
	WindWaveDirection   string `json:"wind_wave_direction"`
	WindWavePeriod      string `json:"wind_wave_period"`
	WindWavePeakPeriod  string `json:"wind_wave_peak_period"`
	SwellWaveHeight     string `json:"swell_wave_height"`
	SwellWaveDirection  string `json:"swell_wave_direction"`
	SwellWavePeriod     string `json:"swell_wave_period"`
	SwellWavePeakPeriod string `json:"swell_wave_peak_period"`
}

type HourlyUnitsMarine struct {
	Time                string `json:"time"`
	Interval            string `json:"interval"`
	WaveHeight          string `json:"wave_height"`
	WaveDirection       string `json:"wave_direction"`
	WavePeriod          string `json:"wave_period"`
	WindWaveHeight      string `json:"wind_wave_height"`
	WindWaveDirection   string `json:"wind_wave_direction"`
	WindWavePeriod      string `json:"wind_wave_period"`
	WindWavePeakPeriod  string `json:"wind_wave_peak_period"`
	SwellWaveHeight     string `json:"swell_wave_height"`
	SwellWaveDirection  string `json:"swell_wave_direction"`
	SwellWavePeriod     string `json:"swell_wave_period"`
	SwellWavePeakPeriod string `json:"swell_wave_peak_period"`
}
