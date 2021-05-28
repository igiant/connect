package connect

import (
	"net/url"
	"strings"
)

const (
	port = ":4040"
	path = "/admin/api/jsonrpc"
)

// NewConfig returns a pointer to structure with the configuration for connecting to the API server
// Parameters:
//      dest   - control destination (Admin - for server control, Client - for mail client control)
//      server - address without schema and port
func NewConfig(server string) *Config {
	if !strings.Contains(server, ":") {
		server += port
	}
	u := url.URL{
		Scheme: "https",
		Host:   server,
		Path:   path,
	}
	return &Config{
		url: u.String(),
	}
}

// NewApplication returns a pointer to structure with application data
func NewApplication(name, vendor, version string) *ApiApplication {
	if name == "" {
		name = "TempApp"
	}
	if vendor == "" {
		vendor = "TempVendor"
	}
	if version == "" {
		version = "v1.0.1"
	}
	return &ApiApplication{
		Name:    name,
		Vendor:  vendor,
		Version: version,
	}
}

func (c *Config) getID() int {
	c.id++
	return c.id
}
