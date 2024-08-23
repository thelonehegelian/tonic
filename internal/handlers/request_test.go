package handlers

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestParseHeaders(t *testing.T) {
	input := []string{
		"GET / HTTP/1.1",
		"Host: localhost:8080",
		"User-Agent: curl/7.64.1",
		"Accept: */*",
		"",
	}

	expected := map[string]string{
		"Method":     "GET",
		"Path":       "/",
		"Version":    "HTTP/1.1",
		"Host":       "localhost:8080",
		"User-Agent": "curl/7.64.1",
		"Accept":     "*/*",
	}

	result := ParseHeaders(input)

	log.Println(fmt.Sprintf("result: %v", result))
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseHeaders(%v) = %v; want %v", input, result, expected)
	}
}
