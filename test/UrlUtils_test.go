package swock_test

import (
	"testing"

	"github.com/HARDY8118/swock/swock"
)

func TestValidateUrl(t *testing.T) {
	var tests = []struct {
		name     string
		url      string
		expected bool
	}{
		{"[P] HTTP local no port", "http://localhost", true},
		{"[P] HTTP localhost port", "http://localhost:3000", true},
		{"[P] HTTPS local port", "https://localhost:3000", true},
		{"[P] HTTP IP port", "http://0.0.0.0:3000", true},
		{"[P] HTTPS IP port", "https://10.3.111.250:3000", true},
		{"[P] HTTP domain", "http://example.com", true},
		{"[P] HTTPS domain", "https://example.com", true},
		{"[P] HTTP domain port", "http://example.com:5000", true},
		{"[P] HTTPS domain port", "https://example.com:5000", true},
		{"[N] Invalid protocol", "httx://localhost", false},
		{"[N] Missing /", "http:/localhost", false},
		{"[N] Unsupported protocol", "ftp://localhost", false},
		{"[N] Repeated port", "http://localhost:3000:5000", false},
		{"[N] Invalid protocol localhost", "httpd://localhost:5000", false},
		{"[N] Missing address", "http://:5000", false},
		{"[N] Invalid URL", "http::5000", false},
		{"[N] No URL, only protocol", "http::", false},
		{"[N] Missing protocol", "example.com", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if swock.ValidateUrl(test.url) != test.expected {
				t.Errorf("Failed expected %t for %s", test.expected, test.url)
			}
		})
	}
}
