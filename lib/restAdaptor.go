package lib

import (
	. "aws_lambda/app"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func RestApiAdaptor(request events.APIGatewayV2HTTPRequest) RestApiRequest {
	var data map[string]interface{}
	json.Unmarshal([]byte(request.Body), &data)

	adaptedRequest := RestApiRequest{
		Headers: request.Headers,
		Cookies: request.Cookies,
		Method:  request.RequestContext.HTTP.Method,
		Path:    request.RequestContext.HTTP.Path,
		Query:   request.QueryStringParameters,
		Body:    data,
	}

	return adaptedRequest
}
