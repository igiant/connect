package connect

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

// TODO Add Methods
