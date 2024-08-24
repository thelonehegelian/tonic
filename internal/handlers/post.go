package handlers

import "strings"

func ParsePostRequest(req string) (*Request, error) {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)

	headers, err := ParseHeaders(lines[1:])
	if err != nil {
		return nil, err
	}
	body := ParseBody(lines[len(headers)+2:])

	parsedRequest := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    body,
	}

	return parsedRequest, nil
}
