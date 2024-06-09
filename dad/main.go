package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	app := &cli.App{
		Name:  "DadJoke",
		Usage: "Well its just a Joke of your Dad ;)",
		Action: func(*cli.Context) error {
			getRandomJoke()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	asciiArt := `
╔═══╦═══╦═══╗──╔╦═══╦╗╔═╦═══╗
╚╗╔╗║╔═╗╠╗╔╗║──║║╔═╗║║║╔╣╔══╝
─║║║║║─║║║║║║──║║║─║║╚╝╝║╚══╗
─║║║║╚═╝║║║║║╔╗║║║─║║╔╗║║╔══╝
╔╝╚╝║╔═╗╠╝╚╝║║╚╝║╚═╝║║║╚╣╚══╗
╚═══╩╝─╚╩═══╝╚══╩═══╩╝╚═╩═══╝
`
	fmt.Println(asciiArt)
	fmt.Println(string("\033[31m" + "--" + "\033[0m" + "--" + "\033[32m" + "--" + "\033[33m" + "--" + "\033[35m" + "-> " + "\033[0m" + joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)
	if err != nil {
		log.Println("Well you dont have a dad so no joke for you man BTW here is a error code -  %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "DadJoke cli (github.com/7h3-3mp7y-m4n/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("well you already know you dont have a dad , here is God error while making it - %v", err)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("No joke man live a happy life - %v", err)
	}
	return responseByte
}
