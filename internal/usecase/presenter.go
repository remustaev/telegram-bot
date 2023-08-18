package usecase

import (
	"fmt"

	"github.com/remustaev/telegram-bot/internal/entity"
)

// scattered clouds ☁️
// ☀️ temperature: 14.94 °C
// 💧 Humidity: 90%
// 📊 pressure: 1011 hPa
// ☁️ cloudiness: 40%
// 💨 wind: 3.6 Km\h
// 🌅 sunrise: 06:11:16
// 🌄 sunset: 21:21:05

// 08/08/2023:
// ☔️light rain
// 🌡Min: 13.78 - Max: 17.68 °C
//
// 09/08/2023:
// ☔️light rain
// 🌡Min: 12.7 - Max: 19.65 °C
//
// 10/08/2023:
// ☁️scattered clouds
// 🌡Min: 11.92 - Max: 23.24 °C
//
// 11/08/2023:
// ☔️light rain
// 🌡Min: 15.25 - Max: 26.27 °C
//
// 12/08/2023:
// ☔️moderate rain
// 🌡Min: 15.99 - Max: 18.72 °C

// 🌈☀️🌤⛅️🌥☁️🌦🌧⛈🌩🌨❄️💨💧💦☔️☂️🌊🌫🚨🇳
// 🇳🇱

// Example
// {"Condition":1,"Temperature":{"Min":291.44,"Max":292.57,"Current":292.21,"FeelsLike":292.36},
// "Wind":{"Speed":4.12,"Degree":80},"Pressure":1019,"Humidity":84}

func WeatherPresenter(weather entity.Weather) string {
	cond := ConditionPresenter(weather.Condition)
	hum := fmt.Sprintf("💧humidity: %d %%", weather.Humidity)
	temp := TemperaturePresenter(weather.Temperature)
	return fmt.Sprintf("%s\n%s\n%s\n", cond, temp, hum)
}

func TemperaturePresenter(temperature entity.Temperature) string {
	// min := fmt.Sprintf("-%.0f", temperature.Min)
	// if temperature.Min > 0 {
	// 	min = fmt.Sprintf("+%.0f", temperature.Min)
	// }
	//
	// max := fmt.Sprintf("-%.0f", temperature.Max)
	// if temperature.Max > 0 {
	// 	max = fmt.Sprintf("+%.0f", temperature.Max)
	// }

	cur := fmt.Sprintf("-%.0f", temperature.Current)
	if temperature.Current > 0 {
		cur = fmt.Sprintf("+%.0f", temperature.Current)
	}

	return fmt.Sprintf("🌡temperature: %s °C", cur)
}

func ConditionPresenter(condition entity.Condition) string {
	switch condition {
	case entity.Thunderstorm:
		return fmt.Sprintf("⚡️condition: thunderstorm")
	case entity.Drizzle:
		return fmt.Sprintf("💦condition: drizzle")
	case entity.Rain:
		return fmt.Sprintf("☔️condition: rain")
	case entity.Snow:
		return fmt.Sprintf("❄️condition: snow")
	case entity.Clear:
		return fmt.Sprintf("☀️condition: clear")
	case entity.Clouds:
		return fmt.Sprintf("☁️condition: clouds")
	case entity.Smoke:
		return fmt.Sprintf("🌫condition: smoke")
	case entity.Haze:
		return fmt.Sprintf("🌫condition: haze")
	case entity.Dust:
		return fmt.Sprintf("⏳condition: dust")
	case entity.Squall:
		return fmt.Sprintf("💨condition: squall")
	case entity.Tornado:
		return fmt.Sprintf("🌪condition: tornado")
	}
	return ""
}
