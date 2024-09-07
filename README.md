# AWS Lambda Go RestAPI

A quick way to create a RestAPI using AWS Lambda and API Gateway (HTTP) written in Go.

## Instructions

To create a new RestAPI along with registering your routes, you can use the following code:

```go
api := RestApi()

// Register the routes
api.Get("/users/:userId", ShowUser)
api.Put("/users/:userId", UpdateUser)
```

The `ShowUser` and `UpdateUser` functions are the handlers for the routes. You can define them as follows:

```go
// A simple handler to show a user
func ShowUser(request RestApiRequest) RestApiResponse {
	userId := request.PathParams["userId"]
	return RestApiResponse{
		Body: map[string]interface{}{
			"id":   userId,
			"name": "Mickey Mouse",
		},
		StatusCode: 200,
	}
}

// A simple handler to update a user
func UpdateUser(request RestApiRequest) RestApiResponse {
	userId := request.PathParams["userId"]
	name := request.Body["name"]
	return RestApiResponse{
		Body: map[string]interface{}{
			"id":   userId,
			"name": name,
		},
		StatusCode: 200,
	}
}
```

## Local development

To run the application locally, you can use the following command:

```bash
$ go run ./dev/main.go
```

This will create an RestAPI, register the routes and print the response to the console.

#### Example

```go
func main() {
	// Sample request
	request := RestApiRequest{
		Method: "POST",
		Path:   "/users/1234",
		Query: map[string]string{
			"from": "2021-01-01",
		},
		Body: map[string]interface{}{
			"name":  "John Doe",
			"email": "example@test.com",
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
```

## Deployment

To build and zip the application, you can use the following command:

```bash
$ ./build.sh
```

This will create a `deployment.zip` file that you can upload to AWS Lambda.
