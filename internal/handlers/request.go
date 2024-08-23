package handlers

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type Request struct {
	Method  string
	URI     string
	Version string
	Headers map[string]string
	Body    string
}

type ContextManager struct {
	Writer          net.Conn
	Req             Request
	Params          map[string]string
	Keys            map[string]interface{}
	Errors          []error
	AcceptedContent []string
	FullPath        string
}

func (c *ContextManager) SendResponse(statusCode int, body string) {
	statusline := c.CreateStatusLine(statusCode)
	headers := c.Req.Headers
	headers["Content-Length"] = fmt.Sprint(len(c.Req.Body))
	/*
		Example:
		HTTP/1.1 200 OK
		Content-Type: text/html; charset=utf-8
		Content-Length: 13

		<html><body><h1>Hello World</h1></body></html>
	*/

	c.Writer.Write([]byte(statusline + "\r\n"))
	for k, v := range headers {
		c.Writer.Write([]byte(k + ": " + v + "\r\n"))
	}
	c.Writer.Write([]byte("\r\n"))
	c.Writer.Write([]byte(body))
}

func (c *ContextManager) CreateStatusLine(statusCode int) string {
	switch statusCode {
	case 200:
		return "HTTP/1.1 200 OK"
	case 404:
		return "HTTP/1.1 404 Not Found"
	case 500:
		return "HTTP/1.1 500 Internal Server Error"
	case 405:
		return "HTTP/1.1 405 Method Not Allowed"
	default:
		return "HTTP/1.1 200 OK"
	}
}

// ? what should the handler function take
// the handler function should have a context which would have all the information about the request
// the handler function should return a response
type HandlerFunc func() Response

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

type Router struct {
	Routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GET(path string, handler HandlerFunc) {
	// create a Route
	route := Route{
		Method:  "GET",
		Path:    path,
		Handler: handler,
	}

	r.Routes = append(r.Routes, route)
}

// router := NewRouter()

// router("/phones", "GET", PhonesHandler)
// on hitting the /phones route call the handlerFunction, if the function does not exist give error

func GetRequestMethod(req string) string {
	return strings.Split(req, " ")[0]
}

func ParseHeaders(lines []string) (map[string]string, error) {
	headers := make(map[string]string)
	// take the first line and split by whitespace

	requestLine := strings.Fields(lines[0])
	// Example: GET / HTTP 1.1 ...
	if len(requestLine) < 3 {
		return nil, errors.New("Invalid Request Line")
	}
	method := requestLine[0]
	path := requestLine[1]
	version := requestLine[2]

	headers["Method"] = method
	headers["Path"] = path
	headers["version"] = version

	for _, line := range lines[1:] {
		// if we hit an empty line, headers are done
		if line == "" {
			break
		}
		// Host: localhost:8080
		// User-Agent: curl/7.64.1
		headersParts := strings.SplitN(line, ":", 2) // handles the case with port
		// remove whitespace
		key := strings.TrimSpace(headersParts[0])
		value := strings.TrimSpace(headersParts[1])
		// key : value
		headers[key] = value
	}

	return headers, nil
}

func ParseBody(lines []string) string {
	// body example: {"name": "John Doe", "email": "john.doe@example.com"}
	// if body is empty
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\r\n")
}
