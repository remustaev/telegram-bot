package entity

type Weather struct {
	Condition   Condition
	Temperature Temperature
	Wind        Wind
	Pressure    int64
	Humidity    int64 // влажность
}
