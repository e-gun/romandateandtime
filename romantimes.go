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

	if lengthofday(p) == 24*time.Hour {
		return undef(p)
	}

	rh := fmt.Sprintf("hora %s", integerToRoman(whichhour(p)))
	if isitnight(p) {
		rh += " noctis"
	}
	return rh
}

// undef - the day was 24 hours long; so there are no roman hours
func undef(p PlaceAndTime) string {
	season := "summer"
	if p.T.Month() < 4 || p.T.Month() > 9 {
		season = "winter"
	}

	if p.Lat > 0 && season == "winter" {
		return "(hodie hoc loco solis ortum deest)"
	} else if p.Lat < 0 && season == "winter" {
		return "(hodie hoc loco solis occasum deest)"
	} else if p.Lon > 0 {
		return "(hodie hoc loco solis occasum deest)"
	} else {
		return "(hodie hoc loco solis ortum deest)"
	}
}

func lengthofday(p PlaceAndTime) time.Duration {
	// rise  - 2024-07-06 09:43:02 +0000 UTC
	// set   - 2024-07-07 01:01:18 +0000 UTC
	// light - 15h18m16s

	y, m, d := GetYMD(p.T)
	rise, set := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)

	// 45 minutes of darkness...

	// rise, set := sunrise.SunriseSunset(
	//		69.6, -82.89,
	//		2024, time.July, 25, // 2000-01-01
	//	)
	// rise: 2024-07-25 05:56:22 +0000 UTC     set: 2024-07-26 05:19:30 +0000 UTC

	// if the sun is up all day you get
	// rise: 0001-01-01 00:00:00 +0000 UTC     set: 0001-01-01 00:00:00 +0000 UTC

	if set == rise {
		// we actually lost the ability to answer the isitnight() question...
		return 24 * time.Hour
	}

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
	if p.TZ == "" {
		p.TZ = DefaultTimeZone
	}

	zone, err := time.LoadLocation(p.TZ)
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

func lthour(p PlaceAndTime) float64 {
	y, m, d := GetYMD(p.T)
	dur := lengthofdaytimehour(lengthofday(p))
	threshold, _ := sunrise.SunriseSunset(p.Lat, p.Lon, y, m, d)
	since := p.T.Sub(threshold)
	hour := since.Seconds() / dur.Seconds()
	return hour
}

func dkhour(p PlaceAndTime) float64 {
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
	hour := since.Seconds() / dur.Seconds()
	return hour
}

func whichdaytimehour(p PlaceAndTime) int {
	hour := lthour(p)
	return int(hour) + 1
}

func thisdaytimehourpctelapsed(p PlaceAndTime) float64 {
	hour := lthour(p)
	return float64(float64(hour) - float64(int(hour)))
}

func whichnightimehour(p PlaceAndTime) int {
	hour := dkhour(p)
	return int(math.Abs(float64(hour))) + 1
}

func thisnighttimehourpctelapsed(p PlaceAndTime) float64 {
	hour := dkhour(p)
	return float64(float64(hour) - float64(int(hour)))
}
