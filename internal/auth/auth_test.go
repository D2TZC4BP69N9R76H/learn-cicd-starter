package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		name          string
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

			if key != tc.expectedKey {
				t.Errorf("expected key %q, got %q with the following error: %v", tc.expectedKey, key, err)
			}

		})

		// if !reflect.DeepEqual(tc.expectedKey, key) {
		// 	t.Fatalf("expected: %v, got: %v with the following error: %v", tc.expectedKey, key, err)
		// }
	}
}
