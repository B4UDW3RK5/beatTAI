package main

import (
	"fmt"

	"github.com/brandondube/tai"
)

func main() {
	g := tai.Now().AsGregorian()

	// pull the hour, minute, and second from the TAI time and convert the
	// hour and minute to seconds
	hourSeconds := g.Hour * 3600
	minuteSeconds := g.Min * 60
	seconds := g.Sec

	// do the math for beatTAI and store it
	beatTAI := float64(hourSeconds+minuteSeconds+seconds) / 86.4

	// print the value, rounded to two decimal places, padded with leading zeroes
	// if necessary and prefixed with : for proper beatTAI syntax
	fmt.Printf(":%06.2f\n", beatTAI)
}
