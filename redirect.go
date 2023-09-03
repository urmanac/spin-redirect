package main

import (
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
	code := s.getStatusCode()

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
func (s SpinRedirect) getStatusCode() int {
	str := s.cfg.Get(statusCodeKey)
	code, err := strconv.Atoi(str)
	if err != nil {
		return DefaultStatusCode
	}
	return code
}
