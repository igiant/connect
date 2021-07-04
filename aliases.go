package connect

import "encoding/json"

// AliasType - Alias type definition
type AliasType string

const (
	TypePublicFolder AliasType = "TypePublicFolder" // messages are delivered to public folder
	TypeEmailAddress AliasType = "TypeEmailAddress" // messages are delivered to email account
)

// Alias - Alias details
type Alias struct {
	Id          KId        `json:"id"`          // global identification of alias
	DomainId    KId        `json:"domainId"`    // [REQUIRED FOR CREATE] identification in which domain alias exists
	Name        string     `json:"name"`        // [REQUIRED FOR CREATE] [USED BY QUICKSEARCH] left side of alias
	DeliverToId KId        `json:"deliverToId"` // empty if email or contains public folder kid
	DeliverTo   string     `json:"deliverTo"`   // [REQUIRED FOR CREATE] [USED BY QUICKSEARCH] email address or public folder name
	Type        AliasType  `json:"type"`        // type of the alias
	Description string     `json:"description"` // description
	HomeServer  HomeServer `json:"homeServer"`  // [READ-ONLY] Id of alias homeserver if server is in Cluster
}

// AliasList - List of aliases
type AliasList []Alias

// AliasTargetType - Alias Target discriminator
type AliasTargetType string

const (
	TypeUser  AliasTargetType = "TypeUser"  // user
	TypeGroup AliasTargetType = "TypeGroup" // group
)

// AliasTarget - Alias target can be a user or group
type AliasTarget struct {
	Id           KId             `json:"id"`           // unique identifier
	Type         AliasTargetType `json:"type"`         // item type discriminator
	Name         string          `json:"name"`         // loginName for the User, name in square brackets for the Group
	FullName     string          `json:"fullName"`     // fullname for the User, empty string for the Group
	Description  string          `json:"description"`  // description of User/Group
	IsEnabled    bool            `json:"isEnabled"`    // is the User/Group enabled?
	ItemSource   DataSource      `json:"itemSource"`   // is the User/Group stored internally or by LDAP?
	EmailAddress string          `json:"emailAddress"` // first email address
	HomeServer   HomeServer      `json:"homeServer"`   // id of users homeserver if server is in Cluster; groups haven't homeserver
}

// AliasTargetList - List of alias targets
type AliasTargetList []AliasTarget

// Alias management

// AliasesCheck - Obtain a list of mail addresses and/or public folders on which given string will be expanded.
//	checkString - string to be checked
// Return
//	result - list of expansions
func (s *ServerConnection) AliasesCheck(checkString string) (StringList, error) {
	params := struct {
		CheckString string `json:"checkString"`
	}{checkString}
	data, err := s.CallRaw("Aliases.check", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result StringList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Result.Result, err
}

// AliasesCreate - Create new aliases
//	aliases - new alias entities
// Return
//	errors - list of error messages for appropriate new aliases
//	result - list of IDs of created aliases
func (s *ServerConnection) AliasesCreate(aliases AliasList) (ErrorList, CreateResultList, error) {
	params := struct {
		Aliases AliasList `json:"aliases"`
	}{aliases}
	data, err := s.CallRaw("Aliases.create", params)
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

// AliasesGet - Obtain list of aliases.
//	query - query conditions and limits
// Return
//	list - aliases
//  totalItems - amount of aliases for given search condition, useful when limit is defined in query
func (s *ServerConnection) AliasesGet(query SearchQuery, domainId KId) (AliasList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Aliases.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       AliasList `json:"list"`
			TotalItems int       `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// AliasesGetMailPublicFolderList - Obtain a list of mail public folders in the given domain.
//	domainId - global identification of the domain
// Return
//	publicFolderList - list of public folders
func (s *ServerConnection) AliasesGetMailPublicFolderList(domainId KId) (PublicFolderList, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("Aliases.getMailPublicFolderList", params)
	if err != nil {
		return nil, err
	}
	publicFolderList := struct {
		Result struct {
			PublicFolderList PublicFolderList `json:"publicFolderList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &publicFolderList)
	return publicFolderList.Result.PublicFolderList, err
}

// AliasesGetTargetList - Obtain a list of alias targets.
//	query - query conditions and limits
//	domainId - global identification of the domain
// Return
//	list - alias targets
//  totalItems - amount of aliases for given search condition, useful when a limit is defined in the query
func (s *ServerConnection) AliasesGetTargetList(query SearchQuery, domainId KId) (AliasTargetList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Aliases.getTargetList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       AliasTargetList `json:"list"`
			TotalItems int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// AliasesRemove - Delete aliases.
// Return
//	errors - error message list
func (s *ServerConnection) AliasesRemove(aliasIds KIdList) (ErrorList, error) {
	params := struct {
		AliasIds KIdList `json:"aliasIds"`
	}{aliasIds}
	data, err := s.CallRaw("Aliases.remove", params)
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

// AliasesSet - Set an existing alias.
//	aliasIds - list of alias global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) AliasesSet(aliasIds KIdList, pattern Alias) (ErrorList, error) {
	params := struct {
		AliasIds KIdList `json:"aliasIds"`
		Pattern  Alias   `json:"pattern"`
	}{aliasIds, pattern}
	data, err := s.CallRaw("Aliases.set", params)
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
