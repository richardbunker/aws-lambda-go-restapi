package main

import (
	. "aws_lambda/app"
	. "aws_lambda/handlers"
	// . "aws_lambda/middleware"
	"encoding/json"
	"fmt"
)

func main() {
	// Sample GET request
	getPostsRequest := RestApiRequest{
		Method: "GET",
		Path:   "/posts/1234",
		Query: map[string]string{
			"from": "2021-01-01",
		},
		Body: nil,
	}
	// Sample GET request
	getRequest := RestApiRequest{
		// Headers: map[string]string{
		// 	"authorization": "Bearer 987654321",
		// },
		Cookies: []string{},
		Method:  "GET",
		Path:    "/users/1234",
		Query: map[string]string{
			"from": "2021-01-01",
		},
		Body: nil,
	}

	// Sample PUT request
	postRequest := RestApiRequest{
		Headers: map[string]string{},
		Cookies: []string{},
		Method:  "PUT",
		Path:    "/users/1234",
		Query:   nil,
		Body: map[string]interface{}{
			"name": "Minny Mouse",
		},
	}

	// Create a new API application
	api := RestApi()
	api.Middleware([]Middleware{})

	// Register the routes
	api.Group("/users", []Middleware{}, func() {
		api.Get("/:userId", RouteOptions{
			Handler: ShowUser,
		})
		api.Put("/:userId", RouteOptions{
			Handler: UpdateUser,
		})
	})

	api.Get("/posts/:postId", RouteOptions{
		Handler: ShowUser,
	})

	// Handle the requests
	getResponse := api.HandleRequest(getRequest)
	getPostsResponse := api.HandleRequest(getPostsRequest)
	postResponse := api.HandleRequest(postRequest)

	// Print the response bodies
	getBody, _ := json.Marshal(getResponse.Body)
	postBody, _ := json.Marshal(postResponse.Body)
	getPostsBody, _ := json.Marshal(getPostsResponse.Body)
	fmt.Println("GET Request Body: ", string(getBody))
	fmt.Println("PUT Response Body: ", string(postBody))
	fmt.Println("GET Posts Response Body: ", string(getPostsBody))
}
