package handlers

import "strings"

type Request struct {
	Method  string
	URI     string
	Version string
	Headers map[string]string
	Body    string
}

func GetRequestMethod(req string) string {
	return strings.Split(req, " ")[0]
}

func ParseHeaders(lines []string) map[string]string {

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		// if we hit an empty line, headers are done
		if line == "" {
			break
		}
		// Host: localhost:8080
		// User-Agent: curl/7.64.1
		headersParts := strings.Split(line, ":")
		// key : value
		headers[headersParts[0]] = headersParts[1]
	}

	return headers
}

func ParseBody(lines []string) string {
	// body example: {"name": "John Doe", "email": "john.doe@example.com"}
	// if body is empty
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\r\n")
}
