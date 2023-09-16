package main

import "fmt"

func main() {
	const secondsInMin = 60
	const minsInHour = 60
	const secondsInHour = secondsInMin * minsInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)
}
