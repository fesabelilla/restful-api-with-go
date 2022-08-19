package restful_api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"restful-api-with-go/service"
)

func GetResponse() {

	responses, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		log.Fatal("Error occurred during Get API response")
		os.Exit(1)
	}

	//ioutil.ReadAll(response.Body) to read in data from the incoming byte stream
	responseData, err := ioutil.ReadAll(responses.Body)

	if err != nil {
		log.Fatal(err)
	}
	// convert this []byte response into a string using string(responseData)

	responsesValue := string(responseData)
	fmt.Println(responsesValue)

	service.Unmarshalling(responseData)
}
