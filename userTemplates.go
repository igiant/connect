package connect

import "encoding/json"

// ValidFor - User Template Scope
type ValidFor string

const (
	OneDomain  ValidFor = "OneDomain"
	AllDomains ValidFor = "AllDomains"
)

// UserTemplate - Details of user template - meaning is the same as in structure User
type UserTemplate struct {
	Id                   KId                  `json:"id"`
	Name                 string               `json:"name"`                 // [REQUIRED FOR CREATE] [USED BY QUICKSEARCH] name of template (displayed in list of templates)
	Description          string               `json:"description"`          // [USED BY QUICKSEARCH] description of template (displayed after its selection)
	AuthType             UserAuthType         `json:"authType"`             // supported values must be retrieved from engine by ServerInfo::getSupportedAuthTypes()
	IsPasswordReversible bool                 `json:"isPasswordReversible"` // typically SHA1
	HasDefaultSpamRule   bool                 `json:"hasDefaultSpamRule"`   // should be spam rule enabled?
	Role                 UserRight            `json:"role"`                 // list of user roles (excluding public/archive folder rights)
	Scope                ValidFor             `json:"scope"`                // scope of template
	DomainId             KId                  `json:"domainId"`             // not relevant for templating, only for filter (condition)
	EmailAddresses       UserEmailAddressList `json:"emailAddresses"`       // filled only if domain is set
	UserGroups           UserGroupList        `json:"userGroups"`           // filled only if domain is set
	EmailForwarding      EmailForwarding      `json:"emailForwarding"`      // email forwarding setting
	ItemLimit            ItemCountLimit       `json:"itemLimit"`            // max. number of items
	DiskSizeLimit        SizeLimit            `json:"diskSizeLimit"`        // max. disk usage
	HasDomainRestriction bool                 `json:"hasDomainRestriction"` // user can send/receive from/to his domain only
	OutMessageLimit      SizeLimit            `json:"outMessageLimit"`      // limit of outgoing message
	PublishInGal         bool                 `json:"publishInGal"`         // publish user in global address list
	CleanOutItems        CleanOut             `json:"cleanOutItems"`        // Items clean-out settings
	AllowPasswordChange  bool                 `json:"allowPasswordChange"`  // if it is set to false the password can be changed only by the administrator
	AccessPolicy         IdEntity             `json:"accessPolicy"`         // ID and name of Access Policy applied for user. Only ID is writable.
	CompanyContactId     KId                  `json:"companyContactId"`     // ID of company contact associated with this template
	HomeServerId         KId                  `json:"homeServerId"`         // ID of distributed domain home server guid associated with this template
}

type UserTemplateList []UserTemplate

// UserTemplatesCreate - Create user templates.
// Parameters
//	userTemplates - new user template entities
// Return
//	errors - error message list
//	result - list of IDs of created templates
func (c *ServerConnection) UserTemplatesCreate(userTemplates UserTemplateList) (ErrorList, CreateResultList, error) {
	params := struct {
		UserTemplates UserTemplateList `json:"userTemplates"`
	}{userTemplates}
	data, err := c.CallRaw("UserTemplates.create", params)
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

// UserTemplatesGet - Obtain a list of user templates.
// Parameters
//	query - query attributes and limits
// Return
//	userTemplateList - list of user templates
//	totalItems - number of all returned templates
func (c *ServerConnection) UserTemplatesGet(query SearchQuery) (UserTemplateList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("UserTemplates.get", params)
	if err != nil {
		return nil, 0, err
	}
	userTemplateList := struct {
		Result struct {
			UserTemplateList UserTemplateList `json:"userTemplateList"`
			TotalItems       int              `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &userTemplateList)
	return userTemplateList.Result.UserTemplateList, userTemplateList.Result.TotalItems, err
}

// UserTemplatesGetAvailable - - Only templates without administrative rights are listed.
// Parameters
//	domainId - only templates with this domain and templates without domain are listed
// Return
//	userTemplateList - list of user templates
func (c *ServerConnection) UserTemplatesGetAvailable(domainId KId) (UserTemplateList, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := c.CallRaw("UserTemplates.getAvailable", params)
	if err != nil {
		return nil, err
	}
	userTemplateList := struct {
		Result struct {
			UserTemplateList UserTemplateList `json:"userTemplateList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &userTemplateList)
	return userTemplateList.Result.UserTemplateList, err
}

// UserTemplatesRemove - Remove list of user template records.
// Parameters
//	idList - list of identifiers of deleted user templates
// Return
//	errors - error message list
func (c *ServerConnection) UserTemplatesRemove(idList KIdList) (ErrorList, error) {
	params := struct {
		IdList KIdList `json:"idList"`
	}{idList}
	data, err := c.CallRaw("UserTemplates.remove", params)
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

// UserTemplatesSet - Set user templates according a given pattern.
// Parameters
//	idList - list of domain global identifier(s) of items to be changed
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (c *ServerConnection) UserTemplatesSet(idList KIdList, pattern UserTemplate) (ErrorList, error) {
	params := struct {
		IdList  KIdList      `json:"idList"`
		Pattern UserTemplate `json:"pattern"`
	}{idList, pattern}
	data, err := c.CallRaw("UserTemplates.set", params)
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
