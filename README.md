# romandateandtime

```
import rdt "github.com/e-gun/romandateandtime"

	var pt rdt.PlaceAndTime
	pt.T = time.Now()
	pt.Lat = LATITUDE
	pt.Lon = LONGITUDE
	
	rdt.Timezone = "America/New_York"
	
    fmt.Println(pt.GetRomanDateAndTime())

```

if you build as a binary...

```
(romandateandtime: get a roman date and time like "hora III a. d. iv Nonas Ianuarias MMVI")

-h      this help
-d      date with the following layout: "2006-01-02 15:04"
-lg     longitude (default: -79.38)
-lt     latitude (default: 43.65)

```

```

     25 Jul 24 12:42 EDT
     hora VI a. d. vii Quinctilis MMXXIV

Comparanda as one moves north from the equator
	Quito
	Miami
	Toronto
	Chisasibi
Local time                  11:42:00 	12:42:00    12:42:00 	12:42:00
Length of daylight hour     01:00:32 	01:07:13    01:14:02 	01:20:26
Roman hour                   hora VI 	 hora VI     hora VI 	 hora VI
Time left in this hour      00:38:27 	00:45:08    00:41:54 	00:39:58
Longitude                     -78.52      -80.19      -79.38      -78.90
Latitude                       -0.23       25.77       43.65       53.78

```