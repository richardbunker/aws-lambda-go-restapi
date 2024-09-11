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

// Mock handler for testing that returns as 200
func Return200(request RestApiRequest) RestApiResponse {
	return RestApiResponse{
		Body: map[string]interface{}{
			"message": "ok",
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

func TestApiGroupRoute(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register the group of routes
	api.Group("/group", []Middleware{}, func() {
		api.Get("/one", RouteOptions{
			Handler: Return200,
		})
		api.Get("/two", RouteOptions{
			Handler: Return200,
		})
	})

	// Test case: Group one
	reqOne := RestApiRequest{
		Method: "GET",
		Path:   "/group/one",
	}
	// Test case: Group two
	reqTwo := RestApiRequest{
		Method: "GET",
		Path:   "/group/two",
	}
	// Test case: Not in group
	reqThree := RestApiRequest{
		Method: "GET",
		Path:   "/group/three",
	}
	reqFour := RestApiRequest{
		Method: "GET",
		Path:   "/not-here",
	}

	resOne := api.HandleRequest(reqOne)
	resTwo := api.HandleRequest(reqTwo)
	resThree := api.HandleRequest(reqThree)
	resFour := api.HandleRequest(reqFour)

	// Assert status code
	if resOne.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resOne.StatusCode)
	}
	if resTwo.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resTwo.StatusCode)
	}
	if resThree.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", resThree.StatusCode)
	}
	if resFour.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", resFour.StatusCode)
	}
}

func TestApiHandleRequest(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	api.Get("/test", RouteOptions{
		Handler: Return200,
	})

	// Test Requests
	reqOne := RestApiRequest{
		Method: "GET",
		Path:   "/test",
	}
	reqTwo := RestApiRequest{
		Method: "PUT",
		Path:   "/test",
	}

	resOne := api.HandleRequest(reqOne)
	resTwo := api.HandleRequest(reqTwo)

	// Assert status code
	if resOne.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resOne.StatusCode)
	}
	if resTwo.StatusCode != 405 {
		t.Errorf("Expected status code 405, got %d", resTwo.StatusCode)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	api.Get("/test", RouteOptions{
		Handler: Return200,
	})

	// Test Requests
	req := RestApiRequest{
		Method: "PUT",
		Path:   "/test",
	}

	res := api.HandleRequest(req)

	// Assert status code
	if res.StatusCode != 405 {
		t.Errorf("Expected status code 405, got %d", res.StatusCode)
	}
}

func TestRouteMiddleware(t *testing.T) {
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

func TestRouteMiddlewareShouldInterceptRequest(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a GET route
	api.Get("/posts/:postId", RouteOptions{
		Handler: ShowPost,
		Middleware: []Middleware{
			func(request RestApiRequest) (error, *MiddlewareReason) {
				return fmt.Errorf("Unauthorized"), &MiddlewareReason{
					StatusCode: 401,
					Message:    "Unauthorized",
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

func TestHandlesInvalidRouteRegistration(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	// Register a GET route
	api.Get("/posts/:postId[", RouteOptions{
		Handler: ShowPost,
	})

	// Test case: Unauthorized request
	request := RestApiRequest{
		Method: "GET",
		Path:   "/posts/123",
	}

	response := api.HandleRequest(request)

	// Assert status code
	if response.StatusCode != 405 {
		t.Errorf("Expected status code 405, got %d", response.StatusCode)
	}
}

func TestRegistersMiddleware(t *testing.T) {
	// Create a new API instance
	api := RestApi()

	api.Middleware([]Middleware{
		func(request RestApiRequest) (error, *MiddlewareReason) {
			return nil, nil
		},
	})

	if api.globalMiddleware == nil {
		t.Error("Expected globalMiddleware to be not nil")
	}
}

func TestErrorResponse(t *testing.T) {
	response := ApiErrorResponse(401, "Unauthorized")

	// Assert status code
	if response.StatusCode != 401 {
		t.Errorf("Expected status code 401, got %d", response.StatusCode)
	}
}
