package connect

import "encoding/json"

type Alias struct {
	ID          KId        `json:"id"`
	DomainID    KId        `json:"domainId"`
	Name        string     `json:"name"`
	DeliverToID KId        `json:"deliverToId"`
	DeliverTo   string     `json:"deliverTo"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	HomeServer  HomeServer `json:"homeServer"`
}

type AliasList []Alias

// HomeServer User's home server in a distributed domain.
type HomeServer struct {
	ID   KId    `json:"id"`   // server's id
	Name string `json:"name"` // server's Internet hostname
}

type AliasTarget struct {
	ID           string     `json:"id"`           // unique identifier
	Type         string     `json:"type"`         // item type discriminator
	Name         string     `json:"name"`         // loginName for the User, name in square brackets for the Group
	FullName     string     `json:"fullName"`     // fullname for the User, empty string for the Group
	Description  string     `json:"description"`  // description of User/Group
	IsEnabled    string     `json:"isEnabled"`    // is the User/Group enabled?
	ItemSource   string     `json:"itemSource"`   // is the User/Group stored internally or by LDAP?
	EmailAddress string     `json:"emailAddress"` // first email address
	HomeServer   HomeServer `json:"homeServer"`   // id of users homeserver if server is in Cluster; groups haven't homeserver
}

type AliasTargetList []AliasTarget

// AliasesCheck obtains a list of mail addresses and/or public folders on which given string will be expanded.
// Parameters
//      checkString	- string to be checked
// Return
//      list of expansions
func (c *Connection) AliasesCheck(checkString string) ([]string, error) {
	params := struct {
		CheckString string `json:"checkString"`
	}{checkString}
	data, err := c.CallRaw("Aliases.check", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result []string `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Result.Result, err
}

// AliasesCreate - create new aliases.
// Parameters
//      aliases	- new alias entities
// Return
//  list of IDs of created aliases
//  list of error messages for appropriate new aliases
func (c *Connection) AliasesCreate(aliases AliasList) (CreateResultList, ErrorList, error) {
	params := struct {
		Aliases AliasList `json:"aliases"`
	}{aliases}
	data, err := c.CallRaw("Aliases.create", params)
	if err != nil {
		return nil, nil, err
	}
	result := struct {
		Result struct {
			Result CreateResultList `json:"result"`
			Errors ErrorList        `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Result.Result, result.Result.Errors, err
}

// AliasesGet obtains a list of aliases.
// Parameters
//      query	    - query conditions and limits
//      domainKId	- domain identification
// Return
//      aliases
func (c *Connection) AliasesGet(domainID string, query SearchQuery) (AliasList, error) {
	params := struct {
		DomainID string      `json:"domainId"`
		Query    SearchQuery `json:"query"`
	}{domainID, query}
	data, err := c.CallRaw("Aliases.get", params)
	if err != nil {
		return nil, err
	}
	aliasList := struct {
		Result struct {
			AliasList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &aliasList)
	return aliasList.Result.AliasList, err
}

// AliasesGetMailPublicFolderList obtains a list of mail public folders in the given domain.
// Parameters
//      domainId	- global identification of the domain
// Return
//      list of public folders
func (c *Connection) AliasesGetMailPublicFolderList(domainID string) (PublicFolderList, error) {
	params := struct {
		DomainID string `json:"domainId"`
	}{domainID}
	data, err := c.CallRaw("Aliases.getMailPublicFolderList", params)
	if err != nil {
		return nil, err
	}
	publicFolderList := struct {
		Result struct {
			PublicFolderList `json:"publicFolderList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &publicFolderList)
	return publicFolderList.Result.PublicFolderList, err
}

// AliasesGetTargetList obtains a list of alias targets.
// Parameters
//      query	    - query conditions and limits
//      domainId	- global identification of the domain
// Return
//      alias targets
func (c *Connection) AliasesGetTargetList(domainID string, query SearchQuery) (AliasTargetList, error) {
	params := struct {
		DomainID string      `json:"domainId"`
		Query    SearchQuery `json:"query"`
	}{domainID, query}
	data, err := c.CallRaw("Aliases.getTargetList", params)
	if err != nil {
		return nil, err
	}
	aliasTargetList := struct {
		Result struct {
			AliasTargetList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &aliasTargetList)
	return aliasTargetList.Result.AliasTargetList, err
}

// AliasesRemove - Delete aliases.
// Parameters
//      aliasList	- list of global identifiers of aliases to be deleted
// Return
//      error message list
func (c *Connection) AliasesRemove(aliasIDs []string) (ErrorList, error) {
	params := struct {
		AliasIDs []string `json:"aliasIds"`
	}{aliasIDs}
	data, err := c.CallRaw("Aliases.remove", params)
	if err != nil {
		return nil, err
	}
	errorList := struct {
		Result struct {
			ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errorList)
	return errorList.Result.ErrorList, err
}

// AliasesSet - Set an existing alias.
// Parameters
//      aliasList   - list of alias global identifier(s)
//      pattern     - pattern to use for new values
// Return
//      error message list
func (c *Connection) AliasesSet(aliasIDs []string, pattern Alias) (ErrorList, error) {
	params := struct {
		AliasIDs []string `json:"aliasIds"`
		Pattern  Alias
	}{aliasIDs, pattern}
	data, err := c.CallRaw("Aliases.set", params)
	if err != nil {
		return nil, err
	}
	errorList := struct {
		Result struct {
			ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errorList)
	return errorList.Result.ErrorList, err
}
