package head_to_get

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	PutOriginalMethodInHeader string
}

// CreateConfig creates the DEFAULT plugin configuration - no access to config yet!
func CreateConfig() *Config {
	return &Config{}
}

// HeadToGetPlugin header
type HeadToGetPlugin struct {
	putOriginalMethodInHeader string
	name                      string
	next                      http.Handler
}

// New created a new plugin, with a config that's been set (possibly) by the admin
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config == nil {
		return nil, fmt.Errorf("config can not be nil")
	}

	plugin := &HeadToGetPlugin{
		putOriginalMethodInHeader: config.PutOriginalMethodInHeader,
		next:                      next,
		name:                      name,
	}

	return plugin, nil
}

func (t *HeadToGetPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodHead {
		if t.putOriginalMethodInHeader != "" {
			req.Header.Add(t.putOriginalMethodInHeader, req.Method)
		}
		req.Method = http.MethodGet
		rw = NewBodyDroppingResponseWriter(rw)
	}
	t.next.ServeHTTP(rw, req)
}

func NewBodyDroppingResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	r := &BodyDroppingResponseWriter{ResponseWriter: w}
	return r
}

type BodyDroppingResponseWriter struct {
	http.ResponseWriter
}

func (w *BodyDroppingResponseWriter) CloseNotify() <-chan bool {
	if cn, ok := w.ResponseWriter.(http.CloseNotifier); ok {
		return cn.CloseNotify()
	}
	return nil
}

func (w *BodyDroppingResponseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *BodyDroppingResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *BodyDroppingResponseWriter) Write(bytes []byte) (int, error) {
	return 0, nil // dropping body
}

func GetToRemoveResponseHeaders() []string {
	return []string{
		"Content-Type",
		"Content-Length",
		"Content-Disposition",
		"Content-Digest",
		"Digest",
		"Content-Location",
		"Content-Language",
		"Content-Encoding",
		"Accept-Ranges",
		"Transfer-Encoding",
	}
}

func (w *BodyDroppingResponseWriter) WriteHeader(statusCode int) {
	for _, v := range GetToRemoveResponseHeaders() {
		w.ResponseWriter.Header().Del(v)
	}
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *BodyDroppingResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}
