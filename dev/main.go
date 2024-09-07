package main

import (
	. "aws_lambda/app"
	. "aws_lambda/handlers"
	"encoding/json"
	"fmt"
)

func main() {
	// Sample GET request
	getRequest := RestApiRequest{
		Method: "GET",
		Path:   "/users/1234",
		Query: map[string]string{
			"from": "2021-01-01",
		},
		Body: nil,
	}

	// Sample PUT request
	postRequest := RestApiRequest{
		Method: "PUT",
		Path:   "/users/1234",
		Query:  nil,
		Body: map[string]interface{}{
			"name": "Minny Mouse",
		},
	}

	// Create a new API application
	api := RestApi()

	// Register the routes
	api.Get("/users/:userId", ShowUser)
	api.Put("/users/:userId", UpdateUser)

	// Handle the requests
	getResponse := api.HandleRequest(getRequest)
	postResponse := api.HandleRequest(postRequest)

	// Print the response bodies
	getBody, _ := json.Marshal(getResponse.Body)
	postBody, _ := json.Marshal(postResponse.Body)
	fmt.Println("GET Request Body: ", string(getBody))
	fmt.Println("PUT Response Body: ", string(postBody))
}
