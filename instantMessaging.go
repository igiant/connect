package connect

import "encoding/json"

type XmppSettings struct {
	SendOutsideEnabled          bool `json:"sendOutsideEnabled"`          // Sending messages outside the company is enabled
	SendOutsideEnabledIsRunning bool `json:"sendOutsideEnabledIsRunning"` // [READ-ONLY] Sending messages outside the company is really running and is functional
}

type XMPPConfiguration struct {
	DnsARecord             bool `json:"dnsARecord"`
	DnsSRVRecordClient     bool `json:"dnsSRVRecordClient"`
	DnsSRVRecordServer     bool `json:"dnsSRVRecordServer"`
	XmppPingExternalServer bool `json:"xmppPingExternalServer"`
}

// InstantMessagingGet - Get settings of XMPP server
// Return
//	settings - Sign On settings
func (s *ServerConnection) InstantMessagingGet() (*XmppSettings, error) {
	data, err := s.CallRaw("InstantMessaging.get", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings XmppSettings `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// InstantMessagingSet - Set settings of XMPP server
// Parameters
//	settings - Sign On settings
func (s *ServerConnection) InstantMessagingSet(settings XmppSettings) error {
	params := struct {
		Settings XmppSettings `json:"settings"`
	}{settings}
	_, err := s.CallRaw("InstantMessaging.set", params)
	return err
}

// InstantMessagingCheckXMPPConfiguration - Check XMPP configuration for all domains
func (s *ServerConnection) InstantMessagingCheckXMPPConfiguration(domainId KId) (*XMPPConfiguration, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("InstantMessaging.checkXMPPConfiguration", params)
	if err != nil {
		return nil, err
	}
	configuration := struct {
		Result struct {
			Configuration XMPPConfiguration `json:"configuration"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &configuration)
	return &configuration.Result.Configuration, err
}
