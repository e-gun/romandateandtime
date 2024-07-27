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

// GetRomanDateAndTime - 25 Jul 24 10:09 EDT -->  hora IV a. d. vii Quinctilis MMXXIV (in Toronto...)
func (p *PlaceAndTime) GetRomanDateAndTime() string {
	// hora VII pridie Nonas Quinctilis MMXXIV
	return p.GetRomanTime() + " " + p.GetRomanDate()
}

// GetRomanTime - 10:09 EDT -->  hora IV (in Toronto...)
func (p *PlaceAndTime) GetRomanTime() string {
	// hora VII
	return getromantime(*p)
}

// GetRomanDate - a. d. vii Quinctilis MMXXIV
func (p *PlaceAndTime) GetRomanDate() string {
	// pridie Nonas Quinctilis MMXXIV
	return getromandate(p.T)
}

// GetRTRemainder - how long is a daytime hour for this PlaceAndTime? e.g., 1h16m1.583333333s
func (p *PlaceAndTime) GetLenOfDaytimeHour() time.Duration {
	// 1h16m1.583333333s
	return lengthofdaytimehour(lengthofday(*p))
}

// GetRTRemainder - how much time is left in this Roman hour at this PlaceAndTime before the next hour arrives? e.g. 22m26s
func (p *PlaceAndTime) GetRTRemainder() time.Duration {
	// 22m26s
	elapsed := p.GetRTElapsed()
	var h time.Duration
	if isitnight(*p) {
		h = lengthofnighttimehour(lengthofday(*p))
	} else {
		h = lengthofdaytimehour(lengthofday(*p))
	}
	left := h - elapsed
	return left
}

// GetRTElapsed - how much time has passed in this Roman hour at this PlaceAndTime? e.g. 22m26s
func (p *PlaceAndTime) GetRTElapsed() time.Duration {
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
	return time.Duration(elapsed) * time.Second
}

func (p *PlaceAndTime) GetArabicHour() int {
	return whichhour(*p)
}
