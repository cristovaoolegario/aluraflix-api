package dto

import "errors"

func MissingFieldError(missingField string) error {
	return errors.New(missingField + " is required.")
}
