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
	hum := fmt.Sprintf("💧 humidity: \t%d %%", weather.Humidity)
	wind := fmt.Sprintf("💨 wind: \t%.2f Km\\h", weather.Wind.Speed)
	pressure := fmt.Sprintf("📊 pressure: \t%d hPa", weather.Pressure)
	temp := TemperaturePresenter(weather.Temperature)
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", cond, temp, hum, wind, pressure)
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

	return fmt.Sprintf("🌡 temperature: \t%s °C", cur)
}

func ConditionPresenter(condition entity.Condition) string {
	var icon, name string
	switch condition {
	case entity.Thunderstorm:
		icon = "⚡️"
		name = "thunderstorm"
	case entity.Drizzle, entity.Rain:
		icon = "☔️"
		name = "rain"
	case entity.Snow:
		icon = "❄️"
		name = "snow"
	case entity.Clear:
		icon = "☀️"
		name = "clear"
	case entity.Clouds:
		icon = "☁️"
		name = "clouds"
	case entity.Smoke, entity.Haze:
		icon = "🌫"
		name = "haze"
	case entity.Dust:
		icon = "⏳"
		name = "dust"
	case entity.Tornado, entity.Squall:
		icon = "🌪"
		name = "tornado"
	}

	return fmt.Sprintf("%s condition: \t%s", icon, name)
}
