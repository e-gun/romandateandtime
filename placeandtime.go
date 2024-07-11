package romandateandtime

import (
	"time"
)

const (
	LATITUDE  = 43.65
	LONGITUDE = -79.38
)

var (
	TimeZone  = "America/New_York"
	DefaultPT = PlaceAndTime{
		T:   time.Now(),
		Lat: LATITUDE,
		Lon: LONGITUDE,
	}
)

type PlaceAndTime struct {
	T   time.Time
	Lat float64
	Lon float64
}

func (p *PlaceAndTime) GetRomanDateAndTime() string {
	// hora VII pridie Nonas Quinctilis MMXXIV
	return p.GetRomanTime() + " " + p.GetRomanDate()
}

func (p *PlaceAndTime) GetRomanTime() string {
	// hora VII
	return getromantime(*p)
}

func (p *PlaceAndTime) GetRomanDate() string {
	// pridie Nonas Quinctilis MMXXIV
	return getromandate(p.T)
}

func (p *PlaceAndTime) GetLenOfDaytimeHour() time.Duration {
	// 1h16m1.583333333s
	return lengthofdaytimehour(lengthofday(*p))
}
