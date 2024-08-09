package server

import (
	"fmt"
	"net"
	"tonic/internal/handlers"
)

// findHandler finds a handler for a given path in the router.
//
// Parameters:
// - path: The path to search for a handler.
// - router: The router to search in.
//
// Returns:
// - *handlers.Route: The handler for the given path if found, nil otherwise.
func findHandler(path string, router *handlers.Router) *handlers.Route {
	if router == nil {
		return nil
	}

	for _, route := range router.Routes {
		if route.Path == path {
			return &route
		}
	}

	return nil
}

// Entry point of the app
func HandleRequest(conn net.Conn, router *handlers.Router) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	req := string(buffer)
	if req == "" {
		fmt.Println("Empty Request")
		return
	}

	reqObject := handlers.ParseRequest(req)
	handler := findHandler(reqObject.URI, router)
	handler.Handler()

	method := handlers.GetRequestMethod(req)

	fmt.Println("Method:", method)

	var response *Response
	switch method {
	case "GET":
		request := handlers.ParseGetRequest(req)
		response = CreateResponse(200, map[string]string{"Content-Type": "text/html; charset=utf-8"}, request.Body)
	case "POST":
		request := handlers.ParsePostRequest(req)
		response = CreateResponse(200, map[string]string{"Content-Type": "application/json"}, request.Body)
	default:
		response = CreateResponse(405, map[string]string{"Content-Type": "text/plain"}, "Method Not Allowed")
	}

	SendResponse(conn, response)
	fmt.Println("Response:", response)
}
