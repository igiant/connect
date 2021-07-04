package connect

import "encoding/json"

// PrincipalType - Principal type
type PrincipalType string

const (
	AuthDomainPrincipal PrincipalType = "AuthDomainPrincipal" // represents all users from a specified domain
	AuthUserPrincipal   PrincipalType = "AuthUserPrincipal"   // represents all users from all domains
	AnyonePrincipal     PrincipalType = "AnyonePrincipal"     // represents authorized users from all domains
	UserPrincipal       PrincipalType = "UserPrincipal"       // represents an user
	GroupPrincipal      PrincipalType = "GroupPrincipal"      // represents an group
	NonePrincipal       PrincipalType = "NonePrincipal"       // target is empty
)

// PrincipalDescription - Principal descriptor [READ-ONLY]
type PrincipalDescription struct {
	Type        PrincipalType `json:"type"`
	Id          KId           `json:"id"`
	Name        string        `json:"name"`
	DomainName  string        `json:"domainName"`
	FullName    string        `json:"fullName"`
	Description string        `json:"description"`
	IsEnabled   bool          `json:"isEnabled"`
	ItemSource  DataSource    `json:"itemSource"` // internal/LDAP
	HomeServer  HomeServer    `json:"homeServer"` // id of users homeserver if server is in Cluster; groups haven't homeserver
}

// PrincipalList - List of principals
type PrincipalList []PrincipalDescription

// ResourceType - Export format type
type ResourceType string

const (
	Room      ResourceType = "Room"      // resource is a room
	Equipment ResourceType = "Equipment" // resource is something else, eg: a car
)

// Resource - Resource details
type Resource struct {
	Id            KId                  `json:"id"`            // [READ-ONLY] global identification of resource
	DomainId      KId                  `json:"domainId"`      // [REQUIRED FOR CREATE] [WRITE-ONCE] identification in which domain resource exists
	Name          string               `json:"name"`          // [REQUIRED FOR CREATE] [WRITE-ONCE] resource name
	Address       string               `json:"address"`       // [READ-ONLY] email of resource
	Description   string               `json:"description"`   // resource description
	Type          ResourceType         `json:"type"`          // type of the resource
	IsEnabled     bool                 `json:"isEnabled"`     // is resource enabled? default == true
	ResourceUsers PrincipalList        `json:"resourceUsers"` // list of groups / users /
	Manager       PrincipalDescription `json:"manager"`       // identification of user who is a manager
	HomeServer    HomeServer           `json:"homeServer"`    // [READ-ONLY] Id of users homeserver if server is in Cluster
}

// ResourceList - List of resources
type ResourceList []Resource

// Resource management

// ResourcesCreate - Create new resources.
// Parameters
//	resources - new resource entities
// Return
//	errors - error message list
//	result - list of IDs of created resources
func (s *ServerConnection) ResourcesCreate(resources ResourceList) (ErrorList, CreateResultList, error) {
	params := struct {
		Resources ResourceList `json:"resources"`
	}{resources}
	data, err := s.CallRaw("Resources.create", params)
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

// ResourcesGet - Obtain a list of resources.
// Parameters
//	query - query conditions and limits
//	domainId - domain identification
// Return
//	list - resources
//  totalItems - amount of resources for given search condition, useful when limit is defined in SearchQuery
func (s *ServerConnection) ResourcesGet(query SearchQuery, domainId KId) (ResourceList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Resources.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ResourceList `json:"list"`
			TotalItems int          `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ResourcesGetPrincipalList - Obtain a list of potential resource targets (principals).
// Parameters
//	query - query attributes and limits
// Return
//	list - principals
//  totalItems - amount of resources for given search condition, useful when limit is defined in SearchQuery
func (s *ServerConnection) ResourcesGetPrincipalList(query SearchQuery, domainId KId) (PrincipalList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Resources.getPrincipalList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       PrincipalList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ResourcesRemove - Remove resources.
// Parameters
//	resourceIds - list of global identifiers of resource(s) to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) ResourcesRemove(resourceIds KIdList) (ErrorList, error) {
	params := struct {
		ResourceIds KIdList `json:"resourceIds"`
	}{resourceIds}
	data, err := s.CallRaw("Resources.remove", params)
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

// ResourcesSet - Set existing resources.
// Parameters
//	resourceIds - a list resource global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) ResourcesSet(resourceIds KIdList, pattern Resource) (ErrorList, error) {
	params := struct {
		ResourceIds KIdList  `json:"resourceIds"`
		Pattern     Resource `json:"pattern"`
	}{resourceIds, pattern}
	data, err := s.CallRaw("Resources.set", params)
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
