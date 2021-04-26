package testing

import (
	"net/http"
)

// RoundTripFunc is used to mock http.Transport.
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip implements http.Transport.RoundTrip function.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns the mocked http.Client.
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
