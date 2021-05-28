package connect

type PublicFolder struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

type PublicFolderList []PublicFolder

// List of email addresses
type UserEmailAddressList []string

// Export format type.
type FileFormatType string

const (
	TypeXml FileFormatType = "TypeXml" // Extensible Markup Language
	TypeCsv FileFormatType = "TypeCsv" // Comma Separated Values
)

// Type of user role.
type UserRoleType string

const (
	UserRole           UserRoleType = "UserRole"           // regular user without any administration rights
	Auditor            UserRoleType = "Auditor"            // read only access to administration
	AccountAdmin       UserRoleType = "AccountAdmin"       // can administer Users,Groups,Aliases,MLs
	FullAdmin          UserRoleType = "FullAdmin"          // unlimited administration
	BuiltInAdmin       UserRoleType = "BuiltInAdmin"       // BuiltIn admin role can be returned only in Session::WhoAmI method for built-in administrator. This role must NOT be assigned.
	BuiltInDomainAdmin UserRoleType = "BuiltInDomainAdmin" // BuiltIn domain admin role can be returned only in Session::WhoAmI method for built-in domain administrator. This role must NOT be assigned.
)

//
// Note: all fields must be assigned if used in set methods.
type UserRight struct {
	UserRole           UserRoleType `json:"userRole"`
	PublicFolderRight  bool         `json:"publicFolderRight"`
	ArchiveFolderRight bool         `json:"archiveFolderRight"`
}

// Forwarding setup for user.
type UserForwardMode string

const (
	UForwardNone    UserForwardMode = "UForwardNone"    // Forwarding is disabled
	UForwardYes     UserForwardMode = "UForwardYes"     // Forward all messages for this user to some addresses, don't deliver the message to the mailbox.
	UForwardDeliver UserForwardMode = "UForwardDeliver" // Forward all messages for this user to some addresses, and also deliver the message to user's mailbox.
)

// Type of deleting folder of the user
type UserDeleteFolderMode string

const (
	UDeleteUser   UserDeleteFolderMode = "UDeleteUser"   // Delete user without deleting his folder.
	UDeleteFolder UserDeleteFolderMode = "UDeleteFolder" // Delete user and delete his folder.
	UMoveFolder   UserDeleteFolderMode = "UMoveFolder"   // Delete user and his folder will move into another user's folder.
)

// Settings of email forwarding.
// Note: all fields must be assigned if used in set methods.
type EmailForwarding struct {
	Mode           UserForwardMode      `json:"mode"`
	EmailAddresses UserEmailAddressList `json:"emailAddresses"` // list of email addresses, make sense only for UForwardDeliver
}

// Properties of user's groups.
type UserGroup struct {
	Id          KId        `json:"id"` // global identification
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ItemSource  DataSource `json:"itemSource"`
}

// List of user's groups.
type UserGroupList []UserGroup

// Settings of items limit.
// Note: all fields must be assigned if used in set methods.
type ItemCountLimit struct {
	IsActive bool `json:"isActive"`
	Limit    int  `json:"limit"`
}

// Amount of storage used and items currently stored in user's store.
type QuotaUsage struct {
	Items   int                `json:"items"`
	Storage ByteValueWithUnits `json:"storage"`
}

// List of QuotaUsage.
type QuotaUsageList []QuotaUsage

// Last login information.
type LastLogin struct {
	DateTime DateTimeStamp `json:"dateTime"` // date and time of last login
	Protocol string        `json:"protocol"` // protocol name of last login, example POP3
}

// Per-user message retention policy.
type CleanOut struct {
	IsUsedDomain bool            `json:"isUsedDomain"` // use domain settings
	DeletedItems ActionAfterDays `json:"deletedItems"` // clean Deleted Items folder (maximum: 24855)
	JunkEmail    ActionAfterDays `json:"junkEmail"`    // clean Junk Email folder (maximum: 24855)
	SentItems    ActionAfterDays `json:"sentItems"`    // clean Sent Items folder (maximum: 24855)
	AutoDelete   ActionAfterDays `json:"autoDelete"`   // clean all folders (maximum: 24855)
}

// User details.
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
	IsWritableByMe       bool                 `json:"isWritableByMe"`       // Does caller have right to change the user? E.g. if Account Admin gets User structure for Full Admin, isWritableByMe will be false. This field is read-only and cannot be used in kerio::web::SearchQuery conditions.
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

// List of users.
type UserList []User

// User effective rights (inherited from groups)
type EffectiveUserRights struct {
	UserId               KId  `json:"userId"`               // [READ-ONLY] global identification
	HasDomainRestriction bool `json:"hasDomainRestriction"` // user can send/receive from/to his/her domain only
}

// List of users effective rights
type EffectiveUserRightsList []EffectiveUserRights

// Type of user directory
type ServerDirectoryType string

const (
	WinNT            ServerDirectoryType = "WinNT"            // Windows NT Domain directory (Win NT 4.0)
	ActiveDirectory  ServerDirectoryType = "ActiveDirectory"  // Active Directory (Windows 2000 and newer)
	NovellEDirectory ServerDirectoryType = "NovellEDirectory" // Novell eDirectory
)

// Properties of the server from which users are imported.
type ImportServer struct {
	DirectoryType      ServerDirectoryType `json:"directoryType"`
	RemoteDomainName   string              `json:"remoteDomainName"`
	Address            string              `json:"address"` // server IP or FQDN
	LoginName          string              `json:"loginName"`
	Password           string              `json:"password"`
	LdapFilter         string              `json:"ldapFilter"`
	IsSecureConnection bool                `json:"isSecureConnection"`
}

// Login statistics - count and timestamp of the last login.
type LoginStats struct {
	Count     int    `json:"count"`
	LastLogin string `json:"lastLogin"`
}

// Statistics about user's usage of quota, logins to different services.
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

// List of users' statistics.
type UserStatList []UserStats

// Result of a mass operation.
type ResultTriplet struct {
	InputIndex int `json:"inputIndex"`
	ItemsCount int `json:"itemsCount"`
}

// List of mass operation results.
type ResultTripletList []ResultTriplet

// User to be removed, what to do with his/her mailbox.
type RemovalRequest struct {
	UserId           KId                        `json:"userId"`           // ID of user to be removed
	Method           UserDeleteFolderMode       `json:"method"`           // removal method
	RemoveReferences bool                       `json:"removeReferences"` // if true all reference to this user is going to be removed as well
	TargetUserId     KId                        `json:"targetUserId"`     // applicable only when moving user's store to another user, use empty string if not moving user's messages to target mailbox
	Mode             DirectoryServiceDeleteMode `json:"mode"`             // delete mode
}

type RemovalRequestList []RemovalRequest

// A user being imported from directory server.
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

// Resut of autentication.
type AuthResult string

const (
	AuthOK           AuthResult = "AuthOK"           // User was autenticated
	AuthFail         AuthResult = "AuthFail"         // Wrong login name or password.
	AuthUserDisabled AuthResult = "AuthUserDisabled" // User cannot to log in, because his account is disabled.
	AuthLicense      AuthResult = "AuthLicense"      // User cannot log in, because license limit was reached.
	AuthDenied       AuthResult = "AuthDenied"       // User is denied to log in.
	AuthTryLater     AuthResult = "AuthTryLater"     // User cannot to log in at this moment, try later.
)
