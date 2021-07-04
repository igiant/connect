package connect

import "encoding/json"

// DeliveryType - Delivery Type
type DeliveryType string

const (
	Online           DeliveryType = "Online"           // deliver online, immediatelly
	OfflineScheduler DeliveryType = "OfflineScheduler" // delivery is started by scheduler
	OfflineEtrn      DeliveryType = "OfflineEtrn"      // delivery is started by ETRN command from remote host
)

// Forwarding action
// Forwarding - Note: all fields must be assigned if used in set methods
type Forwarding struct {
	IsEnabled   bool         `json:"isEnabled"`   // is forwarding enabled?
	Host        string       `json:"host"`        // hostname or IP address to forward
	Port        int          `json:"port"`        // host port
	How         DeliveryType `json:"how"`         // how to deliver
	PreventLoop bool         `json:"preventLoop"` // do not deliver to domain alias (applicable when Domain.aliasList is not empty)
}

type DirectoryServiceType string

const (
	WindowsActiveDirectory DirectoryServiceType = "WindowsActiveDirectory" // Windows Active Directory
	AppleDirectoryKerberos DirectoryServiceType = "AppleDirectoryKerberos" // Apple Open Directory with Kerberos authentication
	AppleDirectoryPassword DirectoryServiceType = "AppleDirectoryPassword" // Apple Open Directory with Password Server authentication
	KerioDirectory         DirectoryServiceType = "KerioDirectory"         // Kerio Directory (reserved for future use)
	CustomLDAP             DirectoryServiceType = "CustomLDAP"             // Custom Generic LDAP
)

// DirectoryAuthentication - Note: all fields must be assigned if used in set methods (except password)
type DirectoryAuthentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsSecure bool   `json:"isSecure"` // is used LDAPS?
}

// DirectoryServiceConfiguration - Directory service configuration
type DirectoryServiceConfiguration struct {
	ServiceType    DirectoryServiceType    `json:"serviceType"`    // type of the service
	Authentication DirectoryAuthentication `json:"authentication"` // authentication information
	DirectoryName  string                  `json:"directoryName"`  // Active Directory only: Directory name
	LdapSuffix     string                  `json:"ldapSuffix"`     // Apple Directory, Kerio Directory: LDAP Search Suffix
}

// DirectoryService - Directory service information
type DirectoryService struct {
	IsEnabled      bool                    `json:"isEnabled"`      // directory service is in use / isEnabled must be always assigned if used in set methods
	ServiceType    DirectoryServiceType    `json:"serviceType"`    // type of the service
	CustomMapFile  string                  `json:"customMapFile"`  // Custom Generic LDAP only: custom map filename
	Authentication DirectoryAuthentication `json:"authentication"` // authentication information
	Hostname       string                  `json:"hostname"`       // directory service hostname
	BackupHostname string                  `json:"backupHostname"` // directory service backup hostname
	DirectoryName  string                  `json:"directoryName"`  // Active Directory only: Directory name
	LdapSuffix     string                  `json:"ldapSuffix"`     // Apple Directory, Kerio Directory: LDAP Search Suffix
}

// Footer - Note: all fields must be assigned if used in set methods
type Footer struct {
	IsUsed         bool   `json:"isUsed"`         // is footer used
	Text           string `json:"text"`           // text that will be appended to every message sent from this domain
	IsHtml         bool   `json:"isHtml"`         // if is value false the text is precessed as plaintext
	IsUsedInDomain bool   `json:"isUsedInDomain"` // footer is used also for e-mails within domain
}

// WebmailLogo - Note: all fields must be assigned if used in set methods
type WebmailLogo struct {
	IsUsed bool   `json:"isUsed"` // has domain user defined logo?
	Url    string `json:"url"`    // user defined logo URL
}

type DomainRenameInfo struct {
	IsRenamed bool   `json:"isRenamed"`
	OldName   string `json:"oldName"`
	NewName   string `json:"newName"`
}

type DomainQuota struct {
	DiskSizeLimit SizeLimit          `json:"diskSizeLimit"` // max. disk usage
	ConsumedSize  ByteValueWithUnits `json:"consumedSize"`  // [READ-ONLY] current disk usage
	Notification  QuotaNotification  `json:"notification"`  // option for notification
	WarningLimit  int                `json:"warningLimit"`  // limit in per cent
	Email         string             `json:"email"`         // if quota is exceeded the notification will be sent to this address
	Blocks        bool               `json:"blocks"`        // if reaching the quota will block creation of a new items
}

// Domain - Domain details
type Domain struct {
	Id                        KId              `json:"id"`                        // [READ-ONLY] global identification of domain
	Name                      string           `json:"name"`                      // [REQUIRED FOR CREATE] [WRITE-ONCE] name
	Description               string           `json:"description"`               // description
	IsPrimary                 bool             `json:"isPrimary"`                 // is this domain primary?
	UserMaxCount              int              `json:"userMaxCount"`              // maximum users per domain, 'unlimited' constant can be used
	PasswordExpirationEnabled bool             `json:"passwordExpirationEnabled"` // is password expiration enabled for this domain?
	PasswordExpirationDays    int              `json:"passwordExpirationDays"`    // password expiration interval
	PasswordHistoryCount      int              `json:"passwordHistoryCount"`      // lenght of password history
	PasswordComplexityEnabled bool             `json:"passwordComplexityEnabled"` // is password complexity enabled for this domain?
	PasswordMinimumLength     int              `json:"passwordMinimumLength"`     // minimum password length for complexity feature
	OutgoingMessageLimit      SizeLimit        `json:"outgoingMessageLimit"`      // outgoing message size limit
	DeletedItems              ActionAfterDays  `json:"deletedItems"`              // clean Deleted Items folder (AC maximum: 24855)
	JunkEmail                 ActionAfterDays  `json:"junkEmail"`                 // clean Junk Email folder (AC maximum: 24855)
	SentItems                 ActionAfterDays  `json:"sentItems"`                 // clean Sent Items folder (AC maximum: 24855)
	AutoDelete                ActionAfterDays  `json:"autoDelete"`                // clean all folders (AC minimun:30, maximum: 24855)
	KeepForRecovery           ActionAfterDays  `json:"keepForRecovery"`           // keep deleted messages for recovery
	AliasList                 StringList       `json:"aliasList"`                 // list of domain alternative names
	ForwardingOptions         Forwarding       `json:"forwardingOptions"`         // forwarding settings
	Service                   DirectoryService `json:"service"`                   // directory service configuration
	DomainFooter              Footer           `json:"domainFooter"`              // domain footer setting
	KerberosRealm             string           `json:"kerberosRealm"`             // Kerberos Realm name
	WinNtName                 string           `json:"winNtName"`                 // Windows NT domain name - available on windows only
	PamRealm                  string           `json:"pamRealm"`                  // PAM Realm name - available on linux only
	IpAddressBind             OptionalString   `json:"ipAddressBind"`             // specific IP address bind
	Logo                      WebmailLogo      `json:"logo"`                      // user defined logo
	CustomClientLogo          CustomImage      `json:"customClientLogo"`          // Use custom logo in Kerio Connect Client (if not enabled global option from AdvancedOptionsSetting.webMail is used)
	CheckSpoofedSender        bool             `json:"checkSpoofedSender"`        //
	RenameInfo                DomainRenameInfo `json:"renameInfo"`                // [READ-ONLY] if domain was renamed, contain old and new domain name
	DomainQuota               DomainQuota      `json:"domainQuota"`               // domain's quota settings
	IsDistributed             bool             `json:"isDistributed"`             // [READ-ONLY] if domain is distributed
	IsDkimEnabled             bool             `json:"isDkimEnabled"`             // true if DKIM is used for this domain
	IsLdapManagementAllowed   bool             `json:"isLdapManagementAllowed"`   // [READ-ONLY] true if directory service user/group can be created/deleted
	IsInstantMessagingEnabled bool             `json:"isInstantMessagingEnabled"` // true if Instant Messaging is enabled for this domain
	UseRemoteArchiveAddress   bool             `json:"useRemoteArchiveAddress"`   // if true emails are archived to remoteArchiveAddress
	RemoteArchiveAddress      string           `json:"remoteArchiveAddress"`      // remote archiving address
	ArchiveLocalMessages      bool             `json:"archiveLocalMessages"`      // if true emails from emails are archived to remoteArchiveAddress
	ArchiveIncomingMessages   bool             `json:"archiveIncomingMessages"`   // if true emails are archived to remoteArchiveAddress
	ArchiveOutgoingMessages   bool             `json:"archiveOutgoingMessages"`   // if true emails are archived to remoteArchiveAddress
	ArchiveBeforeFilter       bool             `json:"archiveBeforeFilter"`       // if true emails are archived before content filter check
}

// DomainSetting - Identical settings for all domains
type DomainSetting struct {
	Hostname               string `json:"hostname"`               // internet hostname - how this machine introduces itself in SMTP,POP3...
	PublicFoldersPerDomain bool   `json:"publicFoldersPerDomain"` // true=public folders are unique per each domain / false=global for all domains
	ServerId               KId    `json:"serverId"`               // id of server primary used in cluster
}

// DomainList - List of domains
type DomainList []Domain

// UserLimitType - Types of user amount limit
type UserLimitType string

const (
	DomainLimit  UserLimitType = "DomainLimit"  // stricter limit for amount of users is on domain
	LicenseLimit UserLimitType = "LicenseLimit" // stricter limit for amount of users is on license
)

// MaximumUsers - User limit information
type MaximumUsers struct {
	IsUnlimited  bool          `json:"isUnlimited"`  // is it a special case with no limit for users ?
	AllowedUsers int           `json:"allowedUsers"` // number of allowed users (take minimum of server and domain limit)
	Limit        int           `json:"limit"`        // max. user limit
	LimitType    UserLimitType `json:"limitType"`    // max. user limit type, if domain limit == license limit -> use license
}

// UserDomainCountInfo - User count information
type UserDomainCountInfo struct {
	CurrentUsers int          `json:"currentUsers"` // number of created users on domain
	AllowedUsers MaximumUsers `json:"allowedUsers"` // number of allowed users, take stricter limit from max. number for domain, max. number by license
}

// Domain management

// DomainsCheckPublicFoldersIntegrity - If corrupted folder is found, try to fix it.
func (s *ServerConnection) DomainsCheckPublicFoldersIntegrity() error {
	_, err := s.CallRaw("Domains.checkPublicFoldersIntegrity", nil)
	return err
}

// DomainsCreate - Create new domains.
// Parameters
//	domains - new domain entities
// Return
//	errors - error message list
//	result - particular results for all items
func (s *ServerConnection) DomainsCreate(domains DomainList) (ErrorList, CreateResultList, error) {
	params := struct {
		Domains DomainList `json:"domains"`
	}{domains}
	data, err := s.CallRaw("Domains.create", params)
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

// DomainsGeneratePassword - Generate password which meets current password policy of a given domain.
// Parameters
//	domainId - ID of the domain
// Return
//	password - generated password
func (s *ServerConnection) DomainsGeneratePassword(domainId KId) (string, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("Domains.generatePassword", params)
	if err != nil {
		return "", err
	}
	password := struct {
		Result struct {
			Password string `json:"password"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &password)
	return password.Result.Password, err
}

// DomainsGet - Obtain a list of domains.
// Parameters
//	query - query conditions and limits
// Return
//	list - domains
//  totalItems - amount of domains for given search condition, useful when limit is defined in SearchQuery
func (s *ServerConnection) DomainsGet(query SearchQuery) (DomainList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Domains.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       DomainList `json:"list"`
			TotalItems int        `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DomainsGetDkimDnsRecord - Returns DNS TXT record to be added into DNS.
func (s *ServerConnection) DomainsGetDkimDnsRecord(domain string) (string, error) {
	params := struct {
		Domain string `json:"domain"`
	}{domain}
	data, err := s.CallRaw("Domains.getDkimDnsRecord", params)
	if err != nil {
		return "", err
	}
	detail := struct {
		Result struct {
			Detail string `json:"detail"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &detail)
	return detail.Result.Detail, err
}

// DomainsGetSettings - Get settings common in all domains.
// Return
//	setting - domain global setting
func (s *ServerConnection) DomainsGetSettings() (*DomainSetting, error) {
	data, err := s.CallRaw("Domains.getSettings", nil)
	if err != nil {
		return nil, err
	}
	setting := struct {
		Result struct {
			Setting DomainSetting `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return &setting.Result.Setting, err
}

// DomainsGetUserCountInfo - Get information about user count and limit for domain. Disabled users are not counted.
// Parameters
//	domainId - ID of the domain which will be renamed
// Return
//	countInfo - structure with users count and limit
func (s *ServerConnection) DomainsGetUserCountInfo(domainId KId) (*UserDomainCountInfo, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("Domains.getUserCountInfo", params)
	if err != nil {
		return nil, err
	}
	countInfo := struct {
		Result struct {
			CountInfo UserDomainCountInfo `json:"countInfo"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &countInfo)
	return &countInfo.Result.CountInfo, err
}

// DomainsRemove - Remove domains.
// Parameters
//	domainIds - list of global identifiers of domains to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) DomainsRemove(domainIds KIdList) (ErrorList, error) {
	params := struct {
		DomainIds KIdList `json:"domainIds"`
	}{domainIds}
	data, err := s.CallRaw("Domains.remove", params)
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

// DomainsRename - Start domain renaming process.
// Parameters
//	domainId - ID of the domain which will be renamed
//	newName - new domain name
// Return
//	error - error message
func (s *ServerConnection) DomainsRename(domainId KId, newName string) (*ClusterError, error) {
	params := struct {
		DomainId KId    `json:"domainId"`
		NewName  string `json:"newName"`
	}{domainId, newName}
	data, err := s.CallRaw("Domains.rename", params)
	if err != nil {
		return nil, err
	}
	error := struct {
		Result struct {
			Error ClusterError `json:"error"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &error)
	return &error.Result.Error, err
}

// DomainsSaveFooterImage - Save a new footer's image.
// Parameters
//	fileId - id of uploaded file
// Return
//	imgUrl - url to saved image
func (s *ServerConnection) DomainsSaveFooterImage(fileId string) (string, error) {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	data, err := s.CallRaw("Domains.saveFooterImage", params)
	if err != nil {
		return "", err
	}
	imgUrl := struct {
		Result struct {
			ImgUrl string `json:"imgUrl"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &imgUrl)
	return imgUrl.Result.ImgUrl, err
}

// DomainsSaveWebMailLogo - Save a new logo.
// Parameters
//	fileId - ID of the uploaded file
//	domainId - global domain identifier
// Return
//	logoUrl - path to the saved file
func (s *ServerConnection) DomainsSaveWebMailLogo(fileId string, domainId KId) (string, error) {
	params := struct {
		FileId   string `json:"fileId"`
		DomainId KId    `json:"domainId"`
	}{fileId, domainId}
	data, err := s.CallRaw("Domains.saveWebMailLogo", params)
	if err != nil {
		return "", err
	}
	logoUrl := struct {
		Result struct {
			LogoUrl string `json:"logoUrl"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &logoUrl)
	return logoUrl.Result.LogoUrl, err
}

// DomainsSet - Set existing domains to given pattern.
// Parameters
//	domainIds - list of the domain's global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) DomainsSet(domainIds KIdList, pattern Domain) (ErrorList, error) {
	params := struct {
		DomainIds KIdList `json:"domainIds"`
		Pattern   Domain  `json:"pattern"`
	}{domainIds, pattern}
	data, err := s.CallRaw("Domains.set", params)
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

// DomainsSetSettings - Set settings for all domains.
// Parameters
//	setting - domain global settings
func (s *ServerConnection) DomainsSetSettings(setting DomainSetting) error {
	params := struct {
		Setting DomainSetting `json:"setting"`
	}{setting}
	_, err := s.CallRaw("Domains.setSettings", params)
	return err
}

// DomainsTestDomainController - Test connection between Kerio Connect and domain controller.
// Parameters
//	hostnames - directory server (primary and secondary if any)
//	config - directory service configuration. If password is empty then it is taken from domain by 'domainId'.
//	domainId - global domain identifier
// Return
//	errors - error message
func (s *ServerConnection) DomainsTestDomainController(hostnames StringList, config DirectoryServiceConfiguration, domainId KId) (ErrorList, error) {
	params := struct {
		Hostnames StringList                    `json:"hostnames"`
		Config    DirectoryServiceConfiguration `json:"config"`
		DomainId  KId                           `json:"domainId"`
	}{hostnames, config, domainId}
	data, err := s.CallRaw("Domains.testDomainController", params)
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

// DomainsTestDkimDnsStatus - Tests DKIM DNS TXT status for domain list.
// Parameters
//	hostnames - hostnames checked for DKIM public key in DNS
// Return
//	errors - error message
func (s *ServerConnection) DomainsTestDkimDnsStatus(hostnames StringList) (ErrorList, error) {
	params := struct {
		Hostnames StringList `json:"hostnames"`
	}{hostnames}
	data, err := s.CallRaw("Domains.testDkimDnsStatus", params)
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

// DomainsGetDomainFooterPlaceholders - Return all supported placeholders for domain footer
func (s *ServerConnection) DomainsGetDomainFooterPlaceholders() (NamedConstantList, error) {
	data, err := s.CallRaw("Domains.getDomainFooterPlaceholders", nil)
	if err != nil {
		return nil, err
	}
	placeholders := struct {
		Result struct {
			Placeholders NamedConstantList `json:"placeholders"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &placeholders)
	return placeholders.Result.Placeholders, err
}
