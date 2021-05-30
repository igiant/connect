package connect

import "encoding/json"

type SecurityPolicyMode string

const (
	SPNoRestrictions         SecurityPolicyMode = "SPNoRestrictions"         // No restriction
	SPAuthenticationRequired SecurityPolicyMode = "SPAuthenticationRequired" // Require secure authentication
	SPEncryptionRequired     SecurityPolicyMode = "SPEncryptionRequired"     // Require encrypted connection
)

type AuthenticationMethodList []OptionalString

type AntiHammeringOptions struct {
	IsEnabled           bool           `json:"isEnabled"`           // Enable/disable Anti-Hammering
	FailedLoginsToBlock int            `json:"failedLoginsToBlock"` // Count of failed logins within minute to start blocking
	MinutesToBlock      int            `json:"minutesToBlock"`      // Minutes to keep blocking IP
	ExceptionIpGroup    OptionalEntity `json:"exceptionIpGroup"`    // switchable custom white list IP group
}

type SecurityPolicyOptions struct {
	Mode                         SecurityPolicyMode       `json:"mode"`
	AuthenticationExceptionGroup OptionalEntity           `json:"authenticationExceptionGroup"` // Is used if mode == SPAuthenticationRequired
	EncryptionExceptionGroup     OptionalEntity           `json:"encryptionExceptionGroup"`     // Is used if mode == SPEncryptionRequired
	AuthenticationMethods        AuthenticationMethodList `json:"authenticationMethods"`        // List of authentication methods and its status. In any set operation, all methods should be present in list. Methods which are not specified are reset to 'enabled'.
	AllowNtlmForKerberosUsers    bool                     `json:"allowNtlmForKerberosUsers"`    // Allow NTLM authentication for users with Kerberos� authentication (for Active Directory� users)
	EnableLockout                bool                     `json:"enableLockout"`                // Enable/disable account lockout feature
	FailedLoginsToLock           int                      `json:"failedLoginsToLock"`           // Count of failed logins to lock user account
	MinutesToUnlock              int                      `json:"minutesToUnlock"`              // Minutes to unlock locked account
	AntiHammering                AntiHammeringOptions     `json:"antiHammering"`                // Anti-Hammering settings
}

// SecurityPolicyGet - Obtain Security Policy options.
// Return
//	options - current security options
func (c *ServerConnection) SecurityPolicyGet() (*SecurityPolicyOptions, error) {
	data, err := c.CallRaw("SecurityPolicy.get", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options SecurityPolicyOptions `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// SecurityPolicySet - Set Security Policy options.
// Parameters
//	options - options to be updated
func (c *ServerConnection) SecurityPolicySet(options SecurityPolicyOptions) error {
	params := struct {
		Options SecurityPolicyOptions `json:"options"`
	}{options}
	_, err := c.CallRaw("SecurityPolicy.set", params)
	return err
}

// SecurityPolicyUnlockAllAccounts - Unlock all locked accounts immediately.
func (c *ServerConnection) SecurityPolicyUnlockAllAccounts() error {
	_, err := c.CallRaw("SecurityPolicy.unlockAllAccounts", nil)
	return err
}
