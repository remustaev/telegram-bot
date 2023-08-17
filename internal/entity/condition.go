package entity

const (
	Unknown Condition = iota
	Thunderstorm
	Drizzle
	Rain
	Snow
	Clear
	Clouds
	Smoke
	Haze // Туман
	Dust
	Squall // Шквал
	Tornado
)

type Condition int64
