package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		// name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		// Test cases will go here
		"Valid API Key": {
			// name:          "Valid API Key",
			headers:       http.Header{"Authorization": []string{"ApiKey abc123xyz789"}},
			expectedKey:   "abc123xyz789",
			expectedError: nil,
		},
		"Missing Authorization Header": {
			// name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		"Malformed Authorization Header": {
			// name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"InvalidFormat"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			// assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, err, "hhh")

			assert.Equal(t, key, tc.expectedKey)

		})

	}

}
