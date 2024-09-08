package middleware

import (
	. "aws_lambda/app"
	"testing"
)

func TestAuthorizePasses(t *testing.T) {
	request := RestApiRequest{
		Headers: map[string]string{
			"authorization": "Bearer ",
		},
	}

	err, reason := Authorize(request)

	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
	if reason != nil {
		t.Errorf("Expected reason to be nil, got %v", reason)
	}
}
func TestAuthorizeRejects(t *testing.T) {
	request := RestApiRequest{
		Headers: map[string]string{
			"authorization": "Bearer wrong",
		},
	}

	err, reason := Authorize(request)

	if err == nil {
		t.Errorf("Expected error to be 'Unauthorized', got %v", err)
	}
	if reason == nil {
		t.Errorf("Expected reason to be nil, got %v", reason)
	}
	if reason.StatusCode != 401 {
		t.Errorf("Expected status code to be 401, got %v", reason.StatusCode)
	}
	if reason.Message != "Unauthorized" {
		t.Errorf("Expected message to be 'Unauthorized', got %v", reason.Message)
	}
}
