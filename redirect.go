package main

import (
	"fmt"
	"net/http"
	"strconv"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

const (
	// Default value for HTTP status code
	DefaultStatusCode int = http.StatusFound
	// Key for loading desired destination
	destinationKey string = "destination"
	// Key for loading desired HTTP status code
	statusCodeKey string = "statuscode"
)

func init() {
	r := NewSpinRedirect()
	spinhttp.Handle(r.handleFunc)
}

func main() {
}

// SpinRedirect is a struct that provides a handleFunc
// for redirecting to a destination URL using configurable HTTP status code.
type SpinRedirect struct {
	cfg ConfigReader
}

// NewSpinRedirect returns a new SpinRedirect
func NewSpinRedirect() SpinRedirect {
	return SpinRedirect{
		cfg: NewDefaultConfigReader(),
	}
}

func (s SpinRedirect) handleFunc(w http.ResponseWriter, r *http.Request) {
	dest := s.getDestination()
	code := s.getStatusCode(r.Method)

	w.Header().Set("Location", dest)
	w.WriteHeader(code)
}

// getDestination returns the destination URL
// If no destination is found, an empty string is returned.
func (s SpinRedirect) getDestination() string {
	return s.cfg.Get(destinationKey)
}

// getStatusCode returns the HTTP status code
// If no status code is found, or if the provided value is invalid,
// DefaultStatusCode is returned.
func (s SpinRedirect) getStatusCode(method string) int {
	str := s.cfg.Get(statusCodeKey)
	code, err := strconv.Atoi(str)
	if err != nil {
		return DefaultStatusCode
	}
	if !isValidRedirectStatusCode(code, method) {
		fmt.Printf("Invalid status code provided: %d. Will use %d instead.\n", code, DefaultStatusCode)
		return DefaultStatusCode
	}

	return code
}

// isValidRedirectStatusCode returns true if the provided status code is valid for redirection
func isValidRedirectStatusCode(code int, method string) bool {
	if code == http.StatusSeeOther &&
		(method == http.MethodPut || method == http.MethodPost) {
		return true
	}
	validCodes := []int{
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
	}

	for _, c := range validCodes {
		if c == code {
			return true
		}
	}
	return false
}
