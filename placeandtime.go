package romandateandtime

import (
	"time"
)

const (
	LATITUDE  = 43.65
	LONGITUDE = -79.38
)

var (
	DefaultTimeZone = "America/New_York"
	DefaultPT       = PlaceAndTime{
		T:   time.Now(),
		Lat: LATITUDE,
		Lon: LONGITUDE,
	}
)

type PlaceAndTime struct {
	T   time.Time
	Lat float64
	Lon float64
	TZ  string // e.g., "America/New_York"
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

// GetRTRemainder - how long is a daytime hour at p?
func (p *PlaceAndTime) GetLenOfDaytimeHour() time.Duration {
	// 1h16m1.583333333s
	return lengthofdaytimehour(lengthofday(*p))
}

// GetRTRemainder - how much time is left in this Roman hour before the next hour arrives
func (p *PlaceAndTime) GetRTRemainder() time.Duration {
	// 22m26s
	var h time.Duration
	var r float64
	if isitnight(*p) {
		h = lengthofnighttimehour(lengthofday(*p))
		r = thisnighttimehourpctelapsed(*p)
	} else {
		h = lengthofdaytimehour(lengthofday(*p))
		r = thisdaytimehourpctelapsed(*p)
	}

	elapsed := h.Seconds() * r
	left := h.Seconds() - elapsed
	return time.Duration(left) * time.Second
}
