package testing

import (
	"net/http"
	"testing"
)

func TestTesting(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{}
	})

	client.Transport.RoundTrip(&http.Request{})
}
