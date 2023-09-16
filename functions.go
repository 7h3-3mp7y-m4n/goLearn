package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

//same funtcion with diffrent apporach

func concatV2(s1, s2 string) string {
	return s1 + " " + s2
}

func main() {
	fmt.Println(concat("hello", "github"))
	fmt.Println(concatV2("Still", "learning"))
	fmt.Println(concat("until", "I make it "))
}
