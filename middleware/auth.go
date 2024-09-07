package middleware

import (
	. "aws_lambda/app"
	"fmt"
	"os"
)

func getAuthToken() string {
	return "Bearer " + os.Getenv("AUTH_TOKEN")
}

func Authorize(request RestApiRequest) (error, *MiddlewareReason) {
	if request.Headers["authorization"] == getAuthToken() {
		return nil, nil
	}
	return fmt.Errorf("Unauthorized"), &MiddlewareReason{
		StatusCode: 401,
		Messsage:   "Unauthorized",
	}
}
