package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		authHeader  string
		expectedKey string
		expectErr   bool
	}{
		{
			name:        "valid api key",
			authHeader:  "ApiKey abc123",
			expectedKey: "abc123",
			expectErr:   false,
		},
		{
			name:       "missing authorization header",
			authHeader: "",
			expectErr:  true,
		},
		{
			name:       "malformed header wrong prefix",
			authHeader: "Bearer abc123",
			expectErr:  true,
		},
		{
			name:       "malformed header missing key",
			authHeader: "ApiKey",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			key, err := GetAPIKey(headers)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Fatalf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
