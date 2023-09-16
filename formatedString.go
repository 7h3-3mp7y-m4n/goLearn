package main

import "fmt"

func main() {
	const Name = "LocalHost"
	//fmt.Printf = prnt formates string to standared out
	//fmt.Sprintf() = return formated strings
	// %v default repe  %s for string %d for inetegr decimal %f for float
	//create "hi Name + , your open rate is OpenRate percent"

	const OpenRate = 30.5
	fmt.Printf("Hey %s , your open rate is %.2f percent", Name, OpenRate)
	//or we can do Sprintf

	msg := fmt.Sprintf("\nHey %s , your open rate is %.2f percent", Name, OpenRate)
	fmt.Println(msg)

}
