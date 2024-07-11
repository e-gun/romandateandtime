package romandateandtime

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	const (
		HELP = `(romandateandtime: get a roman date and time like "hora III a. d. iv Nonas Ianuarias MMVI")

-h      this help
-d      date with the following layout: "2006-01-02 15:04"
-lg     longitude (default: %.2f)
-lt     latitude (default: %.2f)
`
	)

	args := os.Args[1:len(os.Args)]
	pt := DefaultPT

	for i, a := range args {
		switch a {
		case "-h":
			fmt.Printf(HELP, LONGITUDE, LATITUDE)
			os.Exit(0)
		case "-d":
			d, err := time.Parse("2006-01-02 15:04", args[i+1])
			if err != nil {
				panic(err)
			}
			pt.T = d
		case "-lg":
			lg, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				panic(err)
			}
			pt.Lon = lg
		case "-lt":
			lt, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				panic(err)
			}
			pt.Lat = lt
		}
	}

	fmt.Println("romandateandtime...")
	// fmt.Println(pt.T)
	fmt.Println(pt.GetRomanDateAndTime())
}
