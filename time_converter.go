package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	amPm := s[len(s)-2:]
	hour, _ := strconv.Atoi(s[0:2])
	convertedHour := hour
	if amPm == "PM" {
		if convertedHour != 12 {
			convertedHour += 12
		}
	} else if convertedHour == 12 {
		convertedHour = 0
	}
	convertedHourString := strconv.Itoa(convertedHour)
	if convertedHour <= 9 {
		convertedHourString = strings.Join([]string{"0", convertedHourString}, "")
	}
	return strings.Join([]string{convertedHourString, s[2 : len(s)-2]}, "")
}

func main() {
	//s := "12:05:39AM"
	s := "11:45:54PM"
	result := timeConversion(s)
	fmt.Printf("%s\n", result)
}
