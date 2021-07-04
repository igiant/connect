package connect

import "encoding/json"

type PublicFolder struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

type PublicFolderList []PublicFolder

// UserEmailAddressList - List of email addresses
type UserEmailAddressList []string

// FileFormatType - Export format type.
type FileFormatType string

const (
	TypeXml FileFormatType = "TypeXml" // Extensible Markup Language
	TypeCsv FileFormatType = "TypeCsv" // Comma Separated Values
)

// UserRoleType - Type of user role.
type UserRoleType string

const (
	UserRole           UserRoleType = "UserRole"           // regular user without any administration rights
	Auditor            UserRoleType = "Auditor"            // read only access to administration
	AccountAdmin       UserRoleType = "AccountAdmin"       // can administer Users,Groups,Aliases,MLs
	FullAdmin          UserRoleType = "FullAdmin"          // unlimited administration
	BuiltInAdmin       UserRoleType = "BuiltInAdmin"       // BuiltIn admin role can be returned only in Session::WhoAmI method for built-in administrator. This role must NOT be assigned.
	BuiltInDomainAdmin UserRoleType = "BuiltInDomainAdmin" // BuiltIn domain admin role can be returned only in Session::WhoAmI method for built-in domain administrator. This role must NOT be assigned.
)

// UserRight - Note: all fields must be assigned if used in set methods.
type UserRight struct {
	UserRole           UserRoleType `json:"userRole"`
	PublicFolderRight  bool         `json:"publicFolderRight"`
	ArchiveFolderRight bool         `json:"archiveFolderRight"`
}

// UserForwardMode - Forwarding setup for user.
type UserForwardMode string

const (
	UForwardNone    UserForwardMode = "UForwardNone"    // Forwarding is disabled
	UForwardYes     UserForwardMode = "UForwardYes"     // Forward all messages for this user to some addresses, don't deliver the message to the mailbox.
	UForwardDeliver UserForwardMode = "UForwardDeliver" // Forward all messages for this user to some addresses, and also deliver the message to user's mailbox.
)

// UserDeleteFolderMode - Type of deleting folder of the user
type UserDeleteFolderMode string

const (
	UDeleteUser   UserDeleteFolderMode = "UDeleteUser"   // Delete user without deleting his folder.
	UDeleteFolder UserDeleteFolderMode = "UDeleteFolder" // Delete user and delete his folder.
	UMoveFolder   UserDeleteFolderMode = "UMoveFolder"   // Delete user and his folder will move into another user's folder.
)

// EmailForwarding - Settings of email forwarding.
// Note: all fields must be assigned if used in set methods.
type EmailForwarding struct {
	Mode           UserForwardMode      `json:"mode"`
	EmailAddresses UserEmailAddressList `json:"emailAddresses"` // list of email addresses, make sense only for UForwardDeliver
}

// UserGroup - Properties of user's groups.
type UserGroup struct {
	Id          KId        `json:"id"` // global identification
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ItemSource  DataSource `json:"itemSource"`
}

// UserGroupList - List of user's groups.
type UserGroupList []UserGroup

// ItemCountLimit - Settings of items limit.
// Note: all fields must be assigned if used in set methods.
type ItemCountLimit struct {
	IsActive bool `json:"isActive"`
	Limit    int  `json:"limit"`
}

// QuotaUsage - Amount of storage used and items currently stored in user's store.
type QuotaUsage struct {
	Items   int                `json:"items"`
	Storage ByteValueWithUnits `json:"storage"`
}

// QuotaUsageList - List of QuotaUsage.
type QuotaUsageList []QuotaUsage

// LastLogin - Last login information.
type LastLogin struct {
	DateTime DateTimeStamp `json:"dateTime"` // date and time of last login
	Protocol string        `json:"protocol"` // protocol name of last login, example POP3
}

// CleanOut - Per-user message retention policy.
type CleanOut struct {
	IsUsedDomain bool            `json:"isUsedDomain"` // use domain settings
	DeletedItems ActionAfterDays `json:"deletedItems"` // clean Deleted Items folder (maximum: 24855)
	JunkEmail    ActionAfterDays `json:"junkEmail"`    // clean Junk Email folder (maximum: 24855)
	SentItems    ActionAfterDays `json:"sentItems"`    // clean Sent Items folder (maximum: 24855)
	AutoDelete   ActionAfterDays `json:"autoDelete"`   // clean all folders (maximum: 24855)
}

// User - User details.
type User struct {
	Id                   KId                  `json:"id"`                   // [READ-ONLY] global identification
	DomainId             KId                  `json:"domainId"`             // [REQUIRED FOR CREATE] ID of domain where user belongs to
	CompanyContactId     KId                  `json:"companyContactId"`     // ID of company contact associated with this user
	LoginName            string               `json:"loginName"`            // [REQUIRED FOR CREATE] [USED BY QUICKSEARCH] loginName name
	FullName             string               `json:"fullName"`             // [USED BY QUICKSEARCH]
	Description          string               `json:"description"`          // [USED BY QUICKSEARCH]
	IsEnabled            bool                 `json:"isEnabled"`            // user account is enabled/disabled
	ItemSource           DataSource           `json:"itemSource"`           // is user stored internally or by LDAP? This field cannot be used with Or queries.
	AuthType             UserAuthType         `json:"authType"`             // supported values must be retrieved from engine by ServerInfo::getSupportedAuthTypes()
	Password             string               `json:"password"`             // [WRITE-ONLY]
	IsPasswordReversible bool                 `json:"isPasswordReversible"` // typically triple DES
	AllowPasswordChange  bool                 `json:"allowPasswordChange"`  // if it is set to false the password can be changed only by the administrator
	HasDefaultSpamRule   bool                 `json:"hasDefaultSpamRule"`   // now: available only on user creation
	Role                 UserRight            `json:"role"`                 // user role
	GroupRole            UserRight            `json:"groupRole"`            // the mightiest user role obtained via group membership
	EffectiveRole        UserRight            `json:"effectiveRole"`        // the mightiest user role from role and groupRole
	IsWritableByMe       bool                 `json:"isWritableByMe"`       // Does caller have right to change the user? E.g. if Account Admin gets User structure for Full Admin, isWritableByMe will be false. This field is read-only and cannot be used in SearchQuery conditions.
	EmailAddresses       UserEmailAddressList `json:"emailAddresses"`       // List of user email addresses. His default one (loginName@domain) is not listed here
	EmailForwarding      EmailForwarding      `json:"emailForwarding"`      // email forwarding setting
	UserGroups           UserGroupList        `json:"userGroups"`           // groups membership
	ItemLimit            ItemCountLimit       `json:"itemLimit"`            // max. number of items
	DiskSizeLimit        SizeLimit            `json:"diskSizeLimit"`        // max. disk usage
	ConsumedItems        int                  `json:"consumedItems"`        // current items used
	ConsumedSize         ByteValueWithUnits   `json:"consumedSize"`         // current disk usage
	HasDomainRestriction bool                 `json:"hasDomainRestriction"` // user can send/receive from/to his/her domain only
	OutMessageLimit      SizeLimit            `json:"outMessageLimit"`      // limit of outgoing message
	LastLoginInfo        LastLogin            `json:"lastLoginInfo"`        // information about last login datetime and protocol
	PublishInGal         bool                 `json:"publishInGal"`         // publish user in global address list? Default is true - the user will be published in Global Address Book.
	CleanOutItems        CleanOut             `json:"cleanOutItems"`        // Items clean-out settings
	AccessPolicy         IdEntity             `json:"accessPolicy"`         // ID and name of Access Policy applied for user. Only ID is writable.
	HomeServer           HomeServer           `json:"homeServer"`           // [WRITE-ONCE] Id of user's homeserver if server is in a distributed domain.
	Migration            OptionalEntity       `json:"migration"`            // [READ-ONLY] migration.enabled is true if user's store is just being migrated and migration.id contains migration task id
}

// UserList - List of users.
type UserList []User

// EffectiveUserRights - User effective rights (inherited from groups)
type EffectiveUserRights struct {
	UserId               KId  `json:"userId"`               // [READ-ONLY] global identification
	HasDomainRestriction bool `json:"hasDomainRestriction"` // user can send/receive from/to his/her domain only
}

// EffectiveUserRightsList - List of users effective rights
type EffectiveUserRightsList []EffectiveUserRights

// ServerDirectoryType - Type of user directory
type ServerDirectoryType string

const (
	WinNT            ServerDirectoryType = "WinNT"            // Windows NT Domain directory (Win NT 4.0)
	ActiveDirectory  ServerDirectoryType = "ActiveDirectory"  // Active Directory (Windows 2000 and newer)
	NovellEDirectory ServerDirectoryType = "NovellEDirectory" // Novell eDirectory
)

// ImportServer - Properties of the server from which users are imported.
type ImportServer struct {
	DirectoryType      ServerDirectoryType `json:"directoryType"`
	RemoteDomainName   string              `json:"remoteDomainName"`
	Address            string              `json:"address"` // server IP or FQDN
	LoginName          string              `json:"loginName"`
	Password           string              `json:"password"`
	LdapFilter         string              `json:"ldapFilter"`
	IsSecureConnection bool                `json:"isSecureConnection"`
}

// LoginStats - Login statistics - count and timestamp of the last login.
type LoginStats struct {
	Count     int    `json:"count"`
	LastLogin string `json:"lastLogin"`
}

// UserStats - Statistics about user's usage of quota, logins to different services.
type UserStats struct {
	Name             string     `json:"name"` // user's loginName
	OccupiedSpace    QuotaUsage `json:"occupiedSpace"`
	Pop3             LoginStats `json:"pop3"`
	SecurePop3       LoginStats `json:"securePop3"`
	Imap             LoginStats `json:"imap"`
	SecureImap       LoginStats `json:"secureImap"`
	Http             LoginStats `json:"http"`
	SecureHttp       LoginStats `json:"secureHttp"`
	Ldap             LoginStats `json:"ldap"`
	SecureLdap       LoginStats `json:"secureLdap"`
	Nntp             LoginStats `json:"nntp"`
	SecureNntp       LoginStats `json:"secureNntp"`
	ActiveSync       LoginStats `json:"activeSync"`
	SecureActiveSync LoginStats `json:"secureActiveSync"`
	Xmpp             LoginStats `json:"xmpp"`
	SecureXmpp       LoginStats `json:"secureXmpp"`
}

// UserStatList - List of users' statistics.
type UserStatList []UserStats

// ResultTriplet - Result of a mass operation.
type ResultTriplet struct {
	InputIndex int `json:"inputIndex"`
	ItemsCount int `json:"itemsCount"`
}

// ResultTripletList - List of mass operation results.
type ResultTripletList []ResultTriplet

// RemovalRequest - User to be removed, what to do with his/her mailbox.
type RemovalRequest struct {
	UserId           KId                        `json:"userId"`           // ID of user to be removed
	Method           UserDeleteFolderMode       `json:"method"`           // removal method
	RemoveReferences bool                       `json:"removeReferences"` // if true all reference to this user is going to be removed as well
	TargetUserId     KId                        `json:"targetUserId"`     // applicable only when moving user's store to another user, use empty string if not moving user's messages to target mailbox
	Mode             DirectoryServiceDeleteMode `json:"mode"`             // delete mode
}

type RemovalRequestList []RemovalRequest

// Importee - A user being imported from directory server.
type Importee struct {
	UserItem     User   `json:"userItem"`     // user data
	IsImportable bool   `json:"isImportable"` // [READ-ONLY] user can be imported
	Message      string `json:"message"`      // [READ-ONLY] error message if user is not importable
}

type ImporteeList []Importee

type MailboxCount struct {
	Active int `json:"active"` // the number of active mailboxes on server
	Total  int `json:"total"`  // the number of created users on server
}

// AuthResult - Resut of autentication.
type AuthResult string

const (
	AuthOK           AuthResult = "AuthOK"           // User was autenticated
	AuthFail         AuthResult = "AuthFail"         // Wrong login name or password.
	AuthUserDisabled AuthResult = "AuthUserDisabled" // User cannot to log in, because his account is disabled.
	AuthLicense      AuthResult = "AuthLicense"      // User cannot log in, because license limit was reached.
	AuthDenied       AuthResult = "AuthDenied"       // User is denied to log in.
	AuthTryLater     AuthResult = "AuthTryLater"     // User cannot to log in at this moment, try later.
)

// User accounts management.

// UsersActivate - Activate user(s) from a directory service.
// Parameters
//	userIds - list of global user identifiers
// Return
//	errors - list of error messages for appropriate users
func (s *ServerConnection) UsersActivate(userIds KIdList) (ErrorList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := s.CallRaw("Users.activate", params)
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

// UsersActivateOnServer - Activate user(s) from a directory service in distributed domain environment.
// Parameters
//	userIds - list of global user identifiers
//	homeServerId - Id of server in distributed domain on which users will be activated
// Return
//	errors - list of error messages for appropriate users
func (s *ServerConnection) UsersActivateOnServer(userIds KIdList, homeServerId KId) (ErrorList, error) {
	params := struct {
		UserIds      KIdList `json:"userIds"`
		HomeServerId KId     `json:"homeServerId"`
	}{userIds, homeServerId}
	data, err := s.CallRaw("Users.activateOnServer", params)
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

// UsersConnectFromExternalService - Register connection.
// Parameters
//	service - service name (should be some real service ID, returned by Services.get)
//	connectionId - unique connection identifier
//	port - host port
//	isSecure - ssl connection
func (s *ServerConnection) UsersConnectFromExternalService(service string, connectionId string, clientIpAddress string, port int, isSecure bool) (bool, error) {
	params := struct {
		Service         string `json:"service"`
		ConnectionId    string `json:"connectionId"`
		ClientIpAddress string `json:"clientIpAddress"`
		Port            int    `json:"port"`
		IsSecure        bool   `json:"isSecure"`
	}{service, connectionId, clientIpAddress, port, isSecure}
	data, err := s.CallRaw("Users.connectFromExternalService", params)
	if err != nil {
		return false, err
	}
	result := struct {
		Result struct {
			Result bool `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Result.Result, err
}

// UsersAuthenticateConnectionFromExternalService - Authenticate given user and create session. connectionId must be registered by function connectFromExternalService otherwise authenticate fails.
// Parameters
//	userName - login name + domain name (can be omitted if primary) of the user to be logged in, e.g. "jdoe" or "jdoe@company.com"
//	password - password of the user to be authenticate (base64-encoded)
//	connectionId - connection identifier, must be the same as in connectFromExternalService
//	isSecure - ssl connection
// Return
//	result - resut of autentication.
func (s *ServerConnection) UsersAuthenticateConnectionFromExternalService(userName string, password string, service string, connectionId string, isSecure bool) (*AuthResult, error) {
	params := struct {
		UserName     string `json:"userName"`
		Password     string `json:"password"`
		Service      string `json:"service"`
		ConnectionId string `json:"connectionId"`
		IsSecure     bool   `json:"isSecure"`
	}{userName, password, service, connectionId, isSecure}
	data, err := s.CallRaw("Users.authenticateConnectionFromExternalService", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result AuthResult `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return &result.Result.Result, err
}

// UsersDisconnectFromExternalService - Unregister connection registered by connectFromExternalService and destroy session created if authenticateFromExternalService was called.
// Parameters
//	service - service name
//	connectionId - unique connection identifier
func (s *ServerConnection) UsersDisconnectFromExternalService(service string, connectionId string) error {
	params := struct {
		Service      string `json:"service"`
		ConnectionId string `json:"connectionId"`
	}{service, connectionId}
	_, err := s.CallRaw("Users.disconnectFromExternalService", params)
	return err
}

// UsersCancelWipeMobileDevice - Cancel wiping of user's mobile device.
// Parameters
//	userId - global user identifier
//	deviceId - ID of user's mobile device to cancel wipe
func (s *ServerConnection) UsersCancelWipeMobileDevice(userId KId, deviceId string) error {
	params := struct {
		UserId   KId    `json:"userId"`
		DeviceId string `json:"deviceId"`
	}{userId, deviceId}
	_, err := s.CallRaw("Users.cancelWipeMobileDevice", params)
	return err
}

// UsersCheckMailboxIntegrity - If corrupted folder is found, try to fix it.
// Parameters
//	userIds - list of user identifiers
func (s *ServerConnection) UsersCheckMailboxIntegrity(userIds KIdList) error {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	_, err := s.CallRaw("Users.checkMailboxIntegrity", params)
	return err
}

// UsersCreate - Create new users.
// Parameters
//	users - new user entities
// Return
//	errors - error message list
//	result - list of IDs of created users
func (s *ServerConnection) UsersCreate(users UserList) (ErrorList, CreateResultList, error) {
	params := struct {
		Users UserList `json:"users"`
	}{users}
	data, err := s.CallRaw("Users.create", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// UsersCreateLdap - Create new users in directory service
// Parameters
//	users - new user entities
// Return
//	errors - error message list
//	result - list of IDs of created users
func (s *ServerConnection) UsersCreateLdap(users UserList) (ErrorList, CreateResultList, error) {
	params := struct {
		Users UserList `json:"users"`
	}{users}
	data, err := s.CallRaw("Users.createLdap", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// UsersExportStatistics - Export statistics of given users in given format.
// Parameters
//	userIds - list of IDs of given users
//	format - output data format
// Return
//	fileDownload - description of output file
func (s *ServerConnection) UsersExportStatistics(userIds KIdList, format FileFormatType) (*Download, error) {
	params := struct {
		UserIds KIdList        `json:"userIds"`
		Format  FileFormatType `json:"format"`
	}{userIds, format}
	data, err := s.CallRaw("Users.exportStatistics", params)
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

// UsersExportToCsv - Export given domain users to comma-separated values file format.
// Parameters
//	filename - part of filename; full filename is compound as user_<domainname>_<filename>_<date>.csv
//	query - query attributes and limits
//	domainId - domain identification
// Return
//	fileDownload - description of output file
func (s *ServerConnection) UsersExportToCsv(filename string, query SearchQuery, domainId KId) (*Download, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Filename string      `json:"filename"`
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{filename, query, domainId}
	data, err := s.CallRaw("Users.exportToCsv", params)
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

// UsersGet - Obtain a list of users in given domain.
// Parameters
//	query - query attributes and limits
//	domainId - domain identification
// Return
//	list - users
//  totalItems - number of users found in given domain
func (s *ServerConnection) UsersGet(query SearchQuery, domainId KId) (UserList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Users.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       UserList `json:"list"`
			TotalItems int      `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// UsersGetContactPublicFolderList - Obtain a list of contact public folders in given domain.
// Parameters
//	domainId - global identification of domain
func (s *ServerConnection) UsersGetContactPublicFolderList(domainId KId) (PublicFolderList, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("Users.getContactPublicFolderList", params)
	if err != nil {
		return nil, err
	}
	publicFolders := struct {
		Result struct {
			PublicFolders PublicFolderList `json:"publicFolders"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &publicFolders)
	return publicFolders.Result.PublicFolders, err
}

// UsersGetFromServer - Obtain list of users from given LDAP server potentially importable to the Connect server.
// Parameters
//	importServer - properties of the server to import from
//	domainToImport - the mailserver domain where users are imported
// Return
//	newUsers - list of users
func (s *ServerConnection) UsersGetFromServer(importServer ImportServer, domainToImport KId) (ImporteeList, error) {
	params := struct {
		ImportServer   ImportServer `json:"importServer"`
		DomainToImport KId          `json:"domainToImport"`
	}{importServer, domainToImport}
	data, err := s.CallRaw("Users.getFromServer", params)
	if err != nil {
		return nil, err
	}
	newUsers := struct {
		Result struct {
			NewUsers ImporteeList `json:"newUsers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &newUsers)
	return newUsers.Result.NewUsers, err
}

// UsersGetMailboxCount - This method may take a long time if a directory service for mapped users is not available.
// Return
//	count - Number of users created on the server and number of active mailboxes.
func (s *ServerConnection) UsersGetMailboxCount() (*MailboxCount, error) {
	data, err := s.CallRaw("Users.getMailboxCount", nil)
	if err != nil {
		return nil, err
	}
	count := struct {
		Result struct {
			Count MailboxCount `json:"count"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &count)
	return &count.Result.Count, err
}

// UsersGetMobileDeviceList - Obtain a list of mobile devices of given user.
// Parameters
//	userId - name of user
//	query - query attributes and limits
// Return
//	list - mobile devices of given user
//  totalItems - number of mobile devices found for given user
func (s *ServerConnection) UsersGetMobileDeviceList(userId KId, query SearchQuery) (MobileDeviceList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		UserId KId         `json:"userId"`
		Query  SearchQuery `json:"query"`
	}{userId, query}
	data, err := s.CallRaw("Users.getMobileDeviceList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       MobileDeviceList `json:"list"`
			TotalItems int              `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// UsersGetNotActivated - Only user's ID, loginName, fullName, description are set in structures.
// Parameters
//	domainId - global identification of domain
// Return
//	newUsers - list of users
func (s *ServerConnection) UsersGetNotActivated(domainId KId) (ImporteeList, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("Users.getNotActivated", params)
	if err != nil {
		return nil, err
	}
	newUsers := struct {
		Result struct {
			NewUsers ImporteeList `json:"newUsers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &newUsers)
	return newUsers.Result.NewUsers, err
}

// UsersGetRecoveryDeletedItemsSize - Obtain a size of items stored for recovering.
// Parameters
//	userIds - global identification of user
// Return
//	errors - error message list
func (s *ServerConnection) UsersGetRecoveryDeletedItemsSize(userIds KIdList) (ErrorList, QuotaUsageList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := s.CallRaw("Users.getRecoveryDeletedItemsSize", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors   ErrorList      `json:"errors"`
			SizeList QuotaUsageList `json:"sizeList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.SizeList, err
}

// UsersGetStatistics - Obtain statistics of given users.
// Parameters
//	userIds - list of IDs of given users
//	query - query parameters and limits
// Return
//	list - users' statistics
func (s *ServerConnection) UsersGetStatistics(userIds KIdList, query SearchQuery) (UserStatList, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		UserIds KIdList     `json:"userIds"`
		Query   SearchQuery `json:"query"`
	}{userIds, query}
	data, err := s.CallRaw("Users.getStatistics", params)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List UserStatList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// UsersParseFromCsv - 'PRESERVE_FROM_CSV': preserve domain from CSV file (use primary if not defined)
// Parameters
//	fileId - ID of the uploaded file
//	domainToImport - import to given domain, magic constants
// Return
//	users - list of parsed users with appropriate status and message
func (s *ServerConnection) UsersParseFromCsv(fileId string, domainToImport KId) (ImporteeList, error) {
	params := struct {
		FileId         string `json:"fileId"`
		DomainToImport KId    `json:"domainToImport"`
	}{fileId, domainToImport}
	data, err := s.CallRaw("Users.parseFromCsv", params)
	if err != nil {
		return nil, err
	}
	users := struct {
		Result struct {
			Users ImporteeList `json:"users"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &users)
	return users.Result.Users, err
}

// UsersRecoverDeletedItems - If the user quota is exceeded an error with code 4000 will be returned.
// Parameters
//	userIds - list of user IDs
// Return
//	recoveryMessages - list of recovery messages
func (s *ServerConnection) UsersRecoverDeletedItems(userIds KIdList) (ErrorList, ResultTripletList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := s.CallRaw("Users.recoverDeletedItems", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors           ErrorList         `json:"errors"`
			RecoveryMessages ResultTripletList `json:"recoveryMessages"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.RecoveryMessages, err
}

// UsersRemove - Remove user(s).
// Parameters
//	requests - list of user IDs to be removed, method, and owner of deleted messages
// Return
//	errors - list of users failed to remove only (successfully removed are NOT listed)
func (s *ServerConnection) UsersRemove(requests RemovalRequestList) (ErrorList, error) {
	params := struct {
		Requests RemovalRequestList `json:"requests"`
	}{requests}
	data, err := s.CallRaw("Users.remove", params)
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

// UsersRemoveMobileDevice - Remove mobile device from the list of user's mobile devices.
// Parameters
//	userId - name of user
//	deviceId - ID of user's mobile device to be removed
func (s *ServerConnection) UsersRemoveMobileDevice(userId KId, deviceId string) error {
	params := struct {
		UserId   KId    `json:"userId"`
		DeviceId string `json:"deviceId"`
	}{userId, deviceId}
	_, err := s.CallRaw("Users.removeMobileDevice", params)
	return err
}

// UsersResetBuddyList - IM: Reset buddy list of selected users
// Parameters
//	userIds - list of user identifiers
func (s *ServerConnection) UsersResetBuddyList(userIds KIdList) error {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	_, err := s.CallRaw("Users.resetBuddyList", params)
	return err
}

// UsersGetEffectiveUserRights - Obtains user effective rights (inherited from groups)
// Parameters
//	userIds - list of IDs of users
// Return
//	errors - list of users failed to get effective rights
//	result - list of effective rights
func (s *ServerConnection) UsersGetEffectiveUserRights(userIds KIdList) (ErrorList, EffectiveUserRightsList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := s.CallRaw("Users.getEffectiveUserRights", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList               `json:"errors"`
			Result EffectiveUserRightsList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// UsersSet - Set users' details according given pattern.
// Parameters
//	userIds - list of IDs of users to be changed
//	pattern - pattern to use for new values
// Return
//	errors - create a new user
func (s *ServerConnection) UsersSet(userIds KIdList, pattern User) (ErrorList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
		Pattern User    `json:"pattern"`
	}{userIds, pattern}
	data, err := s.CallRaw("Users.set", params)
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

// UsersWipeMobileDevice - Wipe user's mobile device.
// Parameters
//	userId - global user identifier
//	deviceId - ID of user's mobile device to be wiped
func (s *ServerConnection) UsersWipeMobileDevice(userId KId, deviceId string) error {
	params := struct {
		UserId   KId    `json:"userId"`
		DeviceId string `json:"deviceId"`
	}{userId, deviceId}
	_, err := s.CallRaw("Users.wipeMobileDevice", params)
	return err
}

// UsersGetPersonalContact - Get personal user contacts
func (s *ServerConnection) UsersGetPersonalContact(userIds KIdList) (ErrorList, PersonalContactList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := s.CallRaw("Users.getPersonalContact", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors   ErrorList           `json:"errors"`
			Contacts PersonalContactList `json:"contacts"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Contacts, err
}

// UsersSetPersonalContact - Set personal user contacts
func (s *ServerConnection) UsersSetPersonalContact(userIds KIdList, contact PersonalContact) (ErrorList, error) {
	params := struct {
		UserIds KIdList         `json:"userIds"`
		Contact PersonalContact `json:"contact"`
	}{userIds, contact}
	data, err := s.CallRaw("Users.setPersonalContact", params)
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
