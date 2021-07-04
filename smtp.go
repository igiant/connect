package connect

import "encoding/json"

type SmtpAuthentication string

const (
	Auth      SmtpAuthentication = "Auth"      // AUTH SMTP command
	Pop3Based SmtpAuthentication = "Pop3Based" // POP3 authentication before SMTP
)

// Relay - SMTP Relay options
type Relay struct {
	IsRestricted              bool           `json:"isRestricted"`              // is relay restricted? false == open relay enabled
	IpAddressGroupName        string         `json:"ipAddressGroupName"`        // output only: name of IP group from which open relay is enabled
	IpAddressGroupId          KId            `json:"ipAddressGroupId"`          // ID of IP group from which open relay is enabled
	AuthenticateOutgoing      bool           `json:"authenticateOutgoing"`      // client must authenticate before sending a message
	ForwardPop3authentication ItemCountLimit `json:"forwardPop3authentication"` // client can be authenticated by POP3 before SMTP relay for certain number of minutes
	HideLocalIpInRcv          bool           `json:"hideLocalIpInRcv"`          // hide local IP address in receive headers
}

type RelayAuthentication struct {
	IsRequired bool               `json:"isRequired"` // is authentication required?
	UserName   string             `json:"userName"`   // a user name for relay server authentication
	Password   string             `json:"password"`   // a password for relay server authentication, it's allowed to send ID of a rule to copy a password
	AuthType   SmtpAuthentication `json:"authType"`   // type of authentication
}

type IpLimit struct {
	MaximumFromIp      ItemCountLimit `json:"maximumFromIp"`      // maximum number of messages per hour from 1 IP address
	MaximumConnections ItemCountLimit `json:"maximumConnections"` // maximum number of concurrent SMTP connections from 1 IP address
	MaximumUnknowns    ItemCountLimit `json:"maximumUnknowns"`    // maximum number of unknown recipients
	IpAddressGroupName string         `json:"ipAddressGroupName"` // output only: name of IP group on which limits are NOT applied
	IpAddressGroupId   KId            `json:"ipAddressGroupId"`   // ID of IP group on which limits are NOT applied
}

type SmtpServerSettings struct {
	RelayControl                 Relay          `json:"relayControl"`                 // relay control options
	IpBasedLimit                 IpLimit        `json:"ipBasedLimit"`                 // limits based on IP address
	BlockUnknownDns              bool           `json:"blockUnknownDns"`              // block is sender's mail domain was not found in DNS
	VerifyClientRDnsEntry        bool           `json:"verifyClientRDnsEntry"`        // block if client's IP address has no reverse DNS entry (PTR)
	RequireLocalDomainSenderAuth bool           `json:"requireLocalDomainSenderAuth"` // require SMTP authentication if the sender's address is from a local domain
	MaximumRecipients            ItemCountLimit `json:"maximumRecipients"`            // maximum number of recipients in 1 message
	MaximumSmtpFailures          ItemCountLimit `json:"maximumSmtpFailures"`          // maximum number of failed commands in 1 SMTP session
	MessageSize                  SizeLimit      `json:"messageSize"`                  // limit for incomming SMTP message size
	ReaderHops                   int            `json:"readerHops"`                   // maximum number of accepted received headers (hops)
	UseSSL                       bool           `json:"useSSL"`                       // use SSL if supported by remote SMTP server
	MaximumThreads               int            `json:"maximumThreads"`               // maximum number of delivery threads
	RetryInterval                TimeLimit      `json:"retryInterval"`                // delivery retry interval
	BounceInterval               TimeLimit      `json:"bounceInterval"`               // bounce the message to the sender if not delivered in defined time
	WarningInterval              TimeLimit      `json:"warningInterval"`              // warn sender if a message is not delivered within define time
	ReportLanguage               string         `json:"reportLanguage"`               // 2 char abbreviation; we don't support reports added by user
}

type RelayRuleCondType string

const (
	RelayCondNone      RelayRuleCondType = "RelayCondNone"      // there is none condition; it's always true
	RelayCondRecipient RelayRuleCondType = "RelayCondRecipient" // recipient from envelope
	RelayCondSender    RelayRuleCondType = "RelayCondSender"    // sender from envelope
)

type RelayRuleComparatorType string

const (
	RelayCompEqual    RelayRuleComparatorType = "RelayCompEqual"
	RelayCompNotEqual RelayRuleComparatorType = "RelayCompNotEqual"
)

type RelayRuleCondition struct {
	Test       RelayRuleCondType       `json:"test"`
	Comparator RelayRuleComparatorType `json:"comparator"`
	Pattern    string                  `json:"pattern"` // patern shoud be type match E.g. *@example.com
}

type RelayDeliveryRule struct {
	Id             KId                 `json:"id"` // [READ-ONLY] global identification
	IsEnabled      bool                `json:"isEnabled"`
	Description    string              `json:"description"`    // contains rule description
	HostName       string              `json:"hostName"`       // relay server hostname
	Port           int                 `json:"port"`           // relay server port
	Authentication RelayAuthentication `json:"authentication"` // relay server authentication parameters
	Condition      RelayRuleCondition  `json:"condition"`
}

type RelayDeliveryRuleList []RelayDeliveryRule

type FilterRule struct {
	IsIncomplete   bool                `json:"isIncomplete"`   // if rule is not completed (it does not contain any definition of conditions and actions)
	EvaluationMode EvaluationModeType  `json:"evaluationMode"` // determines evaluation mod of initial conditions
	Conditions     FilterConditionList `json:"conditions"`     // list of rule's initial conditions
	Actions        FilterActionList    `json:"actions"`        // list of rule's actions (performed if initial conditions are meet)
}

type DeliveryRule struct {
	Id          KId        `json:"id"`          // [READ-ONLY] global identification
	IsEnabled   bool       `json:"isEnabled"`   // says whether rule is enabled
	Description string     `json:"description"` // contains rules description
	Rule        FilterRule `json:"rule"`
}

type DeliveryRuleList []DeliveryRule

// SmtpGet - Obtain SMTP server settings.
// Return
//	server - SMTP settings
func (s *ServerConnection) SmtpGet() (*SmtpServerSettings, error) {
	data, err := s.CallRaw("Smtp.get", nil)
	if err != nil {
		return nil, err
	}
	server := struct {
		Result struct {
			Server SmtpServerSettings `json:"server"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &server)
	return &server.Result.Server, err
}

// SmtpSet - Change SMTP server settings.
//	server - SMTP settings
func (s *ServerConnection) SmtpSet(server SmtpServerSettings) error {
	params := struct {
		Server SmtpServerSettings `json:"server"`
	}{server}
	_, err := s.CallRaw("Smtp.set", params)
	return err
}

// SmtpGetRelayDeliveryRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpGetRelayDeliveryRuleList() (RelayDeliveryRuleList, error) {
	data, err := s.CallRaw("Smtp.getRelayDeliveryRuleList", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List RelayDeliveryRuleList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// SmtpSetRelayDeliveryRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpSetRelayDeliveryRuleList(list RelayDeliveryRuleList) error {
	params := struct {
		List RelayDeliveryRuleList `json:"list"`
	}{list}
	_, err := s.CallRaw("Smtp.setRelayDeliveryRuleList", params)
	return err
}

// SmtpGetIncomingRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpGetIncomingRuleList() (DeliveryRuleList, error) {
	data, err := s.CallRaw("Smtp.getIncomingRuleList", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List DeliveryRuleList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// SmtpSetIncomingRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpSetIncomingRuleList(list DeliveryRuleList) error {
	params := struct {
		List DeliveryRuleList `json:"list"`
	}{list}
	_, err := s.CallRaw("Smtp.setIncomingRuleList", params)
	return err
}

// SmtpGetOutgoingRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpGetOutgoingRuleList() (DeliveryRuleList, error) {
	data, err := s.CallRaw("Smtp.getOutgoingRuleList", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List DeliveryRuleList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// SmtpSetOutgoingRuleList - Change SMTP server settings.
func (s *ServerConnection) SmtpSetOutgoingRuleList(list DeliveryRuleList) error {
	params := struct {
		List DeliveryRuleList `json:"list"`
	}{list}
	_, err := s.CallRaw("Smtp.setOutgoingRuleList", params)
	return err
}
