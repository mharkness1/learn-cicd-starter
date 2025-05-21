package auth

import (
	"net/http"
	"testing"
)

func TestGetApi(t *testing.T) {
	tests := []struct {
		name          string
		inputHeader   http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name: "basic implementation",
			inputHeader: http.Header{
				"Content-Type":  []string{"text/html"},
				"Authorization": []string{"ApiKey key123"},
			},
			expectedKey:   "key123",
			expectedError: false,
		},
		{
			name: "missing api key",
			inputHeader: http.Header{
				"Authorization": []string{"No Key"},
			},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "missing authorization",
			inputHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.inputHeader)
			if tc.expectedError == true {
				if err == nil {
					t.Errorf("FAIL test %d - %s: expected an error to return, got none", i, tc.name)
				}
				if key != "" {
					t.Errorf("test %d - %s: expected key to return empty but got %s", i, tc.name, key)
					return
				}
			} else if tc.expectedError == false {
				if err != nil {
					t.Errorf("FAIL test %d - %s: expected no error but got: %v", i, tc.name, err)
				}
				if key != tc.expectedKey {
					t.Errorf("test %d - %s: expected key didn't match: expected: %s but got %s", i, tc.name, tc.expectedKey, key)
					return
				}
			}
		})
	}

}
