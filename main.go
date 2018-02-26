package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	toZone := flag.String("zone", "US/Central", "time zone to convert to (US/Central)")
	flag.Parse()

	layout := "2006-01-02 15:04:05"

	if len(flag.Args()) < 1 {
		fmt.Println(`
	Usage : convertTimeFromUTC <--zone=""> <YYYY-MM-DD HH:mm::ss>
	Optional zone: timezone to convert to (default: US/Central)
			`)
		os.Exit(0)
	}

	str := flag.Args()[0]
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var loc *time.Location

	loc, err = time.LoadLocation(*toZone)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(t.In(loc))
}
