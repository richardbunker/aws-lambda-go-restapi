package main

import (
	. "aws_lambda/app"
	. "aws_lambda/handlers"
	. "aws_lambda/lib"
	. "aws_lambda/middleware"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	req := RestApiAdaptor(request)

	// Create a new API application
	api := RestApi()

	// Register the routes
	api.Get("/users/:userId", RouteOptions{
		Middleware: []Middleware{
			Authorize,
		},
		Handler: ShowUser,
	})

	response := api.HandleRequest(req)

	body, _ := json.Marshal(response.Body)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: response.StatusCode,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
