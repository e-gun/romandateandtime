# romandateandtime

## as a binary 
```
% ./romandateandtime
hora IV a. d. iv Sextilis MMXXIV

% ./romandateandtime -h
(romandateandtime: get a roman date and time like "hora III a. d. iv Nonas Ianuarias MMVI")

-h      this help
-d      date with the following layout: "2006-01-02 15:04"
-lg     longitude (default: -79.38)
-lt     latitude (default: 43.65)
 
```

## as a libarary...

```
import rdt "github.com/e-gun/romandateandtime"

	var pt rdt.PlaceAndTime
	pt.T = time.Now()
	pt.Lat = LATITUDE
	pt.Lon = LONGITUDE
	
	rdt.Timezone = "America/New_York"
	
    fmt.Println(pt.GetRomanDateAndTime())

```

## clock

a live 'multi-timezone' Roman clock is (probably) running at https://antisigma.classics.utoronto.ca/t.
