package app

import (
	"fmt"
	"regexp"
	"strings"
)

type RestApiRequest struct {
	Method     string
	Path       string
	PathParams map[string]string
	Query      map[string]string
	Body       map[string]interface{}
}

type RestApiResponse struct {
	Body       map[string]interface{}
	StatusCode int
}

type Handler func(request RestApiRequest) RestApiResponse

type Routes map[string]Handler

type Method string

type Api struct {
	methods map[Method]Routes
}

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

// Create a new API application
func RestApi() *Api {
	return &Api{
		methods: make(map[Method]Routes),
	}
}

// Register a handler for a GET request
func (api *Api) Get(pathToMatch string, handler Handler) {
	api.RegisterRoute(GET, pathToMatch, handler)
}

// Register a handler for a POST request
func (api *Api) Post(pathToMatch string, handler Handler) {
	api.RegisterRoute(POST, pathToMatch, handler)
}

// Register a handler for a PUT request
func (api *Api) Put(pathToMatch string, handler Handler) {
	api.RegisterRoute(PUT, pathToMatch, handler)
}

// Register a handler for a DELETE request
func (api *Api) Delete(pathToMatch string, handler Handler) {
	api.RegisterRoute(DELETE, pathToMatch, handler)
}

// Register a handler for a request
func (api *Api) RegisterRoute(method Method, pathToMatch string, handler Handler) {
	if api.methods[method] == nil {
		api.methods[method] = make(Routes)
	}
	api.methods[method][pathToMatch] = handler
}

// Handle a request
func (api *Api) HandleRequest(request RestApiRequest) RestApiResponse {
	if api.methods[Method(request.Method)] == nil {
		return Error(405, "Method not allowed")
	}
	var handler Handler
	for route := range api.methods[Method(request.Method)] {
		if match(request.Path, route) {
			handler = api.methods[Method(request.Method)][route]
			request.PathParams = extractPathParams(request.Path, route)
			break
		}
	}
	if handler == nil {
		return Error(404, "Route not found")
	}
	return handler(request)
}

// Extract path parameters from the requested path
func extractPathParams(requestedPath string, registeredPath string) map[string]string {
	params := make(map[string]string)
	requestedPathParts := strings.Split(requestedPath, "/")
	registeredPathParts := strings.Split(registeredPath, "/")
	for i, part := range registeredPathParts {
		if strings.HasPrefix(part, ":") {
			params[strings.TrimPrefix(part, ":")] = requestedPathParts[i]
		}
	}
	return params
}

// Check if the requested path matches the registered route
func match(path string, route string) bool {
	pattern := regexp.MustCompile(`:\w+`).ReplaceAllString(route, "([^/]+)")
	regex, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		fmt.Println("Error in regex")
		return false
	}
	return regex.MatchString(path)
}

// A helper function to create an error response
func Error(statusCode int, message string) RestApiResponse {
	return RestApiResponse{
		Body: map[string]interface{}{
			"error": message,
		},
		StatusCode: statusCode,
	}
}
