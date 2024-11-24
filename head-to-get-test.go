package head_to_get

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_Head(t *testing.T) {
	t.Run("test using HEAD", func(t *testing.T) {
		ctx := context.Background()
		next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

		cfg := CreateConfig()
		handler, err := New(ctx, next, cfg, "head-to-get")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost/", nil)
		if err != nil {
			t.Fatalf("error with request: %+v", err)
		}

		handler.ServeHTTP(recorder, req)

		t.Helper()
		if req.Method != http.MethodGet {
			t.Fatalf("Method should be GET, was " + req.Method)
		}
		if req.Header.Get("Content-Type") != "" {
			t.Fatalf("Content-type header should not be present in response!")
		}
	})
}

func TestServeHTTP_Get(t *testing.T) {
	t.Run("test using HEAD", func(t *testing.T) {
		ctx := context.Background()
		next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

		cfg := CreateConfig()
		handler, err := New(ctx, next, cfg, "head-to-get")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/", nil)
		if err != nil {
			t.Fatalf("error with request: %+v", err)
		}

		handler.ServeHTTP(recorder, req)

		t.Helper()
		if req.Method != http.MethodGet {
			t.Fatalf("Method should be GET, was " + req.Method)
		}
		if req.Header.Get("Content-Type") != "" {
			t.Fatalf("Content-type header should not be present in response!")
		}
	})
}

func TestServeHTTP_Post(t *testing.T) {
	t.Run("test using HEAD", func(t *testing.T) {
		ctx := context.Background()
		next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

		cfg := CreateConfig()
		handler, err := New(ctx, next, cfg, "head-to-get")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/", nil)
		if err != nil {
			t.Fatalf("error with request: %+v", err)
		}

		handler.ServeHTTP(recorder, req)

		t.Helper()
		if req.Method != http.MethodPost {
			t.Fatalf("Method should be POST, was " + req.Method)
		}
		if req.Header.Get("Content-Type") != "" {
			t.Fatalf("Content-type header should not be present in response!")
		}
	})
}
