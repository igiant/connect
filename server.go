package connect

import (
	"encoding/json"
)

type ServerVersion struct {
	Product  string `json:"product"`
	Version  string `json:"version"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
	Revision int    `json:"revision"`
	Build    int    `json:"build"`
}

type WebSessionList struct {
	ID             string `json:"id"`
	UserName       string `json:"userName"`
	ClientAddress  string `json:"clientAddress"`
	ExpirationTime string `json:"expirationTime"`
	ComponentType  string `json:"componentType"`
	IsSecure       bool   `json:"isSecure"`
}

type Connections struct {
	Proto       string `json:"proto"`
	Extension   string `json:"extension"`
	IsSecure    bool   `json:"isSecure"`
	Time        string `json:"time"`
	From        string `json:"from"`
	User        string `json:"user"`
	Description string `json:"description"`
}

type ConnectionList []Connections

type DirList struct {
	Name            string `json:"name"`
	HasSubdirectory bool   `json:"hasSubdirectory"`
}

type ExtensionsList []string

type Constant struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NamedConstantList []Constant

type FolderInfo struct {
	FolderName     string   `json:"folderName"`
	ReferenceCount int      `json:"referenceCount"`
	IndexLoaded    bool     `json:"indexLoaded"`
	Users          []string `json:"users"`
}

type FolderInfoList []FolderInfo

type ProductInfo struct {
	ProductName   string `json:"productName"`
	Version       string `json:"version"`
	BuildNumber   string `json:"buildNumber"`
	OsName        string `json:"osName"`
	Os            string `json:"os"`
	ReleaseType   string `json:"releaseType"`
	DeployedType  string `json:"deployedType"`
	IsDockerImage bool   `json:"isDockerImage"`
	UpdateInfo    struct {
		Result      string `json:"result"`
		Description string `json:"description"`
		DownloadURL string `json:"downloadUrl"`
		InfoURL     string `json:"infoUrl"`
	} `json:"updateInfo"`
	CentralManagementSet bool `json:"centralManagementSet"`
}

type Administration struct {
	IsEnabled                   bool   `json:"isEnabled"`
	IsLimited                   bool   `json:"isLimited"`
	GroupID                     string `json:"groupId"`
	GroupName                   string `json:"groupName"`
	BuiltInAdminEnabled         bool   `json:"builtInAdminEnabled"`
	BuiltInAdminUsername        string `json:"builtInAdminUsername"`
	BuiltInAdminPassword        string `json:"builtInAdminPassword"`
	BuiltInAdminPasswordIsEmpty bool   `json:"builtInAdminPasswordIsEmpty"`
	BuiltInAdminUsernameCollide bool   `json:"builtInAdminUsernameCollide"`
}

type ServerTimeInfo struct {
	TimezoneOffset int `json:"timezoneOffset"`
	StartTime      int `json:"startTime"`
	CurrentTime    int `json:"currentTime"`
}

// ServerGetVersion obtains a information about server version.
func (c *Connection) ServerGetVersion() (*ServerVersion, error) {
	data, err := c.CallRaw("Server.getVersion", nil)
	if err != nil {
		return nil, err
	}
	serverVersion := struct {
		Result struct {
			ServerVersion
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &serverVersion)
	return &serverVersion.Result.ServerVersion, err
}

// ServerGetServerHash obtains a hash string created from product name, version, and installation time.
func (c *Connection) ServerGetServerHash() (string, error) {
	data, err := c.CallRaw("Server.getServerHash", nil)
	if err != nil {
		return "", err
	}
	serverHash := ""
	err = json.Unmarshal(data, &serverHash)
	return serverHash, err
}

// ServerGetProductInfo gets basic information about product and its version.
func (c *Connection) ServerGetProductInfo() (*ProductInfo, error) {
	data, err := c.CallRaw("Server.getProductInfo", nil)
	if err != nil {
		return nil, err
	}
	productInfo := struct {
		Result struct {
			ProductInfo `json:"info"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &productInfo)
	return &productInfo.Result.ProductInfo, err
}

// ServerGetRemoteAdministration obtains a information about remote administration settings.
func (c *Connection) ServerGetRemoteAdministration() (*Administration, error) {
	data, err := c.CallRaw("Server.getRemoteAdministration", nil)
	if err != nil {
		return nil, err
	}
	administration := struct {
		Result struct {
			Administration `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &administration)
	return &administration.Result.Administration, err
}

// ServerGetBrowserLanguages returns a list of user-preferred languages set in browser.
func (c *Connection) ServerGetBrowserLanguages() ([]string, error) {
	data, err := c.CallRaw("Server.getBrowserLanguages", nil)
	if err != nil {
		return nil, err
	}
	calculatedLanguage := struct {
		Result struct {
			Languages []string `json:"calculatedLanguage"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &calculatedLanguage)
	return calculatedLanguage.Result.Languages, err
}

// ServerGetServerIpAddresses obtains a list all server IP addresses.
func (c *Connection) ServerGetServerIpAddresses() ([]string, error) {
	data, err := c.CallRaw("Server.getServerIpAddresses", nil)
	if err != nil {
		return nil, err
	}
	addresses := struct {
		Result struct {
			Addresses []string `json:"addresses"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &addresses)
	return addresses.Result.Addresses, err
}

// ServerGetColumnList obtains a list of columns dependent on callee role.
func (c *Connection) ServerGetColumnList(objectName, methodName string) ([]string, error) {
	params := struct {
		ObjectName string `json:"objectName"`
		MethodName string `json:"methodName"`
	}{objectName, methodName}
	data, err := c.CallRaw("Server.getColumnList", params)
	if err != nil {
		return nil, err
	}
	columnList := struct {
		Result struct {
			ColumnList []string `json:"columns"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &columnList)
	return columnList.Result.ColumnList, err
}

// ServerGetDownloadProgress obtains a progress of installation package downloading.
func (c *Connection) ServerGetDownloadProgress() (int, error) {
	data, err := c.CallRaw("Server.getDownloadProgress", nil)
	if err != nil {
		return 0, err
	}
	downloadProgress := struct {
		Result struct {
			Progress int `json:"progress"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &downloadProgress)
	return downloadProgress.Result.Progress, err
}

// ServerGetServerTime obtains server time information.
func (c *Connection) ServerGetServerTime() (*ServerTimeInfo, error) {
	data, err := c.CallRaw("Server.getServerTime", nil)
	if err != nil {
		return nil, err
	}
	serverTimeInfo := struct {
		Result struct {
			ServerTimeInfo `json:"info"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &serverTimeInfo)
	return &serverTimeInfo.Result.ServerTimeInfo, err
}

// ServerGetClientStatistics obtains client statistics settings.
func (c *Connection) ServerGetClientStatistics() (bool, error) {
	data, err := c.CallRaw("Server.getClientStatistics", nil)
	if err != nil {
		return false, err
	}
	clientStatistics := struct {
		Result struct {
			ClientStatistics bool `json:"isEnabled"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &clientStatistics)
	return clientStatistics.Result.ClientStatistics, err
}

// ServerGetLicenseExtensionsList obtains a list of license extensionsList, caller must be authenticated.
func (c *Connection) ServerGetLicenseExtensionsList() (ExtensionsList, error) {
	data, err := c.CallRaw("Server.getLicenseExtensionsList", nil)
	if err != nil {
		return nil, err
	}
	extensions := struct {
		Result struct {
			ExtensionsList `json:"extensions"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &extensions)
	return extensions.Result.ExtensionsList, err
}

// ServerGetNamedConstantList obtains server side list of constants.
func (c *Connection) ServerGetNamedConstantList() (NamedConstantList, error) {
	data, err := c.CallRaw("Server.getNamedConstantList", nil)
	if err != nil {
		return nil, err
	}
	namedConstantList := struct {
		Result struct {
			NamedConstantList `json:"constants"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &namedConstantList)
	return namedConstantList.Result.NamedConstantList, err
}

// ServerGetWebSessions obtains a information about web component sessions.
func (c *Connection) ServerGetWebSessions(query SearchQuery) ([]WebSessionList, error) {
	q := struct {
		SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Server.getWebSessions", q)
	if err != nil {
		return nil, err
	}
	webSessionList := struct {
		Result struct {
			WebSessions []WebSessionList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &webSessionList)
	return webSessionList.Result.WebSessions, err
}

// ServerKillWebSessions Terminate actual web sessions.
func (c *Connection) ServerKillWebSessions(ids []string) error {
	params := struct {
		IDs []string `json:"ids"`
	}{ids}
	_, err := c.CallRaw("Server.killWebSessions", params)
	return err
}

// ServerSendBugReport send a bug report to Kerio.
func (c *Connection) ServerSendBugReport(name, email, language, subject, description string) error {
	params := struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Language    string `json:"language"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}{name, email, language, subject, description}
	_, err := c.CallRaw("Server.sendBugReport", params)
	return err
}

// ServerGetConnections obtains a information about active connections.
func (c *Connection) ServerGetConnections(query SearchQuery) (ConnectionList, error) {
	q := struct {
		SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Server.getConnections", q)
	if err != nil {
		return nil, err
	}
	connectionList := struct {
		Result struct {
			ConnectionList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &connectionList)
	return connectionList.Result.ConnectionList, err
}

// ServerGetOpenedFoldersInfo obtains a information about folders opened on server.
func (c *Connection) ServerGetOpenedFoldersInfo(query SearchQuery) (FolderInfoList, error) {
	q := struct {
		SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Server.getOpenedFoldersInfo", q)
	if err != nil {
		return nil, err
	}
	folderInfoList := struct {
		Result struct {
			FolderInfoList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &folderInfoList)
	return folderInfoList.Result.FolderInfoList, err
}

// ServerGetDirs obtains a list of directories in a particular path.
func (c *Connection) ServerGetDirs(path string) ([]DirList, error) {
	params := struct {
		FullPath string `json:"fullPath"`
	}{path}
	data, err := c.CallRaw("Server.getDirs", params)
	if err != nil {
		return nil, err
	}
	dirList := struct {
		Result struct {
			Dirs []DirList `json:"dirList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dirList)
	return dirList.Result.Dirs, err
}

//ServerPathExists checks if the selected path exists and is accessible from the server.
//Parameters:
//      path	    - directory name
//      credentials	- (optional) user name and password required to access network disk
//      result	    - result of check
func (c *Connection) ServerPathExists(username, password, path string) (string, error) {
	params := struct {
		Credentials Credentials `json:"credentials"`
		Path        string      `json:"path"`
	}{Credentials{username, password}, path}
	data, err := c.CallRaw("Server.pathExists", params)
	if err != nil {
		return "", err
	}
	directoryAccessResult := struct {
		Result struct {
			Result string `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &directoryAccessResult)
	return directoryAccessResult.Result.Result, err
}

// ServerReboot - reboot the host system
func (c *Connection) ServerReboot() error {
	_, err := c.CallRaw("Server.reboot", nil)
	return err
}

// ServerRestart - restart server. The server must run as service.
func (c *Connection) ServerRestart() error {
	_, err := c.CallRaw("Server.restart", nil)
	return err
}

// ServerUpgrade - upgrade server to the latest version. The server must run as service.
func (c *Connection) ServerUpgrade() error {
	_, err := c.CallRaw("Server.upgrade", nil)
	return err
}
