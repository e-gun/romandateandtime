package romandateandtime

import (
	"fmt"
	"github.com/nathan-osman/go-sunrise"
	"math"
	"os"
	"time"
)

func getromantime(p PlaceAndTime) string {
	// hora VII
	rh := fmt.Sprintf("hora %s", integerToRoman(whichhour(p)))
	if isitnight(p) {
		rh += " noctis"
	}
	return rh
}

func lengthofday(p PlaceAndTime) time.Duration {
	// rise  - 2024-07-06 09:43:02 +0000 UTC
	// set   - 2024-07-07 01:01:18 +0000 UTC
	// light - 15h18m16s

	y, m, d := GetYMD(p.T)
	rise, set := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)

	daylight := set.Sub(rise)
	return daylight
}

func lengthofdaytimehour(daylight time.Duration) time.Duration {
	// 15h18m16s --> 1h16m31.333333333s
	return daylight / 12
}

func lengthofnighttimehour(daylight time.Duration) time.Duration {
	// 15h18m16s --> 43m28.666666666s
	dark := (24 * time.Hour) - daylight
	return dark / 12
}

func isitnight(p PlaceAndTime) bool {
	zone, err := time.LoadLocation(TimeZone)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := p.T.In(zone)
	y, m, d := GetYMD(t)
	rise, set := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)

	//fmt.Println(set.In(zone))
	//fmt.Println(rise.In(zone))

	if t.After(set) || t.Before(rise) {
		return true
	} else {
		return false
	}
}

func whichhour(p PlaceAndTime) int {
	if isitnight(p) {
		return whichnightimehour(p)
	} else {
		return whichdaytimehour(p)
	}
}

func whichdaytimehour(p PlaceAndTime) int {
	y, m, d := GetYMD(p.T)
	dur := lengthofdaytimehour(lengthofday(p))
	threshold, _ := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)
	since := p.T.Sub(threshold)
	hour := since / dur

	return int(hour) + 1
}

func whichnightimehour(p PlaceAndTime) int {
	y, m, d := GetYMD(p.T)

	dur := lengthofnighttimehour(lengthofday(p))
	_, setstoday := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)
	_, setyesterday := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d-1)

	var threshold time.Time
	if p.T.After(setstoday) {
		threshold = setstoday
	} else {
		threshold = setyesterday
	}

	since := p.T.Sub(threshold)
	hour := since / dur
	return int(math.Abs(float64(hour))) + 1
}
