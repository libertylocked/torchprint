package errors

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPError is an HTTP error from API
type HTTPError struct {
	StatusCode int
	Body       []byte
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP error %v. %v", e.StatusCode, string(e.Body))
}

// NewHTTPError returns an HTTP error with status code
func NewHTTPError(r *http.Response) error {
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return HTTPError{
		StatusCode: r.StatusCode,
		Body:       b,
	}
}
