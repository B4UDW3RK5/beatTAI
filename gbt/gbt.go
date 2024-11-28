package main

import (
	"fmt"
	"time"
)

func main() {
	// variable to store the number of leap seconds. this is really lazy but
	// later I'll find a way to snag the current number of leap seconds, until
	// then this will need to be manually changed any time a new leap second is
	// added.
	leapSeconds := 37

	// grab the current time in UTC
	currentTime := time.Now().UTC()

	// pull the hour, minute, and second from the current time and convert the
	// hour and minute to seconds
	hourSeconds := currentTime.Hour() * 3600
	minuteSeconds := currentTime.Minute() * 60
	seconds := currentTime.Second()

	// do the math for beatTAI and store it
	beatTAI := float64((hourSeconds + minuteSeconds + seconds + leapSeconds)) / 86.4

	// print the value, rounded to two decimal places, padded with leading zeroes
	// if necessary and prefixed with : for proper beatTAI syntax
	fmt.Printf(":%06.2f\n", beatTAI)
}
