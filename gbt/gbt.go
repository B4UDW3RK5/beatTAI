package main

import (
   "fmt"
   "time"
)

func main() {
   // Variable to store the number of leap seconds
   leapSeconds := 37

   // Get the current time in UTC
   currentTime := time.Now().UTC()

   // Extract the hour, minute, and second from the current time
   hour := currentTime.Hour()
   minute := currentTime.Minute()
   second := currentTime.Second()

   // Print the values
   fmt.Println("Current UTC Time:")
   fmt.Println("Hour:", hour)
   fmt.Println("Minute:", minute)
   fmt.Println("Second:", second)
   fmt.Println("leapSeconds:", leapSeconds)
}
