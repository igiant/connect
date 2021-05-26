package connect

import "net/url"

type Destination uint8

const (
	Admin Destination = iota
	Client
)

// NewConfig returns a pointer to structure with the configuration for connecting to the API server
// Parameters:
//      dest   - control destination (Admin - for server control, Client - for mail client control)
//      server - address without schema and port
func NewConfig(dest Destination, server string) *Config {
	port := "4040"
	path := "/admin/api/jsonrpc"
	if dest == Client {
		port = "443"
		path = "webmail/api/jsonrpc"
	}
	u := url.URL{
		Scheme: "https",
		Host:   server + ":" + port,
		Path:   path,
	}
	return &Config{
		url: u.String(),
	}
}

// NewApplication returns a pointer to structure with application data
func NewApplication(name, vendor, version string) *Application {
	if name == "" {
		name = "TempApp"
	}
	if vendor == "" {
		vendor = "TempVendor"
	}
	if version == "" {
		version = "v1.0.1"
	}
	return &Application{
		Name:    name,
		Vendor:  vendor,
		Version: version,
	}
}

func (c *Config) getID() int {
	c.id++
	return c.id
}
