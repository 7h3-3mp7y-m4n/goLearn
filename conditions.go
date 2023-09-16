package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("trying to send message ", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("meesage not sent")
	}
}
