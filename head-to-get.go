package head_to_get

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the DEFAULT plugin configuration - no access to config yet!
func CreateConfig() *Config {
	return &Config{
	}
}

// HeadToGetPlugin header
type HeadToGetPlugin struct {
	name   string
	next   http.Handler
}

// New created a new plugin, with a config that's been set (possibly) by the admin
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config == nil {
		return nil, fmt.Errorf("config can not be nil")
	}

	plugin := &HeadToGetPlugin{
		next:   next,
		name:   name,
	}

	return plugin, nil
}

func (t *HeadToGetPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if r.Method == "HEAD" {
		r.Method = "GET"
	}
	t.next.ServeHTTP(rw, req)
}
