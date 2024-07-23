package romandateandtime

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type RomanMonth struct {
	Name  string
	Abbr  string
	Days  int
	Nones int
	Ides  int
}

var (
	RomanYear = buildromanyear()
)

func getromandate(t time.Time) string {
	// pridie Nonas Quinctilis MMXXIV
	y, m, d := GetYMD(t)

	leap := false
	if y%400 == 0 || (y%100 != 0 && y%4 == 0) {
		leap = true
	}

	rm := RomanYear[int(m)]

	postides := func(d int) string {
		var rd string
		var ante int
		if leap && m == time.February && d == 24 {
			ante = rm.Days - (d - 2)
			if d == 25 {
				rd = "bis-"
			}
		} else {
			ante = rm.Days - (d - 1)
		}

		if ante > 1 {
			rd += "a. d. " + strings.ToLower(integerToRoman(ante))
		} else if ante == 1 {
			rd += "pridie Kalendas"
		} else {
			// do nothing
		}
		return rd
	}

	// walk through the month ...
	var rd string
	if d == 1 {
		rd = "Kalendis"
	} else if d < rm.Nones-1 {
		an := rm.Nones - (d - 1)
		rd = fmt.Sprintf("a. d. %s Nonas", strings.ToLower(integerToRoman(an)))
	} else if d == rm.Nones-1 {
		rd = "pridie Nonas"
	} else if d == rm.Nones {
		rd = "Nonis"
	} else if d < rm.Ides-1 {
		an := rm.Ides - (d - 1)
		rd = fmt.Sprintf("a. d. %s Idus", strings.ToLower(integerToRoman(an)))
	} else if d == rm.Ides-1 {
		rd = "pridie Idus"
	} else if d == rm.Ides {
		rd = "Idibus"
	} else {
		rd = postides(d)
	}

	return rd + " " + rm.Name + " " + integerToRoman(y)
}

func buildromanyear() map[int]RomanMonth {
	ry := map[int]RomanMonth{
		1:  {"Ianuarias", "Ian.", 31, 5, 13},
		2:  {"Februarias", "Feb.", 28, 5, 13},
		3:  {"Martias", "Mar.", 31, 7, 15},
		4:  {"Aprilis", "Apr.", 30, 5, 13},
		5:  {"Maias", "Mai.", 31, 7, 15},
		6:  {"Iunias", "Iun.", 30, 5, 13},
		7:  {"Quinctilis", "Quint.", 31, 7, 15},
		8:  {"Sextilis", "Sex.", 31, 5, 13},
		9:  {"Septembris", "Sept.", 30, 5, 13},
		10: {"Octobris", "Oct.", 31, 7, 15},
		11: {"Novembris", "Nov.", 30, 5, 13},
		12: {"Decembris", "Dec.", 31, 5, 13},
	}
	return ry
}

func GetYMD(t time.Time) (int, time.Month, int) {
	y := t.Year()
	m := t.Month()
	d := t.Day()
	return y, m, d
}

func testdates() {
	// 2024-07-06 14:14:05.209947 -0400 EDT m=+0.003046251
	// 15h18m16s
	// pridie Nonas Quinctilis MMXXIV
	// 7
	// hora VII pridie Nonas Quinctilis MMXXIV
	// 2024-07-04 05:33:00 -0400 EDT
	// 12
	// hora XII noctis a. d. iv Nonas Quinctilis MMXXIV

	zone, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tt := time.Date(2024, time.July, 4, 5, 33, 0, 0, zone)
	pt := PlaceAndTime{
		T:   tt,
		Lat: LONGITUDE,
		Lon: LATITUDE,
	}
	fmt.Println(time.Now())
	fmt.Println(lengthofday(DefaultPT))
	fmt.Println(getromandate(time.Now()))
	fmt.Println(whichhour(DefaultPT))
	fmt.Println(DefaultPT.GetRomanDateAndTime())

	fmt.Println(tt)
	fmt.Println(whichhour(pt))
	fmt.Println(pt.GetRomanDateAndTime())
}
