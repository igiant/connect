package connect

import (
	"encoding/json"
)

type UserDetails struct {
	ID            string `json:"id"`
	DomainID      string `json:"domainId"`
	LoginName     string `json:"loginName"`
	FullName      string `json:"fullName"`
	EffectiveRole struct {
		UserRole           string `json:"userRole"`
		PublicFolderRight  bool   `json:"publicFolderRight"`
		ArchiveFolderRight bool   `json:"archiveFolderRight"`
	} `json:"effectiveRole"`
}

type OutgoingMessageLimit struct {
	IsActive bool `json:"isActive"`
	Limit    struct {
		Value int    `json:"value"`
		Units string `json:"units"`
	} `json:"limit"`
}

type ForwardingOptions struct {
	IsEnabled   bool   `json:"isEnabled"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	How         string `json:"how"`
	PreventLoop bool   `json:"preventLoop"`
}

type KeepForRecovery struct {
	IsEnabled bool `json:"isEnabled"`
	Days      int  `json:"days"`
}

type RenameInfo struct {
	IsRenamed bool   `json:"isRenamed"`
	OldName   string `json:"oldName"`
	NewName   string `json:"newName"`
}

type Domain struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	IsPrimary                 bool   `json:"isPrimary"`
	UserMaxCount              int    `json:"userMaxCount"`
	PasswordExpirationEnabled bool   `json:"passwordExpirationEnabled"`
	PasswordExpirationDays    int    `json:"passwordExpirationDays"`
	PasswordHistoryCount      int    `json:"passwordHistoryCount"`
	PasswordComplexityEnabled bool   `json:"passwordComplexityEnabled"`
	PasswordMinimumLength     int    `json:"passwordMinimumLength"`
	OutgoingMessageLimit      `json:"outgoingMessageLimit"`
	DeletedItems              struct {
		IsEnabled bool `json:"isEnabled"`
		Days      int  `json:"days"`
	} `json:"deletedItems"`
	JunkEmail struct {
		IsEnabled bool `json:"isEnabled"`
		Days      int  `json:"days"`
	} `json:"junkEmail"`
	SentItems struct {
		IsEnabled bool `json:"isEnabled"`
		Days      int  `json:"days"`
	} `json:"sentItems"`
	AutoDelete struct {
		IsEnabled bool `json:"isEnabled"`
		Days      int  `json:"days"`
	} `json:"autoDelete"`
	KeepForRecovery   `json:"keepForRecovery"`
	AliasList         []interface{} `json:"aliasList"`
	ForwardingOptions `json:"forwardingOptions"`
	Service           struct {
		IsEnabled      bool   `json:"isEnabled"`
		ServiceType    string `json:"serviceType"`
		CustomMapFile  string `json:"customMapFile"`
		Authentication struct {
			Username string `json:"username"`
			Password string `json:"password"`
			IsSecure bool   `json:"isSecure"`
		} `json:"authentication"`
		Hostname       string `json:"hostname"`
		BackupHostname string `json:"backupHostname"`
		DirectoryName  string `json:"directoryName"`
		LdapSuffix     string `json:"ldapSuffix"`
	} `json:"service"`
	DomainFooter struct {
		IsUsed         bool   `json:"isUsed"`
		Text           string `json:"text"`
		IsHTML         bool   `json:"isHtml"`
		IsUsedInDomain bool   `json:"isUsedInDomain"`
	} `json:"domainFooter"`
	KerberosRealm string `json:"kerberosRealm"`
	WinNtName     string `json:"winNtName"`
	PamRealm      string `json:"pamRealm"`
	IPAddressBind struct {
		Enabled bool   `json:"enabled"`
		Value   string `json:"value"`
	} `json:"ipAddressBind"`
	Logo struct {
		IsUsed bool   `json:"isUsed"`
		URL    string `json:"url"`
	} `json:"logo"`
	CustomClientLogo struct {
		IsEnabled bool   `json:"isEnabled"`
		URL       string `json:"url"`
		ID        string `json:"id"`
	} `json:"customClientLogo"`
	CheckSpoofedSender bool `json:"checkSpoofedSender"`
	RenameInfo         `json:"renameInfo"`
	DomainQuota        struct {
		DiskSizeLimit struct {
			IsActive bool `json:"isActive"`
			Limit    struct {
				Value int    `json:"value"`
				Units string `json:"units"`
			} `json:"limit"`
		} `json:"diskSizeLimit"`
		ConsumedSize struct {
			Value int    `json:"value"`
			Units string `json:"units"`
		} `json:"consumedSize"`
		Notification struct {
			Type   string `json:"type"`
			Period struct {
				Value int    `json:"value"`
				Units string `json:"units"`
			} `json:"period"`
		} `json:"notification"`
		WarningLimit int    `json:"warningLimit"`
		Email        string `json:"email"`
		Blocks       bool   `json:"blocks"`
	} `json:"domainQuota"`
	IsDistributed             bool   `json:"isDistributed"`
	IsDkimEnabled             bool   `json:"isDkimEnabled"`
	IsLdapManagementAllowed   bool   `json:"isLdapManagementAllowed"`
	IsInstantMessagingEnabled bool   `json:"isInstantMessagingEnabled"`
	UseRemoteArchiveAddress   bool   `json:"useRemoteArchiveAddress"`
	RemoteArchiveAddress      string `json:"remoteArchiveAddress"`
	ArchiveLocalMessages      bool   `json:"archiveLocalMessages"`
	ArchiveIncomingMessages   bool   `json:"archiveIncomingMessages"`
	ArchiveOutgoingMessages   bool   `json:"archiveOutgoingMessages"`
	ArchiveBeforeFilter       bool   `json:"archiveBeforeFilter"`
}

// Login - log in a given user. Please note that with a session to one module you cannot use another one
// (eg. with admin session you cannot use webmail).
// Parameters:
//      user     - login name + domain name (can be omitted if primary) of the user to be logged in,
//                 e.g. "jdoe" or "jdoe@company.com"
//      password - password of the user to be logged in
//      app      - client application description
func (c *Connection) Login(user, password string, app *Application) error {
	if app == nil {
		app = NewApplication("", "", "")
	}
	params := loginStruct{
		user,
		password,
		*app,
	}
	var err error
	data, err := c.CallRaw("Session.login", params)
	token := struct {
		Result struct {
			Token string `json:"token"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return err
	}
	c.Token = &token.Result.Token
	return nil
}

// Logout - Log out the callee
func (c *Connection) Logout() error {
	_, err := c.CallRaw("Session.logout", nil)
	return err
}

// SessionWhoAmI determines the currently logged user (caller, e.g. administrator).
// Fields id and domainId can be empty if built-in administrator is logged-in.
func (c *Connection) SessionWhoAmI() (*UserDetails, error) {
	data, err := c.CallRaw("Session.whoAmI", nil)
	if err != nil {
		return nil, err
	}
	userDetails := struct {
		Result struct {
			UserDetails `json:"userDetails"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &userDetails)
	return &userDetails.Result.UserDetails, err
}

// SessionGetDomain gets domain information of the currently logged user.
// Only name, displayName, ID, description and password policy related fields are filled.
func (c *Connection) SessionGetDomain() (*Domain, error) {
	data, err := c.CallRaw("Session.getDomain", nil)
	if err != nil {
		return nil, err
	}
	domain := struct {
		Result struct {
			Domain `json:"domain"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &domain)
	return &domain.Result.Domain, err
}
