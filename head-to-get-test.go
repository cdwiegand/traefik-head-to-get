package head_to_get

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		config     *Config
		assertFunc func(t *testing.T) http.Handler
	}{
		{
			name:   "default config",
			config: &Config{},
			assertFunc: func(t *testing.T) http.Handler {
				t.Helper()
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					if req.Method == "HEAD" {
						t.Fatalf("HEAD must become GET!")
					}
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			handler, err := New(ctx, tt.assertFunc(t), tt.config, "head-to-get-test")
			if err != nil {
				t.Fatalf("error creating new plugin instance: %+v", err)
			}
			recorder := httptest.NewRecorder()
			req, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost/", nil)
			if err != nil {
				t.Fatalf("error with request: %+v", err)
			}

			handler.ServeHTTP(recorder, req)
		})
	}
}
