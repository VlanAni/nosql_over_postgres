package validator

import (
	"errors"
	"testing"
)

func TestValidateCollectionName(t *testing.T) {
	testCases := []struct {
		name           string
		collectionName string
		wantError      bool
	}{
		{"Valid simple name", "users", false},
		{"Valid with numbers", "users_2026", false},
		{"Valid starts with underscore", "_hidden_data", false},
		{"Valid max length (63)", "a12345678901234567890123456789012345678901234567890123456789012", false},

		{"Invalid empty string", "", true},
		{"Invalid starts with number", "1_users", true},
		{"Invalid contains space", "my users", true},
		{"Invalid contains hyphen", "my-users", true},
		{"Invalid SQL injection", "users; DROP TABLE users;", true},
		{"Invalid too long (64)", "a123456789012345678901234567890123456789012345678901234567890123", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := ValidateCollectionName(tc.collectionName)

			hasError := err != nil
			if hasError != tc.wantError {
				t.Errorf("ValidateCollectionName(%q) returned err: %v, wantError: %v", tc.collectionName, err, tc.wantError)
			}

			if tc.wantError && !errors.Is(err, ErrInvalidCollectionName) {
				t.Errorf("Expected error to be ErrInvalidCollectionName, got %v", err)
			}
		})
	}
}
