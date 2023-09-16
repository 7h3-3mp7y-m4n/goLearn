package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var userName string

	// Use string concatenation or separate Println calls for each variable
	fmt.Println(smsSendingLimit, costPerSMS, hasPermission, userName)

	// ez way of doing this
	congrtas := "happy Cloud"
	fmt.Println(congrtas)

	//multi deceleration
	milage, company := 898987, "Tesla"
	fmt.Println(milage, company)
}
