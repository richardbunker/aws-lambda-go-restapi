package app

import (
	"fmt"
	"testing"
)

// Mock handler for testing
func ShowPost(request RestApiRequest) RestApiResponse {
	return RestApiResponse{
		Body: map[string]interface{}{
			"postId": request.PathParams["postId"],
		},
		StatusCode: 200,
	}
}

func TestApiGetRoute(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a GET route
	api.Get("/posts/:postId", RouteOptions{
		Handler:    ShowPost,
		Middleware: []Middleware{},
	})

	// Test case: Authorized request
	request := RestApiRequest{
		Method: "GET",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	// Assert response body
	if response.Body["postId"] != "123" {
		t.Errorf("Expected postId to be '123', got %v", response.Body["postId"])
	}
}

func TestApiPostRoute(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a POST route
	api.Post("/posts", RouteOptions{
		Handler:    ShowPost,
		Middleware: []Middleware{},
	})

	// Test case: Authorized request
	request := RestApiRequest{
		Method: "POST",
		Path:   "/posts",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}

func TestApiPutRoute(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a PUT route
	api.Put("/posts/:postId", RouteOptions{
		Handler:    ShowPost,
		Middleware: []Middleware{},
	})

	// Test case: Authorized request
	request := RestApiRequest{
		Method: "PUT",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}

func TestApiDeleteRoute(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a DELETE route
	api.Delete("/posts/:postId", RouteOptions{
		Handler:    ShowPost,
		Middleware: []Middleware{},
	})

	// Test case: Authorized request
	request := RestApiRequest{
		Method: "DELETE",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}

func TestApiMiddleware(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a GET route
	api.Get("/posts/:postId", RouteOptions{
		Handler: ShowPost,
		Middleware: []Middleware{
			func(request RestApiRequest) (error, *MiddlewareReason) {
				return nil, nil
			},
		},
	})

	// Test case: Authorized request
	request := RestApiRequest{
		Method: "GET",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}
func TestApiMiddlewareShouldHalt(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a GET route
	api.Get("/posts/:postId", RouteOptions{
		Handler: ShowPost,
		Middleware: []Middleware{
			func(request RestApiRequest) (error, *MiddlewareReason) {
				return fmt.Errorf("Unauthorized"), &MiddlewareReason{
					StatusCode: 401,
					Messsage:   "Unauthorized",
				}
			},
		},
	})

	// Test case: Unauthorized request
	request := RestApiRequest{
		Method: "GET",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 401 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}
