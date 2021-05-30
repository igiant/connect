package connect

import "encoding/json"

// @defgroup SUBGROUP3 Definitions

// IpAddressGroupType -
type IpAddressGroupType string

const (
	Host        IpAddressGroupType = "Host"
	Network     IpAddressGroupType = "Network"
	Range       IpAddressGroupType = "Range"
	ChildGroup  IpAddressGroupType = "ChildGroup"
	ThisMachine IpAddressGroupType = "ThisMachine"
	IpPrefix    IpAddressGroupType = "IpPrefix"
)

type IpAddressGroup struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

// Note: If type is changed, all fields representing associated value must be also assigned, if used in set method.
//    And conversely, type must be assigned if value was changed.

// IpAddressEntry -    type + host + (addr1, addr2) + childGroupId
type IpAddressEntry struct {
	Id             KId                `json:"id"`
	GroupId        KId                `json:"groupId"`
	SharedId       KId                `json:"sharedId"` // read-only; filled when the item is shared in MyKerio
	GroupName      string             `json:"groupName"`
	Description    string             `json:"description"`
	Type           IpAddressGroupType `json:"type"`
	Enabled        bool               `json:"enabled"`
	Status         StoreStatus        `json:"status"`
	Host           string             `json:"host"`  // name, IP or IP prefix
	Addr1          IpAddress          `json:"addr1"` // network/from, e.g. 192.168.0.0
	Addr2          IpAddress          `json:"addr2"` // mask/to, e.g. 255.255.0.0
	ChildGroupId   KId                `json:"childGroupId"`
	ChildGroupName string             `json:"childGroupName"`
}

type IpAddressEntryList []IpAddressEntry

type IpAddressGroupList []IpAddressGroup

// IpAddressGroupsCreate - Create new groups.
// Parameters
//	groups - details for new groups.
// Return
//	errors - possible errors: - "This address group already exists!" duplicate name-value
//	result - list of IDs of created IpAddressGroups
func (c *ServerConnection) IpAddressGroupsCreate(groups IpAddressEntryList) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups IpAddressEntryList `json:"groups"`
	}{groups}
	data, err := c.CallRaw("IpAddressGroups.create", params)
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

// IpAddressGroupsGet - Get the list of IP groups.
// Parameters
//	query - conditions and limits. Included from weblib. KWF engine implementation notes:
//	- LIKE matches substring (second argument) in a string (first argument). There are no wildcards.
//	- sort and match are not case sensitive. - column alias (first operand):
//	- TODO: "QUICKSEARCH" - requested operator applied as following: (name operator secondOperand ) || (description operator secondOperand)
// Return
//	totalItems - count of all groups on the server (before the start/limit applied)
func (c *ServerConnection) IpAddressGroupsGet(query SearchQuery) (IpAddressEntryList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("IpAddressGroups.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       IpAddressEntryList `json:"list"`
			TotalItems int                `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// IpAddressGroupsGetGroupList - Get the list of groups, sorted in ascending order.
// Return
//	groups - list of IP address groups
func (c *ServerConnection) IpAddressGroupsGetGroupList() (IpAddressGroupList, error) {
	data, err := c.CallRaw("IpAddressGroups.getGroupList", nil)
	if err != nil {
		return nil, err
	}
	groups := struct {
		Result struct {
			Groups IpAddressGroupList `json:"groups"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &groups)
	return groups.Result.Groups, err
}

// IpAddressGroupsRemove - Remove groups.
// Parameters
//	groupIds - IDs of groups that should be removed
// Return
//	errors - Errors by removing groups
func (c *ServerConnection) IpAddressGroupsRemove(groupIds KIdList) (ErrorList, error) {
	params := struct {
		GroupIds KIdList `json:"groupIds"`
	}{groupIds}
	data, err := c.CallRaw("IpAddressGroups.remove", params)
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

// IpAddressGroupsSet - Create groups.
// Parameters
//	groupIds - IDs of groups to be updated.
//	details - details for update.
// Return
//	errors - possible errors: - "This address group already exists!" duplicate name-value
func (c *ServerConnection) IpAddressGroupsSet(groupIds KIdList, details IpAddressEntry) (ErrorList, error) {
	params := struct {
		GroupIds KIdList        `json:"groupIds"`
		Details  IpAddressEntry `json:"details"`
	}{groupIds, details}
	data, err := c.CallRaw("IpAddressGroups.set", params)
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

// IpAddressGroupsValidateRemove - Check if groups removal can cut off the administrator from remote administration
// Parameters
//	groupIds - IDs of groups that should be removed
// Return
//	errors - if the result is false, error argument contains additional error info; possible errors:
//	- "You will be cut off from remote administration!"
func (c *ServerConnection) IpAddressGroupsValidateRemove(groupIds KIdList) (ErrorList, error) {
	params := struct {
		GroupIds KIdList `json:"groupIds"`
	}{groupIds}
	data, err := c.CallRaw("IpAddressGroups.validateRemove", params)
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

// IpAddressGroupsValidateSet - Check update of existing groups to see whether this change cut off the administrator from remote administration.
// Parameters
//	groupIds - IDs of groups to be updated.
//	details - details for update.
// Return
//	errors - if the result is false, error argument contains additional error info; possible errors: - "You will be cut off from remote administration!"
func (c *ServerConnection) IpAddressGroupsValidateSet(groupIds KIdList, details IpAddressEntry) (ErrorList, error) {
	params := struct {
		GroupIds KIdList        `json:"groupIds"`
		Details  IpAddressEntry `json:"details"`
	}{groupIds, details}
	data, err := c.CallRaw("IpAddressGroups.validateSet", params)
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
