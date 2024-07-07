# romandateandtime

```
import rdt "github.com/e-gun/romandateandtime"

	var pt rdt.PlaceAndTime
	pt.T = time.Now()
	pt.Lat = LATITUDE
	pt.Lon = LONGITUDE
	
    fmt.Println(rdt.GetRomanDateAndTime(pt))

```

if you build as a binary...

```
(romandateandtime: get a roman date and time like "hora III a. d. iv Nonas Ianuarias MMVI")

-h      this help
-d      date with the following layout: "2006-01-02 15:04"
-lg     longitude (default: -79.38)
-lt     latitude (default: 43.65)

```