package main

import "fmt"

//the thing is that we made a copy and use it so when we want to modfity it then we have to return it and pass it again

func main() {
	sendSoFar := 430
	const sendsToAdd = 25
	sendSoFar = incrementSends(sendSoFar, sendsToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")

}

func incrementSends(sendSoFar, sendsToAdd int) int {
	sendSoFar = sendSoFar + sendsToAdd
	return sendSoFar
}
