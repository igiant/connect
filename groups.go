package connect

import "encoding/json"

// Group - Group details
type Group struct {
	Id                   KId          `json:"id"`                   // global identification of group
	DomainId             KId          `json:"domainId"`             // identification in which domain group exists
	Name                 string       `json:"name"`                 // group name
	EmailAddresses       StringList   `json:"emailAddresses"`       // group email addresses
	Role                 UserRoleType `json:"role"`                 // access rights list
	HasDomainRestriction bool         `json:"hasDomainRestriction"` // user can send/receive from/to his domain only
	Description          string       `json:"description"`          // description
	ItemSource           DataSource   `json:"itemSource"`           // internal database or mapped from LDAP?
	PublishInGal         bool         `json:"publishInGal"`         // publish user in global address list? - default is true
}

// GroupList - List of groups
type GroupList []Group

type GroupRemovalRequest struct {
	GroupId KId                        `json:"groupId"` // ID of group to be removed
	Mode    DirectoryServiceDeleteMode `json:"mode"`    // delete mode
}

type GroupRemovalRequestList []GroupRemovalRequest

// Group management

// GroupsActivate - Activate groups from a directory service
// Return
//	errors - list of error messages for appropriate groups
func (s *ServerConnection) GroupsActivate(groupIdList KIdList) (ErrorList, error) {
	params := struct {
		GroupIdList KIdList `json:"groupIdList"`
	}{groupIdList}
	data, err := s.CallRaw("Groups.activate", params)
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

// GroupsAddMemberList - Add new member(s) to a group.
// Parameters
//	groupId - global group identifier
//	userList - list of global identifiers of users to be added to a group
// Return
//	errors - error message list
func (s *ServerConnection) GroupsAddMemberList(groupId KId, userList KIdList) (ErrorList, error) {
	params := struct {
		GroupId  KId     `json:"groupId"`
		UserList KIdList `json:"userList"`
	}{groupId, userList}
	data, err := s.CallRaw("Groups.addMemberList", params)
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

// GroupsCreate - Create new groups.
// Parameters
//	groups - new group entities
// Return
//	errors - error message list
//	result - list of IDs of created groups
func (s *ServerConnection) GroupsCreate(groups GroupList) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups GroupList `json:"groups"`
	}{groups}
	data, err := s.CallRaw("Groups.create", params)
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

// GroupsCreateLdap - Create new groups in directory service.
// Parameters
//	groups - new group entities
// Return
//	errors - error message list
//	result - list of IDs of created groups
func (s *ServerConnection) GroupsCreateLdap(groups GroupList) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups GroupList `json:"groups"`
	}{groups}
	data, err := s.CallRaw("Groups.createLdap", params)
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

// GroupsGet - Obtain a list of groups.
// Parameters
//	query - query conditions and limits
// Return
//	list - groups
//  totalItems - amount of groups for given search condition, useful when limit is defined in SearchQuery
func (s *ServerConnection) GroupsGet(query SearchQuery, domainId KId) (GroupList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Groups.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       GroupList `json:"list"`
			TotalItems int       `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// GroupsRemove - Note: it is not necessary to remove members before deleting a group
// Return
//	errors - error message list
func (s *ServerConnection) GroupsRemove(requests GroupRemovalRequestList) (ErrorList, error) {
	params := struct {
		Requests GroupRemovalRequestList `json:"requests"`
	}{requests}
	data, err := s.CallRaw("Groups.remove", params)
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

// GroupsRemoveMemberList - Remove member(s) from a group.
// Parameters
//	groupId - global group identifier
//	userIds - list of global identifiers of users to be add to a group
// Return
//	errors - error message list
func (s *ServerConnection) GroupsRemoveMemberList(groupId KId, userIds KIdList) (ErrorList, error) {
	params := struct {
		GroupId KId     `json:"groupId"`
		UserIds KIdList `json:"userIds"`
	}{groupId, userIds}
	data, err := s.CallRaw("Groups.removeMemberList", params)
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

// GroupsSet - Create a new group.
// Parameters
//	groupIds - a list group global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) GroupsSet(groupIds KIdList, pattern Group) (ErrorList, error) {
	params := struct {
		GroupIds KIdList `json:"groupIds"`
		Pattern  Group   `json:"pattern"`
	}{groupIds, pattern}
	data, err := s.CallRaw("Groups.set", params)
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
