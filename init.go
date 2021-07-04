package connect

import "encoding/json"

// InitGetHostname - Returns FQDN (fully qualified domain name) of the server (e.g. mail.companyname.com).
// Return
//	hostname - name of the server
func (s *ServerConnection) InitGetHostname() (string, error) {
	data, err := s.CallRaw("Init.getHostname", nil)
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
//	hostname - fully qualified domain name of the server
func (s *ServerConnection) InitCheckHostname(hostname string) error {
	params := struct {
		Hostname string `json:"hostname"`
	}{hostname}
	_, err := s.CallRaw("Init.checkHostname", params)
	return err
}

// InitCheckMxRecord - Check existence of MX record in the DNS for specified domain.
//	domainName - fully qualified domain name
func (s *ServerConnection) InitCheckMxRecord(domainName string) error {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	_, err := s.CallRaw("Init.checkMxRecord", params)
	return err
}

// InitSetHostname - Set Internet hostname of the server. This name is used for server identification in SMTP, POP3 and similar protocols.
//	hostname - new fully qualified domain name of the server
func (s *ServerConnection) InitSetHostname(hostname string) error {
	params := struct {
		Hostname string `json:"hostname"`
	}{hostname}
	_, err := s.CallRaw("Init.setHostname", params)
	return err
}

// InitGetDistributableDomains - Retrieve domains, which can be distributed, from the master server as a standalone server.
//	authentication - Structure with a credential. Credential will be used when connected is false.
// Return
//	domainNames - List of domains which can be distributed (they have a directory service set).
func (s *ServerConnection) InitGetDistributableDomains(authentication ClusterAuthentication) (StringList, error) {
	params := struct {
		Authentication ClusterAuthentication `json:"authentication"`
	}{authentication}
	data, err := s.CallRaw("Init.getDistributableDomains", params)
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
//	domainName - domain which can be distributed (they have a directory service set) and exist on master server.
//	authentication - Structure with a credential. Credential will be used when connected is false.
// Return
//	result - if ClusterErrorType is not clSuccess, error argument contains additional error info
func (s *ServerConnection) InitCreateDistributableDomain(domainName string, authentication ClusterAuthentication) (*ClusterError, error) {
	params := struct {
		DomainName     string                `json:"domainName"`
		Authentication ClusterAuthentication `json:"authentication"`
	}{domainName, authentication}
	data, err := s.CallRaw("Init.createDistributableDomain", params)
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
//	domainName - fully qualified name of the domain
func (s *ServerConnection) InitCreatePrimaryDomain(domainName string) error {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	_, err := s.CallRaw("Init.createPrimaryDomain", params)
	return err
}

// InitCreateAdministratorAccount - Creates the administrator account. This account will be created in primary domain.
//	loginName - login name for administrator (without domain name)
//	password - administrator password
func (s *ServerConnection) InitCreateAdministratorAccount(loginName string, password string) error {
	params := struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}{loginName, password}
	_, err := s.CallRaw("Init.createAdministratorAccount", params)
	return err
}

// InitGetMessageStorePath - Get current path to message store. Default path is "store" subdirectory in installation directory.
// Return
//	path - full path to message store directory
//	freeSpace - amount of free space in the directory
func (s *ServerConnection) InitGetMessageStorePath() (string, int, error) {
	data, err := s.CallRaw("Init.getMessageStorePath", nil)
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
//	path - full path to message store directory
func (s *ServerConnection) InitSetMessageStorePath(path string) error {
	params := struct {
		Path string `json:"path"`
	}{path}
	_, err := s.CallRaw("Init.setMessageStorePath", params)
	return err
}

// InitGetDirs - Obtain a list of directories in a particular path.
//	fullPath - directory for listing, if full path is empty logical drives will be listed
// Return
//	dirList - List of directories
func (s *ServerConnection) InitGetDirs(fullPath string) (DirectoryList, error) {
	params := struct {
		FullPath string `json:"fullPath"`
	}{fullPath}
	data, err := s.CallRaw("Init.getDirs", params)
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
//	path - full path to message store directory
// Return
//	result - result of the check
//	freeSpace - amount of free space in the directory
func (s *ServerConnection) InitCheckMessageStorePath(path string) (*DirectoryAccessResult, int, error) {
	params := struct {
		Path string `json:"path"`
	}{path}
	data, err := s.CallRaw("Init.checkMessageStorePath", params)
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
//	isEnabled - flag if statistics are enabled
func (s *ServerConnection) InitSetClientStatistics(isEnabled bool) error {
	params := struct {
		IsEnabled bool `json:"isEnabled"`
	}{isEnabled}
	_, err := s.CallRaw("Init.setClientStatistics", params)
	return err
}

// InitFinish - Finish initial configuration of Kerio Connect.
func (s *ServerConnection) InitFinish() error {
	_, err := s.CallRaw("Init.finish", nil)
	return err
}

// InitGetNamedConstantList - Server side list of constants.
// Return
//	constants - list of constants
func (s *ServerConnection) InitGetNamedConstantList() (NamedConstantList, error) {
	data, err := s.CallRaw("Init.getNamedConstantList", nil)
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
func (s *ServerConnection) InitGetBrowserLanguages() (StringList, error) {
	data, err := s.CallRaw("Init.getBrowserLanguages", nil)
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
func (s *ServerConnection) InitGetProductInfo() (*ProductInfo, error) {
	data, err := s.CallRaw("Init.getProductInfo", nil)
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
func (s *ServerConnection) InitGetEula() (string, error) {
	data, err := s.CallRaw("Init.getEula", nil)
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
