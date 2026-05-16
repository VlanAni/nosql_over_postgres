package validator

import (
	"errors"
	"regexp"
)

var collectionNameRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]{0,62}$`)

var ErrInvalidCollectionName = errors.New("invalid collection name: must start with a letter or underscore, contain only alphanumeric characters or underscores, and be 1 to 63 characters long")

func ValidateCollectionName(name string) error {
	if !collectionNameRegex.MatchString(name) {
		return ErrInvalidCollectionName
	}
	return nil
}
