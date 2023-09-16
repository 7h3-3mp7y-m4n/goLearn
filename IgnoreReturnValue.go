package main

import "fmt"

func main() {
	firstName, _ := getname()
	fmt.Println("Welcome to Textio", firstName)

}

func getname() (string, string) {
	return "Jhon", "Doe"
}
