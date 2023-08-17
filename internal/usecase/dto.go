package usecase

import (
	"github.com/remustaev/telegram-bot/internal/entity"
)

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type ResponseDTO struct {
	Coordinates Coordinates `json:"coord"`
	Weather     []Weather   `json:"weather"`
	Base        string      `json:"base"`
	Main        Main        `json:"main"`
	Visibility  int64       `json:"visibility"`
	Wind        Wind        `json:"wind"`
	Clouds      Clouds      `json:"clouds"`
	Dt          int64       `json:"dt"`
	Sys         Sys         `json:"sys"`
	Timezone    int64       `json:"timezone"`
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Cod         int64       `json:"cod"`
}

func mapToWeather(respDTO ResponseDTO) entity.Weather {
	return entity.Weather{
		Condition: mapToCondition(respDTO.Cod),
		Temperature: entity.Temperature{
			Min:       respDTO.Main.TempMin,
			Max:       respDTO.Main.TempMax,
			Current:   respDTO.Main.Temp,
			FeelsLike: respDTO.Main.FeelsLike,
		},
		Wind: entity.Wind{
			Speed:  respDTO.Wind.Speed,
			Degree: respDTO.Wind.Deg,
		},
		Pressure: respDTO.Main.Pressure,
		Humidity: respDTO.Main.Humidity,
	}
}

func mapToCondition(condDTO int64) entity.Condition {
	switch {
	case condDTO >= 200 && condDTO <= 299:
		return entity.Thunderstorm
	case condDTO >= 300 && condDTO <= 399:
		return entity.Drizzle
	case condDTO >= 500 && condDTO <= 599:
		return entity.Rain
	case condDTO >= 600 && condDTO <= 699:
		return entity.Snow
	case condDTO == 701 || condDTO == 721 || condDTO == 741:
		return entity.Haze
	case condDTO == 711:
		return entity.Smoke
	case condDTO == 731 || condDTO == 751 || condDTO == 761:
		return entity.Dust
	case condDTO == 771:
		return entity.Squall
	case condDTO == 781:
		return entity.Tornado
	case condDTO == 800:
		return entity.Clear
	case condDTO >= 801 && condDTO <= 809:
		return entity.Clouds
	default:
		return entity.Unknown
	}
}
