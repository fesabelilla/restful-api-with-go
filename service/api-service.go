package service

import (
	"encoding/json"
	"fmt"
	"restful-api-with-go/dtos"
)

func Unmarshalling(responseData []byte) {

	var responseObject dtos.Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for _, Pokemon := range responseObject.Pokemon {
		fmt.Println(Pokemon.Species.Name)
	}
}
