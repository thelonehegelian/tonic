package handlers

import (
	"reflect"
	"testing"
)

func TestParseHeaders(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "Valid headers with GET request",
			input: []string{
				"GET / HTTP/1.1",
				"Host: localhost:8080",
				"User-Agent: curl/7.64.1",
				"Accept: */*",
				"",
			},
			expected: map[string]string{
				"Method":     "GET",
				"Path":       "/",
				"Version":    "HTTP/1.1",
				"Host":       "localhost:8080",
				"User-Agent": "curl/7.64.1",
				"Accept":     "*/*",
			},
			wantErr: false,
		},

		{
			name: "Method not allowed",
			input: []string{
				"BET / HTTP/1.1",
				"Host: localhost:8080",
				"User-Agent: curl/7.64.1",
				"Accept: */*",
				"",
			},
			expected: nil,
			wantErr:  true,
		},

		{
			name: "Path is not Valid",
			input: []string{
				"GET path HTTP/1.1",
				"Host: localhost:8080",
				"User-Agent: curl/7.64.1",
				"Accept: */*",
				"",
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "HTTP Version not Allowed",
			input: []string{
				"GET / HTTP/2.2",
				"Host: localhost:8080",
				"User-Agent: curl/7.64.1",
				"Accept: */*",
				"",
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseHeaders(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("ParseHeaders() error = %v, wantErr %v", err, test.wantErr)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ParseHeaders() = %v, want %v", result, test.expected)
			}
		})
	}
}
