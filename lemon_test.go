package lemon

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	client *Client
}

func TestSuite(t *testing.T) {
	client := New(os.Getenv("LEMON_KEY"), WithDebug(true))
	if client.key == "" {
		t.SkipNow()
	}

	suite.Run(t, &Suite{client: client})
}

func TestVerifySign(t *testing.T) {
	client := New("123456", WithDebug(true))
	if client.key == "" {
		t.SkipNow()
	}

	expected := "Hello, World!"
	digest := "fe2cf3347d7855c9d6ecb8f6aeea8eeb24fa1bef21aea4c36a05ed423259bb8a"

	req := httptest.NewRequest("POST", "/path/to/resource", bytes.NewReader([]byte(expected)))
	req.Header.Set("X-Signature", digest)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := client.VerifyRequestSign(r); err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expected))
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	suite.Run(t, &Suite{client: client})
}
