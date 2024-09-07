package handlers

import (
	. "aws_lambda/app"
)

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
