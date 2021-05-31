package connect

import "encoding/json"

type SenderPolicyOptions struct {
	AuthenticationRequired bool           `json:"authenticationRequired"` // Require sender authentication for local domains
	AntiSpoofingEnabled    bool           `json:"antiSpoofingEnabled"`    // Is Antispoofing enabled
	IPWhiteList            OptionalEntity `json:"IPWhiteList"`            // Trusted senders (Not chceked by anti-spoofing)
}

// SenderPolicyGet - Obtain Sender Policy options.
// Return
//	options - current sender options
func (s *ServerConnection) SenderPolicyGet() (*SenderPolicyOptions, error) {
	data, err := s.CallRaw("SenderPolicy.get", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options SenderPolicyOptions `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// SenderPolicySet - Set Sender Policy options.
// Parameters
//	options - options to be updated
func (s *ServerConnection) SenderPolicySet(options SenderPolicyOptions) error {
	params := struct {
		Options SenderPolicyOptions `json:"options"`
	}{options}
	_, err := s.CallRaw("SenderPolicy.set", params)
	return err
}
