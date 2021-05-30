package connect

import "encoding/json"

type ServiceType string

const (
	ServiceActiveSync ServiceType = "ServiceActiveSync" // ActiveSync
	ServiceEWS        ServiceType = "ServiceEWS"        // EWS
	ServiceIMAP       ServiceType = "ServiceIMAP"       // IMAP, Kerio Outlook Connector
	ServiceKoff       ServiceType = "ServiceKoff"       // Kerio Outlook Connector (Offline Edition)
	ServicePOP3       ServiceType = "ServicePOP3"       // POP3
	ServiceWebDAV     ServiceType = "ServiceWebDAV"     // WebDAV, CalDAV, CardDAV
	ServiceWebMail    ServiceType = "ServiceWebMail"    // WebMail
	ServiceXMPP       ServiceType = "ServiceXMPP"       // XMPP
)

type AccessPolicyConnectionRuleType string

const (
	ServiceAllowed   AccessPolicyConnectionRuleType = "ServiceAllowed"   // service is allowed
	ServiceDenied    AccessPolicyConnectionRuleType = "ServiceDenied"    // service is forbidden
	ServiceIpAllowed AccessPolicyConnectionRuleType = "ServiceIpAllowed" // service is allowed for specific IP group
	ServiceIpDenied  AccessPolicyConnectionRuleType = "ServiceIpDenied"  // service is forbidden for specific IP group
)

type AccessPolicyConnectionRule struct {
	Type    AccessPolicyConnectionRuleType `json:"type"`    // type of rule
	GroupId KId                            `json:"groupId"` // if type of rule is 'ServiceIpAllowed/Denied' there is ID of IP Group
}

// AccessPolicyRule - Access policy rule details.
type AccessPolicyRule struct {
	Id      KId                        `json:"id"`      // [READ-ONLY] [REQUIRED FOR SET] global identification
	GroupId KId                        `json:"groupId"` // [REQUIRED FOR CREATE] global identification of AccessPolicyGroup
	Service ServiceType                `json:"service"` // type of service
	Rule    AccessPolicyConnectionRule `json:"rule"`    // rule for connections
}

// AccessPolicyRuleList - List of AccessPolicyRule.
type AccessPolicyRuleList []AccessPolicyRule

type ServiceTypeInfo struct {
	Service     ServiceType `json:"service"`     // type of service
	Description string      `json:"description"` // description of service enum
}

type ServiceTypeInfoList []ServiceTypeInfo

// AccessPolicyGroup - Access policy group details.
type AccessPolicyGroup struct {
	Id        KId    `json:"id"`        // [READ-ONLY] [REQUIRED FOR SET] global identification
	Name      string `json:"name"`      // name of policy
	IsDefault bool   `json:"isDefault"` // [READ-ONLY]
}

// AccessPolicyGroupList - List of AccessPolicy.
type AccessPolicyGroupList []AccessPolicyGroup

// Access policies management.

// AccessPolicyCreate - Add new policies.
// Parameters
//	rules - new policies rules
// Return
//	errors - error message list
//	result - list of IDs of created rules
func (c *Connection) AccessPolicyCreate(rules AccessPolicyRuleList) (ErrorList, CreateResultList, error) {
	params := struct {
		Rules AccessPolicyRuleList `json:"rules"`
	}{rules}
	data, err := c.CallRaw("AccessPolicy.create", params)
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

// AccessPolicyCreateGroupList - Create the list of groups.
// Parameters
//	groups - list of groups to create
// Return
//	errors - error message list
//	result - list of IDs of created groups
func (c *Connection) AccessPolicyCreateGroupList(groups AccessPolicyGroupList) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups AccessPolicyGroupList `json:"groups"`
	}{groups}
	data, err := c.CallRaw("AccessPolicy.createGroupList", params)
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

// AccessPolicyGet - Obtain a list of policies.
// Parameters
//	query - query attributes and limits
// Return
//	list - policies
//	totalItems - number of policies found
func (c *Connection) AccessPolicyGet(query SearchQuery) (AccessPolicyRuleList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("AccessPolicy.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       AccessPolicyRuleList `json:"list"`
			TotalItems int                  `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// AccessPolicyGetGroupList - Get the list of groups, sorted in ascending order.
// Return
//	groups - list of Access policy groups
func (c *Connection) AccessPolicyGetGroupList() (AccessPolicyGroupList, error) {
	data, err := c.CallRaw("AccessPolicy.getGroupList", nil)
	if err != nil {
		return nil, err
	}
	groups := struct {
		Result struct {
			Groups AccessPolicyGroupList `json:"groups"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &groups)
	return groups.Result.Groups, err
}

// AccessPolicyGetServiceList - Get the list of services.
// Return
//	services - list of service info
func (c *Connection) AccessPolicyGetServiceList() (ServiceTypeInfoList, error) {
	data, err := c.CallRaw("AccessPolicy.getServiceList", nil)
	if err != nil {
		return nil, err
	}
	services := struct {
		Result struct {
			Services ServiceTypeInfoList `json:"services"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &services)
	return services.Result.Services, err
}

// AccessPolicyRemove - Remove policies.
// Parameters
//	ruleIds - list of IDs of policy to be removed
// Return
//	errors - error message list
func (c *Connection) AccessPolicyRemove(ruleIds KIdList) (ErrorList, error) {
	params := struct {
		RuleIds KIdList `json:"ruleIds"`
	}{ruleIds}
	data, err := c.CallRaw("AccessPolicy.remove", params)
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

// AccessPolicyRemoveGroupList - Remove the list of groups.
// Parameters
//	groupIds - list of IDs of group policy to be removed
// Return
//	errors - error message list
func (c *Connection) AccessPolicyRemoveGroupList(groupIds KIdList) (ErrorList, error) {
	params := struct {
		GroupIds KIdList `json:"groupIds"`
	}{groupIds}
	data, err := c.CallRaw("AccessPolicy.removeGroupList", params)
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

// AccessPolicySet - Set policy details.
// Parameters
//	rules - rules to save
// Return
//	errors - error message list
func (c *Connection) AccessPolicySet(rules AccessPolicyRuleList) (ErrorList, error) {
	params := struct {
		Rules AccessPolicyRuleList `json:"rules"`
	}{rules}
	data, err := c.CallRaw("AccessPolicy.set", params)
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

// AccessPolicySetGroupList - Set the list of groups.
// Parameters
//	groups - list of group to set
// Return
//	errors - error message list
func (c *Connection) AccessPolicySetGroupList(groups AccessPolicyGroupList) (ErrorList, error) {
	params := struct {
		Groups AccessPolicyGroupList `json:"groups"`
	}{groups}
	data, err := c.CallRaw("AccessPolicy.setGroupList", params)
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
