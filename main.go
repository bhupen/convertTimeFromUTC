package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	toZone := flag.String("zone", "America/Chicago", "time zone to convert to (America/Chicago)")
	flag.Parse()

	layout := "2006-01-02 15:04:05"
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
