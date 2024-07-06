package main

import (
	"fmt"
	"strings"
	"time"
)

func GetRomanDateAndTime(p PlaceAndTime) string {
	// hora VII pridie Nonas Quinctilis MMXXIV
	return GetRomanTime(p) + " " + GetRomanDate(p.T)
}

func GetRomanTime(p PlaceAndTime) string {
	// hora VII
	rh := fmt.Sprintf("hora %s", integerToRoman(whichhour(p)))
	if isitnight(p) {
		rh += " noctis"
	}
	return rh
}

func GetRomanDate(t time.Time) string {
	// pridie Nonas Quinctilis MMXXIV
	y, m, d := getymd(t)

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
