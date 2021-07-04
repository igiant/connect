package connect

import "encoding/json"

// Entity - Available entities, entity prefix due to name collision
type Entity string

const (
	EntityUser           Entity = "EntityUser"           // User Entity
	EntityAlias          Entity = "EntityAlias"          // Alias Entity
	EntityGroup          Entity = "EntityGroup"          // Group Entity
	EntityMailingList    Entity = "EntityMailingList"    // Mailing List Entity
	EntityResource       Entity = "EntityResource"       // Resource Scheduling Entity
	EntityTimeRange      Entity = "EntityTimeRange"      // Time Range Entity
	EntityTimeRangeGroup Entity = "EntityTimeRangeGroup" // Time Range Group Entity
	EntityIpAddress      Entity = "EntityIpAddress"      // Ip Address Entity
	EntityIpAddressGroup Entity = "EntityIpAddressGroup" // Ip Address Group Entity
	EntityService        Entity = "EntityService"        // Service Entity
	EntityDomain         Entity = "EntityDomain"
)

// RestrictionTuple - Restriction Items
type RestrictionTuple struct {
	Name   ItemName        `json:"name"`
	Kind   RestrictionKind `json:"kind"`
	Values StringList      `json:"values"`
}

// RestrictionTupleList - Restriction tuple for 1 entity
type RestrictionTupleList []RestrictionTuple

// Restriction - Entity name restriction definition
type Restriction struct {
	EntityName Entity               `json:"entityName"` // IDL entity name, eg. User
	Tuples     RestrictionTupleList `json:"tuples"`     // Restriction tuples
}

// RestrictionList - List of restrictions
type RestrictionList []Restriction

// JavaScriptDate - JavaScript timestamp
type JavaScriptDate string

// SubscriptionInfo - Subscription information
type SubscriptionInfo struct {
	IsUnlimited      bool           `json:"isUnlimited"`      // is it a special license with expiration == never ?
	ShowAlert        bool           `json:"showAlert"`        // show subscription expiration alert
	RemainingDays    int            `json:"remainingDays"`    // days remaining to subscription expiration
	SubscriptionDate JavaScriptDate `json:"subscriptionDate"` // last date of subscription
}

// AboutInfo - About information
type AboutInfo struct {
	CurrentUsers    int              `json:"currentUsers"`    // number of created users on domain
	AllowedUsers    MaximumUsers     `json:"allowedUsers"`    // number of allowed users, take stricter limit from max. number for domain, max. number by license
	ServerSoftware  string           `json:"serverSoftware"`  // product name and version string, same as SERVER_SOFTWARE
	Subscription    SubscriptionInfo `json:"subscription"`    // information about subscription
	Copyright       string           `json:"copyright"`       // copyright string
	ProductHomepage string           `json:"productHomepage"` // url to homepage of product
}

type ServerVersion struct {
	Product  string `json:"product"`  // name of product
	Version  string `json:"version"`  // version in string consists of values of major, minor, revision, build a dot separated
	Major    int    `json:"major"`    // major version
	Minor    int    `json:"minor"`    // minor version
	Revision int    `json:"revision"` // revision number
	Build    int    `json:"build"`    // build number
}

// AlertName - Type of Alert
type AlertName string

const (
	LicenseExpired                AlertName = "LicenseExpired"                // License has expired
	LicenseInvalidMinVersion      AlertName = "LicenseInvalidMinVersion"      // Invalid minimal version of a product found
	LicenseInvalidEdition         AlertName = "LicenseInvalidEdition"         // The license was not issued for this edition of the product
	LicenseInvalidUser            AlertName = "LicenseInvalidUser"            // The license was not issued for this user
	LicenseInvalidDomain          AlertName = "LicenseInvalidDomain"          // The license was not issued for this domain
	LicenseInvalidOS              AlertName = "LicenseInvalidOS"              // The license was not issued for this operating system
	LicenseCheckForwardingEnabled AlertName = "LicenseCheckForwardingEnabled" // The license was not alowed forward the message to another host
	LicenseTooManyUsers           AlertName = "LicenseTooManyUsers"           // More users try login to their mailboxes then allowed License.
	StorageSpaceLow               AlertName = "StorageSpaceLow"               // Low space in storage
	SubscriptionExpired           AlertName = "SubscriptionExpired"           // Subscription has expired
	SubscriptionSoonExpire        AlertName = "SubscriptionSoonExpire"        // Subscription soon expire
	LicenseSoonExpire             AlertName = "LicenseSoonExpire"             // License soon expire
	CoredumpFound                 AlertName = "CoredumpFound"                 // Some coredump was found after crash
	MacOSServicesKeepsPorts       AlertName = "MacOSServicesKeepsPorts"       // Apache on Lion server keeps ports (Eg. port 443), which are assigned to our services. See Services.stopMacOSServices()
	RemoteUpgradeFailed           AlertName = "RemoteUpgradeFailed"           // Remote server upgrade failed
	RemoteUpgradeSucceeded        AlertName = "RemoteUpgradeSucceeded"        // Remote server upgrade succeeded
)

// TypeAlert - Type of Alert
type TypeAlert string

const (
	Warning  TypeAlert = "Warning"
	Critical TypeAlert = "Critical"
	Info     TypeAlert = "Info"
)

// Alert - Alert
type Alert struct {
	AlertName     AlertName `json:"alertName"`     // Alert Id
	AlertType     TypeAlert `json:"alertType"`     // Alert type
	CurrentValue  string    `json:"currentValue"`  // Current Value
	CriticalValue string    `json:"criticalValue"` // Critical Value
}

type AlertList []Alert

// EntityDuplicate - Potential duplicate
type EntityDuplicate struct {
	Kind             Entity `json:"kind"` // which entity was found as first duplicate
	Name             string `json:"name"` // name of duplicate
	CollisionAddress string `json:"collisionAddress"`
	Win              bool   `json:"win"`       // if entity is winner in this collision of mail address
	IsPattern        bool   `json:"isPattern"` // is true if it is the pattern to check (self duplicity)
}

type EntityDuplicateList []EntityDuplicate

// EntityDetail - Detail about entity to be checked. Kind or id must be filled.
type EntityDetail struct {
	Kind Entity `json:"kind"` // which entity is inserting
	Id   KId    `json:"id"`   // entity global identification of updated entity
}

type UserNameList []string

type FolderInfo struct {
	FolderName     string       `json:"folderName"`
	ReferenceCount int          `json:"referenceCount"`
	IndexLoaded    bool         `json:"indexLoaded"`
	Users          UserNameList `json:"users"`
}

type FolderInfoList []FolderInfo

type WebSession struct {
	Id             string       `json:"id"`
	UserName       string       `json:"userName"`
	ClientAddress  string       `json:"clientAddress"`  // IPv4 address
	ExpirationTime string       `json:"expirationTime"` // format dd.mm.yyyy hh:mm:ss
	ComponentType  WebComponent `json:"componentType"`  // what about CalDav, WebDav, ActiveSync
	IsSecure       bool         `json:"isSecure"`       // is protocol secure
}

type WebSessionList []WebSession

type Protocol string

const (
	protocolAdmin      Protocol = "protocolAdmin"
	protocolSmtp       Protocol = "protocolSmtp"
	protocolSmtps      Protocol = "protocolSmtps"
	protocolSubmission Protocol = "protocolSubmission"
	protocolPop3       Protocol = "protocolPop3"
	protocolPop3s      Protocol = "protocolPop3s"
	protocolImap       Protocol = "protocolImap"
	protocolImaps      Protocol = "protocolImaps"
	protocolNntp       Protocol = "protocolNntp"
	protocolNntps      Protocol = "protocolNntps"
	protocolLdap       Protocol = "protocolLdap"
	protocolLdaps      Protocol = "protocolLdaps"
	protocolHttp       Protocol = "protocolHttp"
	protocolHttps      Protocol = "protocolHttps"
	protocolXmpp       Protocol = "protocolXmpp"
	protocolXmpps      Protocol = "protocolXmpps"
)

type HttpExtension string

const (
	NoExtension HttpExtension = "NoExtension"
	WebGeneric  HttpExtension = "WebGeneric" // WebMail or WebMail Mini or WebAdmin
	WebDav      HttpExtension = "WebDav"
	CalDav      HttpExtension = "CalDav"
	ActiveSync  HttpExtension = "ActiveSync"
	KocOffline  HttpExtension = "KocOffline"
	KBC         HttpExtension = "KBC" // Kerio Connector for BlackBerry Enterprise Server
	EWS         HttpExtension = "EWS" // Exchange Web Services
)

type Connection struct {
	Proto       Protocol      `json:"proto"`
	Extension   HttpExtension `json:"extension"`
	IsSecure    bool          `json:"isSecure"`
	Time        string        `json:"time"`
	From        string        `json:"from"`
	User        string        `json:"user"`
	Description string        `json:"description"`
}

type ConnectionList []Connection

// Administration - Note: isEnabled, isLimited and groupId fields must be assigned if any of them is used in set methods
type Administration struct {
	IsEnabled                   bool   `json:"isEnabled"`                   // administration from other that local machine is enabled/disabled
	IsLimited                   bool   `json:"isLimited"`                   // administration is limited
	GroupId                     KId    `json:"groupId"`                     // IP Address Group identifier on which is limit applied
	GroupName                   string `json:"groupName"`                   // [READ-ONLY] IP Address Group name on which is limit applied
	BuiltInAdminEnabled         bool   `json:"builtInAdminEnabled"`         // if is enabled field builtInAdminPassword is required
	BuiltInAdminUsername        string `json:"builtInAdminUsername"`        // [READ-ONLY] user name
	BuiltInAdminPassword        string `json:"builtInAdminPassword"`        // password
	BuiltInAdminPasswordIsEmpty bool   `json:"builtInAdminPasswordIsEmpty"` // [READ-ONLY] password is empty
	BuiltInAdminUsernameCollide bool   `json:"builtInAdminUsernameCollide"` // [READ-ONLY] username colide with user in primary domain
}

// ServerTimeInfo - Server time information
type ServerTimeInfo struct {
	TimezoneOffset int           `json:"timezoneOffset"` // +/- offset in minutes
	StartTime      DateTimeStamp `json:"startTime"`      // +/- start time of server
	CurrentTime    DateTimeStamp `json:"currentTime"`    // +/- current time on server
}

// ServerCreatePath - Server creates an archive/backup path. If credentials aren't specified, values from current configuration of backup are used.
//	path - new directory to create
//	credentials - (optional) user name and password required to access network disk
// Return
//	result - result of create action
func (s *ServerConnection) ServerCreatePath(path string, credentials Credentials) (*DirectoryAccessResult, error) {
	params := struct {
		Path        string      `json:"path"`
		Credentials Credentials `json:"credentials"`
	}{path, credentials}
	data, err := s.CallRaw("Server.createPath", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result DirectoryAccessResult `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return &result.Result.Result, err
}

// ServerFindEntityByEmail - caller must be authenticated; Note: creating duplicates is often allowed but may cause unwanted effects.
//	addresses - list of email addresses (without domain) to be checked
//	updatedEntity - identification of the current entity (to avoid self duplicity)
//	domainId - domain identification
// Return
//	entities - list of found entities with e-mail address duplicate 'updatedEntity' is included in list and marked, if none duplicate is found list is empty
func (s *ServerConnection) ServerFindEntityByEmail(addresses StringList, updatedEntity EntityDetail, domainId KId) (EntityDuplicateList, error) {
	params := struct {
		Addresses     StringList   `json:"addresses"`
		UpdatedEntity EntityDetail `json:"updatedEntity"`
		DomainId      KId          `json:"domainId"`
	}{addresses, updatedEntity, domainId}
	data, err := s.CallRaw("Server.findEntityByEmail", params)
	if err != nil {
		return nil, err
	}
	entities := struct {
		Result struct {
			Entities EntityDuplicateList `json:"entities"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &entities)
	return entities.Result.Entities, err
}

// ServerGenerateSupportInfo - Generate a file with information for the support.
// Return
//	fileDownload - description of output file
func (s *ServerConnection) ServerGenerateSupportInfo() (*Download, error) {
	data, err := s.CallRaw("Server.generateSupportInfo", nil)
	if err != nil {
		return nil, err
	}
	fileDownload := struct {
		Result struct {
			FileDownload Download `json:"fileDownload"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileDownload)
	return &fileDownload.Result.FileDownload, err
}

// ServerGetAboutInfo - Obtain information about server, caller must be authenticated.
// Return
//	aboutInformation - information about server
func (s *ServerConnection) ServerGetAboutInfo() (*AboutInfo, error) {
	data, err := s.CallRaw("Server.getAboutInfo", nil)
	if err != nil {
		return nil, err
	}
	aboutInformation := struct {
		Result struct {
			AboutInformation AboutInfo `json:"aboutInformation"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &aboutInformation)
	return &aboutInformation.Result.AboutInformation, err
}

// ServerGetAlertList - Obtain a list of alerts.
// Return
//	alerts - list of alerts
func (s *ServerConnection) ServerGetAlertList() (AlertList, error) {
	data, err := s.CallRaw("Server.getAlertList", nil)
	if err != nil {
		return nil, err
	}
	alerts := struct {
		Result struct {
			Alerts AlertList `json:"alerts"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &alerts)
	return alerts.Result.Alerts, err
}

// ServerGetBrowserLanguages - Returns a list of user-preferred languages set in browser.
// Return
//	calculatedLanguage - a list of 2-character language codes
func (s *ServerConnection) ServerGetBrowserLanguages() (StringList, error) {
	data, err := s.CallRaw("Server.getBrowserLanguages", nil)
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

// ServerGetClientStatistics - Obtain client statistics settings.
func (s *ServerConnection) ServerGetClientStatistics() (bool, error) {
	data, err := s.CallRaw("Server.getClientStatistics", nil)
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

// ServerGetColumnList - Obtain a list of columns dependent on callee role.
//	objectName - name of the API object
//	methodName - name of the method of appropriate object
// Return
//	columns - list of available columns
func (s *ServerConnection) ServerGetColumnList(objectName string, methodName string) (StringList, error) {
	params := struct {
		ObjectName string `json:"objectName"`
		MethodName string `json:"methodName"`
	}{objectName, methodName}
	data, err := s.CallRaw("Server.getColumnList", params)
	if err != nil {
		return nil, err
	}
	columns := struct {
		Result struct {
			Columns StringList `json:"columns"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &columns)
	return columns.Result.Columns, err
}

// ServerGetConnections - Obtain information about active connections.
//	query - condition and fields have no effect for this method
// Return
//	list - active connections
//  totalItems - total number of active connections
func (s *ServerConnection) ServerGetConnections(query SearchQuery) (ConnectionList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Server.getConnections", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ConnectionList `json:"list"`
			TotalItems int            `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ServerGetDirs - Obtain a list of directories in a particular path.
//	fullPath - directory for listing, if full path is empty logical drives will be listed
// Return
//	dirList - List of directories
func (s *ServerConnection) ServerGetDirs(fullPath string) (DirectoryList, error) {
	params := struct {
		FullPath string `json:"fullPath"`
	}{fullPath}
	data, err := s.CallRaw("Server.getDirs", params)
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

// ServerGetLicenseExtensionsList - Obtain list of license extensions, caller must be authenticated.
// Return
//	extensions - list of license extensions
func (s *ServerConnection) ServerGetLicenseExtensionsList() (StringList, error) {
	data, err := s.CallRaw("Server.getLicenseExtensionsList", nil)
	if err != nil {
		return nil, err
	}
	extensions := struct {
		Result struct {
			Extensions StringList `json:"extensions"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &extensions)
	return extensions.Result.Extensions, err
}

// ServerGetNamedConstantList - Server side list of constants.
// Return
//	constants - list of constants
func (s *ServerConnection) ServerGetNamedConstantList() (NamedConstantList, error) {
	data, err := s.CallRaw("Server.getNamedConstantList", nil)
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

// ServerGetOpenedFoldersInfo - Obtain information about folders opened on server.
//	query - condition and fields have no effect for this method
// Return
//	list - opened folders with info
//  totalItems - total number of opened folders
func (s *ServerConnection) ServerGetOpenedFoldersInfo(query SearchQuery) (FolderInfoList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Server.getOpenedFoldersInfo", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       FolderInfoList `json:"list"`
			TotalItems int            `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ServerGetProductInfo - Get basic information about product and its version.
// Return
//	info - structure with basic information about product
func (s *ServerConnection) ServerGetProductInfo() (*ProductInfo, error) {
	data, err := s.CallRaw("Server.getProductInfo", nil)
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

// ServerGetRemoteAdministration - Obtain information about remote administration settings.
// Return
//	setting - current settings
func (s *ServerConnection) ServerGetRemoteAdministration() (*Administration, error) {
	data, err := s.CallRaw("Server.getRemoteAdministration", nil)
	if err != nil {
		return nil, err
	}
	setting := struct {
		Result struct {
			Setting Administration `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return &setting.Result.Setting, err
}

// ServerGetServerHash - Obtain a hash string created from product name, version, and installation time.
// Return
//	serverHash - server hash
func (s *ServerConnection) ServerGetServerHash() (string, error) {
	data, err := s.CallRaw("Server.getServerHash", nil)
	if err != nil {
		return "", err
	}
	serverHash := struct {
		Result struct {
			ServerHash string `json:"serverHash"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &serverHash)
	return serverHash.Result.ServerHash, err
}

// ServerGetServerIpAddresses - List all server IP addresses.
// Return
//	addresses - all IP addresses of the server
func (s *ServerConnection) ServerGetServerIpAddresses() (StringList, error) {
	data, err := s.CallRaw("Server.getServerIpAddresses", nil)
	if err != nil {
		return nil, err
	}
	addresses := struct {
		Result struct {
			Addresses StringList `json:"addresses"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &addresses)
	return addresses.Result.Addresses, err
}

// ServerGetServerTime - Get server time information.
// Return
//	info - structure with time information
func (s *ServerConnection) ServerGetServerTime() (*ServerTimeInfo, error) {
	data, err := s.CallRaw("Server.getServerTime", nil)
	if err != nil {
		return nil, err
	}
	info := struct {
		Result struct {
			Info ServerTimeInfo `json:"info"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &info)
	return &info.Result.Info, err
}

// ServerGetVersion - Obtain information about server version.
// Return
//	product - name of product
//	version - version in string consists of values of major, minor, revision, build a dot separated
//	major - major version
//	minor - minor version
//	revision - revision number
//	build - build number
func (s *ServerConnection) ServerGetVersion() (*ServerVersion, error) {
	data, err := s.CallRaw("Server.getVersion", nil)
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

// ServerGetWebSessions - Obtain information about web component sessions.
//	query - condition and fields have no effect for this method
// Return
//	list - web component sessions
//  totalItems - total number of web component sessions
func (s *ServerConnection) ServerGetWebSessions(query SearchQuery) (WebSessionList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Server.getWebSessions", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       WebSessionList `json:"list"`
			TotalItems int            `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ServerKillWebSessions - Terminate actual web sessions.
//	ids - list of web sessions IDs to terminate
func (s *ServerConnection) ServerKillWebSessions(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("Server.killWebSessions", params)
	return err
}

// ServerPathExists - Check if the selected path exists and is accessible from the server.
//	path - directory name
//	credentials - (optional) user name and password required to access network disk
// Return
//	result - result of check
func (s *ServerConnection) ServerPathExists(path string, credentials Credentials) (DirectoryAccessResult, error) {
	params := struct {
		Path        string      `json:"path"`
		Credentials Credentials `json:"credentials"`
	}{path, credentials}
	data, err := s.CallRaw("Server.pathExists", params)
	if err != nil {
		return "", err
	}
	result := struct {
		Result struct {
			Result DirectoryAccessResult `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Result.Result, err
}

// ServerReboot - Reboot the host system.
func (s *ServerConnection) ServerReboot() error {
	_, err := s.CallRaw("Server.reboot", nil)
	return err
}

// ServerRestart - Restart server. The server must run as service.
func (s *ServerConnection) ServerRestart() error {
	_, err := s.CallRaw("Server.restart", nil)
	return err
}

// ServerUpgrade - Upgrade server to the latest version. The server must run as service.
func (s *ServerConnection) ServerUpgrade() error {
	_, err := s.CallRaw("Server.upgrade", nil)
	return err
}

// ServerGetDownloadProgress - Get progress of installation package downloading
// Return
//	progress - download progress in percents (0-100)
func (s *ServerConnection) ServerGetDownloadProgress() (int, error) {
	data, err := s.CallRaw("Server.getDownloadProgress", nil)
	if err != nil {
		return 0, err
	}
	progress := struct {
		Result struct {
			Progress int `json:"progress"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &progress)
	return progress.Result.Progress, err
}

// ServerSendBugReport - Send a bug report to Kerio.
//	name - name of sender
//	email - email of sender
//	language - language of report
//	subject - summary of report
//	description - description of problem
func (s *ServerConnection) ServerSendBugReport(name string, email string, language string, subject string, description string) error {
	params := struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Language    string `json:"language"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}{name, email, language, subject, description}
	_, err := s.CallRaw("Server.sendBugReport", params)
	return err
}

// ServerSetClientStatistics - Set client statistics settings.
func (s *ServerConnection) ServerSetClientStatistics(isEnabled bool) error {
	params := struct {
		IsEnabled bool `json:"isEnabled"`
	}{isEnabled}
	_, err := s.CallRaw("Server.setClientStatistics", params)
	return err
}

// ServerSetRemoteAdministration - Set new remote administration parameters.
//	setting - new settings
func (s *ServerConnection) ServerSetRemoteAdministration(setting Administration) error {
	params := struct {
		Setting Administration `json:"setting"`
	}{setting}
	_, err := s.CallRaw("Server.setRemoteAdministration", params)
	return err
}

// ServerUploadLicense - Upload license manually from a file.
//	fileId - ID of the uploaded file
func (s *ServerConnection) ServerUploadLicense(fileId string) error {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	_, err := s.CallRaw("Server.uploadLicense", params)
	return err
}

// ServerValidateRemoteAdministration - Validate whether the administrator can cut off him/herself from the administration.
//	setting - new setting
func (s *ServerConnection) ServerValidateRemoteAdministration(setting Administration) error {
	params := struct {
		Setting Administration `json:"setting"`
	}{setting}
	_, err := s.CallRaw("Server.validateRemoteAdministration", params)
	return err
}

// ServerIsBritishPreferred - Determine whether to use British or American flag for English.
// Return
//	preferred - use British flag
func (s *ServerConnection) ServerIsBritishPreferred() (bool, error) {
	data, err := s.CallRaw("Server.isBritishPreferred", nil)
	if err != nil {
		return false, err
	}
	preferred := struct {
		Result struct {
			Preferred bool `json:"preferred"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &preferred)
	return preferred.Result.Preferred, err
}
