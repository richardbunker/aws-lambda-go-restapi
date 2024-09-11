package main

import (
	. "aws_lambda/app"
	. "aws_lambda/handlers"
	"encoding/json"
	"fmt"
)

func main() {
	// Sample GET request
	request := RestApiRequest{
		Method: "GET",
		Path:   "/users/1234",
		Query: map[string]string{
			"from": "2021-01-01",
		},
		Body: nil,
	}

	// Create a new API application
	api := RestApi()
	api.Middleware([]Middleware{})

	// Register the routes
	api.Get("/users/:userId", RouteOptions{
		Handler: ShowUser,
	})

	// Handle the requests
	response := api.HandleRequest(request)
	// Print the response bodies
	body, _ := json.Marshal(response.Body)
	fmt.Println(string(body))
}
