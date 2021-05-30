package connect

import "encoding/json"

type SignOn struct {
	IsEnabled bool   `json:"isEnabled"`
	HostName  string `json:"hostName"` // Hostname to the Kerio Unity Sign On server. Non default port can be added Eg: example.com:4444
	UserName  string `json:"userName"` // Administrator username
	Password  string `json:"password"` // [WRITE-ONLY] Administrator password
}

// UnitySignOnGet - Obtain Kerio Unity Sign On settings
// Return
//	settings - Sign On settings
func (c *ServerConnection) UnitySignOnGet() (*SignOn, error) {
	data, err := c.CallRaw("UnitySignOn.get", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings SignOn `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// UnitySignOnSet - Set Kerio Unity Sign On settings
// Parameters
//	settings - Sign On settings
func (c *ServerConnection) UnitySignOnSet(settings SignOn) error {
	params := struct {
		Settings SignOn `json:"settings"`
	}{settings}
	_, err := c.CallRaw("UnitySignOn.set", params)
	return err
}

// UnitySignOnTestConnection - Test connection to Kerio Unity Sign On server
// Parameters
//	hostNames - directory server (primary and secondary if any)
//	credentials - authentication information
// Return
//	errors - error messages list; If no error is listed, connection is successful
func (c *ServerConnection) UnitySignOnTestConnection(hostNames StringList, credentials Credentials) (ErrorList, error) {
	params := struct {
		HostNames   StringList  `json:"hostNames"`
		Credentials Credentials `json:"credentials"`
	}{hostNames, credentials}
	data, err := c.CallRaw("UnitySignOn.testConnection", params)
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
