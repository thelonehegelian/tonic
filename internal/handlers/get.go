package handlers

import (
	"strings"
)

const sampleBody = `<html><body><h1>Hello World</h1></body></html>`

func ParseGetRequest(req string) (*Request, error) {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)
	headers, err := ParseHeaders(lines[1:])
	if err != nil {
		return nil, err
	}

	request := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    sampleBody,
	}
	return request, nil
}

func ParseRequest(req string) (*Request, error) {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)
	headers, err := ParseHeaders(lines[1:])
	if err != nil {
		return nil, err
	}
	request := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    sampleBody,
	}
	return request, nil
}
