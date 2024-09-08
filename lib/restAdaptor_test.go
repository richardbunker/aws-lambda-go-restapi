package lib

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestRestApiAdaptor(t *testing.T) {
	request := events.APIGatewayV2HTTPRequest{
		Body: `{"name": "Mickey Mouse"}`,
		Headers: map[string]string{
			"authorization": "Bearer token",
		},
		Cookies: []string{},
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/users/1234",
			},
		},
	}

	adaptedRequest := RestApiAdaptor(request)

	if adaptedRequest.Method != "GET" {
		t.Errorf("Expected method to be 'GET', got %v", adaptedRequest.Method)
	}
	if adaptedRequest.Path != "/users/1234" {
		t.Errorf("Expected path to be '/users/1234', got %v", adaptedRequest.Path)
	}
	if adaptedRequest.Body["name"] != "Mickey Mouse" {
		t.Errorf("Expected body to be 'Mickey Mouse', got %v", adaptedRequest.Body["name"])
	}
	if adaptedRequest.Headers["authorization"] != "Bearer token" {
		t.Errorf("Expected authorization header to be 'Bearer token', got %v", adaptedRequest.Headers["authorization"])
	}
	if len(adaptedRequest.Cookies) != 0 {
		t.Errorf("Expected cookies to be empty, got %v", adaptedRequest.Cookies)
	}
	if adaptedRequest.Query != nil {
		t.Errorf("Expected query to be nil, got %v", adaptedRequest.Query)
	}
}
