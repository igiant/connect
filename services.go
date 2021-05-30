package connect

import "encoding/json"

// StartUp - type of start-up
type StartUp string

const (
	Manual    StartUp = "Manual"    // service should be started manually
	Automatic StartUp = "Automatic" // service starts automatically
)

type AddressType string

const (
	AllAddresses  AddressType = "AllAddresses"  // all IP addresses for this machine
	Localhost     AddressType = "Localhost"     // localhost - special case
	RealIpAddress AddressType = "RealIpAddress" // a specific address of this machine
)

// Listener - Listening entity
// Note: all fields must be assigned if used in set methods
type Listener struct {
	Type    AddressType `json:"type"`
	Address IpAddress   `json:"address"` // can obtain localizable string "All addresses"
	Port    int         `json:"port"`
}

type ListenerList []Listener

// ConcurrentConnections - Note: all fields must be assigned if used in set methods
type ConcurrentConnections struct {
	IsSet bool `json:"isSet"` // is set maximum of concurrent connections?
	Value int  `json:"value"` // maximum of concurrent connections
}

type GroupLimit struct {
	IsUsed  bool           `json:"isUsed"`  // is group limit set
	IpGroup IpAddressGroup `json:"ipGroup"` // IP address group
}

type Service struct {
	Id              KId                   `json:"id"`
	Name            string                `json:"name"`
	HowToStart      StartUp               `json:"howToStart"`
	Listeners       ListenerList          `json:"listeners"`
	Group           GroupLimit            `json:"group"`
	ConnectionLimit ConcurrentConnections `json:"connectionLimit"`
	DefaultPort     int                   `json:"defaultPort"`
	IsRunning       bool                  `json:"isRunning"`
	AnonymousAccess bool                  `json:"anonymousAccess"` //this property has meaning only for nntp
}

type ServiceList []Service

// ServicesGet - Show a list of services with current status.
// Return
//	services - list of KMS services
func (c *ServerConnection) ServicesGet() (ServiceList, error) {
	data, err := c.CallRaw("Services.get", nil)
	if err != nil {
		return nil, err
	}
	services := struct {
		Result struct {
			Services ServiceList `json:"services"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &services)
	return services.Result.Services, err
}

// ServicesRestart - Restart a given service.
// Parameters
//	service - unique service identifier
func (c *ServerConnection) ServicesRestart(service KId) error {
	params := struct {
		Service KId `json:"service"`
	}{service}
	_, err := c.CallRaw("Services.restart", params)
	return err
}

// ServicesSet - Change current status of service(s).
// Parameters
//	services - list of KMS services
// Return
//	errors - errors of requested changes
func (c *ServerConnection) ServicesSet(services ServiceList) (ErrorList, error) {
	params := struct {
		Services ServiceList `json:"services"`
	}{services}
	data, err := c.CallRaw("Services.set", params)
	if err != nil {
		return nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, err
}

// ServicesStart - Start a given service.
// Parameters
//	service - unique service identifier
func (c *ServerConnection) ServicesStart(service KId) error {
	params := struct {
		Service KId `json:"service"`
	}{service}
	_, err := c.CallRaw("Services.start", params)
	return err
}

// ServicesStop - Stop a given service.
// Parameters
//	service - unique service identifier
func (c *ServerConnection) ServicesStop(service KId) error {
	params := struct {
		Service KId `json:"service"`
	}{service}
	_, err := c.CallRaw("Services.stop", params)
	return err
}

// ServicesStopMacOSServices - Stop the Mac OS services.
func (c *ServerConnection) ServicesStopMacOSServices() error {
	_, err := c.CallRaw("Services.stopMacOSServices", nil)
	return err
}

// ServicesGetIPv6 - Obtain IPv6 settings.
func (c *ServerConnection) ServicesGetIPv6() (bool, error) {
	data, err := c.CallRaw("Services.getIPv6", nil)
	if err != nil {
		return false, err
	}
	isEnabled := struct {
		Result struct {
			IsEnabled bool `json:"isEnabled"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &isEnabled)
	return isEnabled.Result.IsEnabled, err
}

// ServicesSetIPv6 - Set IPv6 settings.
func (c *ServerConnection) ServicesSetIPv6(isEnabled bool) error {
	params := struct {
		IsEnabled bool `json:"isEnabled"`
	}{isEnabled}
	_, err := c.CallRaw("Services.setIPv6", params)
	return err
}
