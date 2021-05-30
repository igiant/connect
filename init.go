package connect

import "encoding/json"

// InitGetHostname - Returns FQDN (fully qualified domain name) of the server (e.g. mail.companyname.com).
// Return
//	hostname - name of the server
func (c *ServerConnection) InitGetHostname() (string, error) {
	data, err := c.CallRaw("Init.getHostname", nil)
	if err != nil {
		return "", err
	}
	hostname := struct {
		Result struct {
			Hostname string `json:"hostname"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hostname)
	return hostname.Result.Hostname, err
}

// InitCheckHostname - Check existence of domain name in the DNS. Existence of DN record with type "A" in appropriate DNS zone.
// Parameters
//	hostname - fully qualified domain name of the server
func (c *ServerConnection) InitCheckHostname(hostname string) error {
	params := struct {
		Hostname string `json:"hostname"`
	}{hostname}
	_, err := c.CallRaw("Init.checkHostname", params)
	return err
}

// InitCheckMxRecord - Check existence of MX record in the DNS for specified domain.
// Parameters
//	domainName - fully qualified domain name
func (c *ServerConnection) InitCheckMxRecord(domainName string) error {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	_, err := c.CallRaw("Init.checkMxRecord", params)
	return err
}

// InitSetHostname - Set Internet hostname of the server. This name is used for server identification in SMTP, POP3 and similar protocols.
// Parameters
//	hostname - new fully qualified domain name of the server
func (c *ServerConnection) InitSetHostname(hostname string) error {
	params := struct {
		Hostname string `json:"hostname"`
	}{hostname}
	_, err := c.CallRaw("Init.setHostname", params)
	return err
}

// InitGetDistributableDomains - Retrieve domains, which can be distributed, from the master server as a standalone server.
// Parameters
//	authentication - Structure with a credential. Credential will be used when connected is false.
// Return
//	domainNames - List of domains which can be distributed (they have a directory service set).
func (c *ServerConnection) InitGetDistributableDomains(authentication ClusterAuthentication) (StringList, error) {
	params := struct {
		Authentication ClusterAuthentication `json:"authentication"`
	}{authentication}
	data, err := c.CallRaw("Init.getDistributableDomains", params)
	if err != nil {
		return nil, err
	}
	domainNames := struct {
		Result struct {
			DomainNames StringList `json:"domainNames"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &domainNames)
	return domainNames.Result.DomainNames, err
}

// InitCreateDistributableDomain - Connect server to cluster as slave and create distributable domain.
// Parameters
//	domainName - domain which can be distributed (they have a directory service set) and exist on master server.
//	authentication - Structure with a credential. Credential will be used when connected is false.
// Return
//	result - if ClusterErrorType is not clSuccess, error argument contains additional error info
func (c *ServerConnection) InitCreateDistributableDomain(domainName string, authentication ClusterAuthentication) (*ClusterError, error) {
	params := struct {
		DomainName     string                `json:"domainName"`
		Authentication ClusterAuthentication `json:"authentication"`
	}{domainName, authentication}
	data, err := c.CallRaw("Init.createDistributableDomain", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result ClusterError `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return &result.Result.Result, err
}

// InitCreatePrimaryDomain - Creates the primary email domain.
// Parameters
//	domainName - fully qualified name of the domain
func (c *ServerConnection) InitCreatePrimaryDomain(domainName string) error {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	_, err := c.CallRaw("Init.createPrimaryDomain", params)
	return err
}

// InitCreateAdministratorAccount - Creates the administrator account. This account will be created in primary domain.
// Parameters
//	loginName - login name for administrator (without domain name)
//	password - administrator password
func (c *ServerConnection) InitCreateAdministratorAccount(loginName string, password string) error {
	params := struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}{loginName, password}
	_, err := c.CallRaw("Init.createAdministratorAccount", params)
	return err
}

// InitGetMessageStorePath - Get current path to message store. Default path is "store" subdirectory in installation directory.
// Return
//	path - full path to message store directory
//	freeSpace - amount of free space in the directory
func (c *ServerConnection) InitGetMessageStorePath() (string, int, error) {
	data, err := c.CallRaw("Init.getMessageStorePath", nil)
	if err != nil {
		return "", 0, err
	}
	path := struct {
		Result struct {
			Path      string `json:"path"`
			FreeSpace int    `json:"freeSpace"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &path)
	return path.Result.Path, path.Result.FreeSpace, err
}

// InitSetMessageStorePath - Set path to message store directory.
// Parameters
//	path - full path to message store directory
func (c *ServerConnection) InitSetMessageStorePath(path string) error {
	params := struct {
		Path string `json:"path"`
	}{path}
	_, err := c.CallRaw("Init.setMessageStorePath", params)
	return err
}

// InitGetDirs - Obtain a list of directories in a particular path.
// Parameters
//	fullPath - directory for listing, if full path is empty logical drives will be listed
// Return
//	dirList - List of directories
func (c *ServerConnection) InitGetDirs(fullPath string) (DirectoryList, error) {
	params := struct {
		FullPath string `json:"fullPath"`
	}{fullPath}
	data, err := c.CallRaw("Init.getDirs", params)
	if err != nil {
		return nil, err
	}
	dirList := struct {
		Result struct {
			DirList DirectoryList `json:"dirList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dirList)
	return dirList.Result.DirList, err
}

// InitCheckMessageStorePath - Check if message store path is correct and can be created in the file system.
// Parameters
//	path - full path to message store directory
// Return
//	result - result of the check
//	freeSpace - amount of free space in the directory
func (c *ServerConnection) InitCheckMessageStorePath(path string) (*DirectoryAccessResult, int, error) {
	params := struct {
		Path string `json:"path"`
	}{path}
	data, err := c.CallRaw("Init.checkMessageStorePath", params)
	if err != nil {
		return nil, 0, err
	}
	result := struct {
		Result struct {
			Result    DirectoryAccessResult `json:"result"`
			FreeSpace int                   `json:"freeSpace"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return &result.Result.Result, result.Result.FreeSpace, err
}

// InitSetClientStatistics - Set client statistics settings.
// Parameters
//	isEnabled - flag if statistics are enabled
func (c *ServerConnection) InitSetClientStatistics(isEnabled bool) error {
	params := struct {
		IsEnabled bool `json:"isEnabled"`
	}{isEnabled}
	_, err := c.CallRaw("Init.setClientStatistics", params)
	return err
}

// InitFinish - Finish initial configuration of Kerio Connect.
func (c *ServerConnection) InitFinish() error {
	_, err := c.CallRaw("Init.finish", nil)
	return err
}

// InitGetNamedConstantList - Server side list of constants.
// Return
//	constants - list of constants
func (c *ServerConnection) InitGetNamedConstantList() (NamedConstantList, error) {
	data, err := c.CallRaw("Init.getNamedConstantList", nil)
	if err != nil {
		return nil, err
	}
	constants := struct {
		Result struct {
			Constants NamedConstantList `json:"constants"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &constants)
	return constants.Result.Constants, err
}

// InitGetBrowserLanguages - Returns a list of user-preferred languages set in browser.
// Return
//	calculatedLanguage - a list of 2-character language codes
func (c *ServerConnection) InitGetBrowserLanguages() (StringList, error) {
	data, err := c.CallRaw("Init.getBrowserLanguages", nil)
	if err != nil {
		return nil, err
	}
	calculatedLanguage := struct {
		Result struct {
			CalculatedLanguage StringList `json:"calculatedLanguage"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &calculatedLanguage)
	return calculatedLanguage.Result.CalculatedLanguage, err
}

// InitGetProductInfo - Get basic information about product and its version.
// Return
//	info - structure with basic information about product
func (c *ServerConnection) InitGetProductInfo() (*ProductInfo, error) {
	data, err := c.CallRaw("Init.getProductInfo", nil)
	if err != nil {
		return nil, err
	}
	info := struct {
		Result struct {
			Info ProductInfo `json:"info"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &info)
	return &info.Result.Info, err
}

// InitGetEula - Obtain EULA.
// Return
//	content - plain text of EULA
func (c *ServerConnection) InitGetEula() (string, error) {
	data, err := c.CallRaw("Init.getEula", nil)
	if err != nil {
		return "", err
	}
	content := struct {
		Result struct {
			Content string `json:"content"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &content)
	return content.Result.Content, err
}
