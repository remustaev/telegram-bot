package entity

const (
	LatitudeAmsterdam  = 52.377956
	LongitudeAmsterdam = 4.897070
	TimezoneCET        = "CET"
)

type Timezone string

type Location struct {
	Latitude  float32
	Longitude float32
	Timezone  Timezone
}

func NewLocationAmsterdam() Location {
	return Location{
		Latitude:  LatitudeAmsterdam,
		Longitude: LongitudeAmsterdam,
		Timezone:  TimezoneCET,
	}
}
