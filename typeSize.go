package main

import "fmt"

func main() {
	accountAge := 2.6

	//create a new "accountAgeInt" here
	//it should be the result of type casting
	accountAgeInt := int(accountAge)

	fmt.Println("Your account existed for around", accountAgeInt, "years")
}
