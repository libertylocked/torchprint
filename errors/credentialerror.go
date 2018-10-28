package errors

import "fmt"

// NoCredentialOrTokenError no credential or token in api
type NoCredentialOrTokenError struct{}

func (e NoCredentialOrTokenError) Error() string {
	return fmt.Sprintf("Missing credential or token in API")
}

// NoCredentialError no credential in api
type NoCredentialError struct{}

func (e NoCredentialError) Error() string {
	return fmt.Sprintf("Missing credential in API")
}
