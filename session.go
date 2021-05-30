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

// Login - log in a given user. Please note that with a session to one module you cannot use another one
// (eg. with admin session you cannot use webmail).
// Parameters:
//      user     - login name + domain name (can be omitted if primary) of the user to be logged in,
//                 e.g. "jdoe" or "jdoe@company.com"
//      password - password of the user to be logged in
//      app      - client application description
func (c *ServerConnection) Login(user, password string, app *ApiApplication) error {
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
func (c *ServerConnection) Logout() error {
	_, err := c.CallRaw("Session.logout", nil)
	return err
}

// SessionWhoAmI determines the currently logged user (caller, e.g. administrator).
// Fields id and domainId can be empty if built-in administrator is logged-in.
func (c *ServerConnection) SessionWhoAmI() (*UserDetails, error) {
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
func (c *ServerConnection) SessionGetDomain() (*Domain, error) {
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
