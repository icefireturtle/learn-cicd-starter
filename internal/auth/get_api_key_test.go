package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := make(http.Header)
	header["Authorization"] = []string{"ApiKey heresakey"}

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	want := "heresakey"
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyMalformedError(t *testing.T) {
	header := make(http.Header)
	header["Authorization"] = []string{"heresakey ApiKey"}

	got, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("error: %v", err)
	}

	want := "error: malformed authorization header"
	if want == got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
