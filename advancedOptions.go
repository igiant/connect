package connect

import "encoding/json"

type MiscellaneousOptions struct {
	LogHostNames              bool `json:"logHostNames"`              // Log hostnames for incoming connections
	ShowProgramNameAndVersion bool `json:"showProgramNameAndVersion"` // Show program name and version in network communication for non-authenticated users
	InsertXEnvelopeTo         bool `json:"insertXEnvelopeTo"`         // Insert X-Envelope-To header to locally delivered messages
	EnableTNEFDecoding        bool `json:"enableTNEFDecoding"`        // Enable decoding of TNEF messages (winmail.dat attachments)
	EnableUUEncodedConversion bool `json:"enableUUEncodedConversion"` // Enable conversion of uuencoded messages to MIME
}

type StoreDirectoryOptions struct {
	StorePath         string             `json:"storePath"`         // Path to the store directory
	ArchivePath       string             `json:"archivePath"`       // Path to the archive
	BackupPath        string             `json:"backupPath"`        // Path to the backup
	WatchdogSoftLimit ByteValueWithUnits `json:"watchdogSoftLimit"` // If the available disk space falls below this value, a warning message is displayed
	WatchdogHardLimit ByteValueWithUnits `json:"watchdogHardLimit"` // If the available disk space falls below this value, Kerio MailServer is stopped and an error message is displayed. Administrator's action is required.
}

type MasterAuthenticationOptions struct {
	IsEnabled        bool   `json:"isEnabled"`        // Enable master authentication to this server
	GroupRestriction KId    `json:"groupRestriction"` // Allow master authentication only from IP address group
	Password         string `json:"password"`         // [WriteOnly] Master password
}

type HttpProxyOptions struct {
	IsEnabled              bool   `json:"isEnabled"` // Use HTTP proxy for antivirus updates, Kerio update checker and other web services
	Address                string `json:"address"`
	Port                   int    `json:"port"`
	RequiresAuthentication bool   `json:"requiresAuthentication"` // Proxy requires authentication
	UserName               string `json:"userName"`
	Password               string `json:"password"`
}

type OperatorOptions struct {
	IsEnabled bool   `json:"isEnabled"`
	Address   string `json:"address"`
}

type UpdateCheckerOptions struct {
	AutoCheck         bool           `json:"autoCheck"`         // Automatically check for new versions
	CheckBetaVersion  bool           `json:"checkBetaVersion"`  // Check also for beta versions
	TimeFromLastCheck DistanceOrNull `json:"timeFromLastCheck"` // [ReadOnly]
	DownloadedFile    string         `json:"downloadedFile"`    // [ReadOnly]
	UpdateInfo        UpdateInfo     `json:"updateInfo"`        // [ReadOnly]
	KocVersion        string         `json:"kocVersion"`        // [ReadOnly]
	KoffVersion       string         `json:"koffVersion"`       // [ReadOnly]
	KspVersion        string         `json:"kspVersion"`        // [ReadOnly]
	KscVersion        string         `json:"kscVersion"`        // [ReadOnly]
}

type KoffUpgradePolicy string

const (
	KoffUPolicyAskVoluntary          KoffUpgradePolicy = "KoffUPolicyAskVoluntary"          // Ask user for each version change and do not allow the update.
	KoffUPolicyAskRequired           KoffUpgradePolicy = "KoffUPolicyAskRequired"           // Ask user for each version change and require the update.
	KoffUPolicyAlwaysSilent          KoffUpgradePolicy = "KoffUPolicyAlwaysSilent"          // Do update for each version change. Update silently when Outlook starts. Ask users when Outlook is running and require update.
	KoffUPolicyOnStartSilent         KoffUpgradePolicy = "KoffUPolicyOnStartSilent"         // default, available in WebAdmin. Do update for each version change. Update silently when Outlook starts. When Outlook is running do nothing and wait for next Outlook start.
	KoffUPolicyOnlyIfNecessaryAsk    KoffUpgradePolicy = "KoffUPolicyOnlyIfNecessaryAsk"    // Update only if necessary. Ask users and require the update.
	KoffUPolicyOnlyIfNecessarySilent KoffUpgradePolicy = "KoffUPolicyOnlyIfNecessarySilent" // available in WebAdmin, Update only if necessary. Update silently when Outlook starts. Ask users when Outlook is running and require update.
)

type KoffOptions struct {
	UpgradePolicy KoffUpgradePolicy `json:"upgradePolicy"`
}

// FulltextStatus - State of index
type FulltextStatus string

const (
	IndexRebuilding    FulltextStatus = "IndexRebuilding"    // reindexing is in progress
	IndexMessages      FulltextStatus = "IndexMessages"      // indexing new delivered messages
	IndexFinished      FulltextStatus = "IndexFinished"      // reindexing is finnished, it also mean "Up To Date"
	IndexDisabled      FulltextStatus = "IndexDisabled"      // indexing is disabled
	IndexError         FulltextStatus = "IndexError"         // some error occured
	IndexErrorLowSpace FulltextStatus = "IndexErrorLowSpace" // available disk space is below Soft Limit
)

// FulltextRebuildStatus - [READ ONLY] progres of index
type FulltextRebuildStatus struct {
	Status       FulltextStatus `json:"status"`       // [READ ONLY] state of rebuild process
	UsersLeft    int            `json:"usersLeft"`    // [status IndexRebuilding] - the current number of user re-indexed mailboxes
	MessagesLeft int            `json:"messagesLeft"` // [status IndexMessages] - number of new delivered messages to index
	Size         int            `json:"size"`         // index size or estimate size in status IndexMessages or IndexRebuilding
	FreeSpace    int            `json:"freeSpace"`    // free space in path for index files
}

type FulltextSetting struct {
	Enabled bool   `json:"enabled"` // enabled/disabled
	Path    string `json:"path"`    // path to directory where are indexes
}

// FulltextScope - Scope of reindex
type FulltextScope string

const (
	IndexAll    FulltextScope = "IndexAll"    // all users to reindex
	IndexDomain FulltextScope = "IndexDomain" // only users from domain to reindex
	IndexUser   FulltextScope = "IndexUser"   // only user to reindex
)

type FulltextRebuildingCommand struct {
	Scope FulltextScope `json:"scope"`
	Id    KId           `json:"id"` // domain id for scope 'IndexDomain' or user id for scope 'IndexUser'
}

type ButtonColor struct {
	IsEnabled       bool   `json:"isEnabled"` // Is used
	TextColor       string `json:"textColor"`
	BackgroundColor string `json:"backgroundColor"`
}

type AdditionalInfo struct {
	IsEnabled bool   `json:"isEnabled"` // Is used
	Text      string `json:"text"`
}

type WebmailCustomLoginPage struct {
	Logo           CustomImage    `json:"logo"`
	ButtonColor    ButtonColor    `json:"buttonColor"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
}

type WebMailOptions struct {
	MessageSizeLimit       int                    `json:"messageSizeLimit"`       // Maximum size of message that can be sent from the WebMail interface (HTTP POST size)
	SessionExpireTimeout   TimeLimit              `json:"sessionExpireTimeout"`   // Session expire timeout
	MaximumSessionDuration TimeLimit              `json:"maximumSessionDuration"` // Maximum session duration
	ForceLogout            bool                   `json:"forceLogout"`            // Force WebMail logout if user's IP address changes (prevents session hijacking and session fixation attacks)
	CustomLoginPage        WebmailCustomLoginPage `json:"customLoginPage"`        // Use custom logo in WebMail login page
	CustomClientLogo       CustomImage            `json:"customClientLogo"`       // Use custom logo in Kerio Connect Client
}

type UserQuota struct {
	Notification QuotaNotification `json:"notification"` // option for notification
	WarningLimit int               `json:"warningLimit"` // limit in per cent
	Email        string            `json:"email"`        // if quota is exceeded the notification will be sent to this address
}

type AdvancedOptionsSetting struct {
	Miscellaneous        MiscellaneousOptions        `json:"miscellaneous"`
	StoreDirectory       StoreDirectoryOptions       `json:"storeDirectory"`
	MasterAuthentication MasterAuthenticationOptions `json:"masterAuthentication"`
	HttpProxy            HttpProxyOptions            `json:"httpProxy"`
	UpdateChecker        UpdateCheckerOptions        `json:"updateChecker"`
	WebMail              WebMailOptions              `json:"webMail"`
	UserQuota            UserQuota                   `json:"userQuota"`
	Fulltext             FulltextSetting             `json:"fulltext"`
	KoffOptions          KoffOptions                 `json:"koffOptions"`
	OperatorOptions      OperatorOptions             `json:"operatorOptions"`
}

// AdvancedOptionsCheckUpdates - Check for updates.
// Return
//	options - new version details
func (s *ServerConnection) AdvancedOptionsCheckUpdates() (*UpdateCheckerOptions, error) {
	data, err := s.CallRaw("AdvancedOptions.checkUpdates", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options UpdateCheckerOptions `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// AdvancedOptionsGet - Obtain Advanced options.
// Return
//	options - current advanced options
func (s *ServerConnection) AdvancedOptionsGet() (*AdvancedOptionsSetting, error) {
	data, err := s.CallRaw("AdvancedOptions.get", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options AdvancedOptionsSetting `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// AdvancedOptionsGetFulltextStatus - Get information about index status.
// Return
//	info - structure with information
func (s *ServerConnection) AdvancedOptionsGetFulltextStatus() (*FulltextRebuildStatus, error) {
	data, err := s.CallRaw("AdvancedOptions.getFulltextStatus", nil)
	if err != nil {
		return nil, err
	}
	info := struct {
		Result struct {
			Info FulltextRebuildStatus `json:"info"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &info)
	return &info.Result.Info, err
}

// AdvancedOptionsSet - Set advanced options.
// Parameters
//	options - options to be updated
func (s *ServerConnection) AdvancedOptionsSet(options AdvancedOptionsSetting) error {
	params := struct {
		Options AdvancedOptionsSetting `json:"options"`
	}{options}
	_, err := s.CallRaw("AdvancedOptions.set", params)
	return err
}

// AdvancedOptionsStartRebuildFulltext - Launch re-index according parameters.
// Parameters
//	parameters - parameters for launching re-index
func (s *ServerConnection) AdvancedOptionsStartRebuildFulltext(parameters FulltextRebuildingCommand) error {
	params := struct {
		Parameters FulltextRebuildingCommand `json:"parameters"`
	}{parameters}
	_, err := s.CallRaw("AdvancedOptions.startRebuildFulltext", params)
	return err
}
