package connect

import "encoding/json"

type AttachmentAction string

const (
	Block  AttachmentAction = "Block"
	Accept AttachmentAction = "Accept"
)

type AttachmentType string

const (
	FileName AttachmentType = "FileName"
	MimeType AttachmentType = "MimeType"
)

// AttachmentItem - Attachment filter rule item
type AttachmentItem struct {
	Id          KId              `json:"id"`
	Enabled     bool             `json:"enabled"`     // the rule is in use
	Type        AttachmentType   `json:"type"`        // type of the rule
	Content     string           `json:"content"`     // *,? wildcards are supported
	Action      AttachmentAction `json:"action"`      // what to do if the rule maches
	Description string           `json:"description"` //
}

type AttachmentItemList []AttachmentItem

// AttachmentSetting - Attachment filter settings
type AttachmentSetting struct {
	Enabled                bool           `json:"enabled"`                // attachment filter is on/off
	WarnSender             bool           `json:"warnSender"`             // sender will (not) obtain warning message
	ForwardOriginal        OptionalString `json:"forwardOriginal"`        // where to forward original message
	ForwardFiltered        OptionalString `json:"forwardFiltered"`        // where to forward filtered message
	EnableZipContentFilter bool           `json:"enableZipContentFilter"` // checks zip content for prohibited extennsions
}

// AntivirusOption - Note: fields name and content must be assigned if used in set methods
type AntivirusOption struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	DefaultValue string `json:"defaultValue"` // read only value
}

type AntivirusOptionList []AntivirusOption

// AntivirusPlugin - Note: field id must be assigned if used in set methods
type AntivirusPlugin struct {
	Id                  string              `json:"id"`          // example: avir_avg
	Description         string              `json:"description"` // example: AVG Email Server Edition
	AreOptionsAvailable bool                `json:"areOptionsAvailable"`
	Options             AntivirusOptionList `json:"options"`
}

type AntivirusPluginList []AntivirusPlugin

type IntegratedEngine struct {
	CheckForUpdates         bool           `json:"checkForUpdates"`   // should we periodically ask for a new version?
	UpdatePeriod            int            `json:"updatePeriod"`      // update checking period in hours
	DatabaseAge             DistanceOrNull `json:"databaseAge"`       // how old is virus database
	LastUpdateCheck         DistanceOrNull `json:"lastUpdateCheck"`   // how long is since last database update check
	DatabaseVersion         string         `json:"databaseVersion"`   // virus database version
	EngineVersion           string         `json:"engineVersion"`     // scanning engine version
	IsPluginAvailable       bool           `json:"isPluginAvailable"` // says if plugins dll is on hardrive
	IsLiveProtectionEnabled bool           `json:"isLiveProtectionEnabled"`
}

// ReactionOnVirus - What to do with an infected file
type ReactionOnVirus string

const (
	DiscardMessage ReactionOnVirus = "DiscardMessage" // completely dicard the message
	RemoveVirus    ReactionOnVirus = "RemoveVirus"    // deliver the message but remove malicious code
)

// ReactionOnNotScanned - What to do with a corrupted or encrypted file
type ReactionOnNotScanned string

const (
	DeliverWithWarning ReactionOnNotScanned = "DeliverWithWarning" // deliver original message with prepended warning
	SameAsVirus        ReactionOnNotScanned = "SameAsVirus"        // the same reaction as for ReactionOnVirus
)

// AntivirusStatus - Are all possible states covered
type AntivirusStatus string

const (
	AntivirusOk     AntivirusStatus = "AntivirusOk"     // no message is needed
	NoAntivirus     AntivirusStatus = "NoAntivirus"     // neither internal nor external antivirus is active
	InternalFailure AntivirusStatus = "InternalFailure" // problem with internal intivirus
	ExternalFailure AntivirusStatus = "ExternalFailure" // problem with external intivirus
	DoubleFailer    AntivirusStatus = "DoubleFailer"    // both internal and external antivirus has failed
)

type FoundVirusBehavior struct {
	Reaction        ReactionOnVirus `json:"reaction"`
	ForwardOriginal OptionalString  `json:"forwardOriginal"` // should be original message forwarded?
	ForwardFiltered OptionalString  `json:"forwardFiltered"` // should be filtered message forwarded?
}

type AntivirusSetting struct {
	UseIntegrated      bool                 `json:"useIntegrated"`      // integrated antivirus is used?
	UseExternal        bool                 `json:"useExternal"`        // an external antivirus is used? note: both internal and extenal can be used together
	Status             AntivirusStatus      `json:"status"`             // status of antivirus to be used for informative massage
	Plugins            AntivirusPluginList  `json:"plugins"`            // list of available antivirus plugins
	SelectedId         string               `json:"selectedId"`         // identifier of currently selected antivirus plugin
	Engine             IntegratedEngine     `json:"engine"`             // integrated engine settings
	VirusReaction      FoundVirusBehavior   `json:"virusReaction"`      // found virus reaction setting
	NotScannedReaction ReactionOnNotScanned `json:"notScannedReaction"` // found corruption or encryption reaction type
}

type UpdateStatus string

const (
	UpdateStarted      UpdateStatus = "UpdateStarted"
	UpdateFinished     UpdateStatus = "UpdateFinished"
	UpdateError        UpdateStatus = "UpdateError"
	UpdateDownloadIni  UpdateStatus = "UpdateDownloadIni"
	UpdateDownloadData UpdateStatus = "UpdateDownloadData"
	UpdateUpToDate     UpdateStatus = "UpdateUpToDate"
)

type IntegratedAvirUpdateStatus struct {
	Status  UpdateStatus `json:"status"`  // state of update process
	Percent int          `json:"percent"` // percent of downloaded data
}

type BlockOrScore string

const (
	BlockMessage BlockOrScore = "BlockMessage" // block the message
	ScoreMessage BlockOrScore = "ScoreMessage" // add SPAM score to the message
)

// BlackListSetting - Custom setting of blacklist spammer IP addresses
type BlackListSetting struct {
	Enabled bool         `json:"enabled"`
	Id      KId          `json:"id"` // global identifier
	Name    string       `json:"name"`
	Action  BlockOrScore `json:"action"`
	Score   int          `json:"score"`
}

type SpamAction string

const (
	LogToSecurity SpamAction = "LogToSecurity" // only log to security log
	BlockAction   SpamAction = "BlockAction"   // block the meassage
	ScoreAction   SpamAction = "ScoreAction"   // increase spam score
)

type CallerId struct {
	Enabled          bool           `json:"enabled"`
	Action           SpamAction     `json:"action"`
	Score            int            `json:"score"`
	ApplyOnTesting   bool           `json:"applyOnTesting"`
	ExceptionIpGroup OptionalEntity `json:"exceptionIpGroup"` // switchable custom white list IP group
}

type Spf struct {
	Enabled          bool           `json:"enabled"`
	Action           SpamAction     `json:"action"`
	Score            int            `json:"score"`
	ExceptionIpGroup OptionalEntity `json:"exceptionIpGroup"` // switchable custom white list IP group
}

type Repellent struct {
	Enabled          bool           `json:"enabled"`
	Delay            int            `json:"delay"`
	CustomWhiteList  OptionalEntity `json:"customWhiteList"`  // switchable custom white list IP group
	ReportToSecurity bool           `json:"reportToSecurity"` // do (not) report a spam attack to security log
}

type BayesState string

const (
	Disabled BayesState = "Disabled" // Bayes database statistics are not provided
	Learning BayesState = "Learning"
	Active   BayesState = "Active"
)

// GreylistingStatus - State of the Greylisting client.
type GreylistingStatus string

const (
	GreylistingOff   GreylistingStatus = "GreylistingOff"
	GreylistingOn    GreylistingStatus = "GreylistingOn"
	GreylistingError GreylistingStatus = "GreylistingError" // Greylisting encountered an error. Call Content.testGreylistConnection() for a more detailed error description.
)

type Greylisting struct {
	Enabled bool `json:"enabled"` // is greylisting enabled?
	/// When enabled is set to true, the setAntiSpamSetting method attempts to connect to the greylisting service asynchronously.
	/// When enabled is set to false, connection to the greylisting service is closed.
	CustomWhiteList  OptionalEntity    `json:"customWhiteList"`  // switchable custom whitelist IP group
	Status           GreylistingStatus `json:"status"`           // read only: current status
	MessagesAccepted string            `json:"messagesAccepted"` // read only: messages accepted
	MessagesDelayed  string            `json:"messagesDelayed"`  // read only: messages temoprarily rejected
	MessagesSkipped  string            `json:"messagesSkipped"`  // read only: messages skipped
}

type IntegratedAntiSpamStatus string

const (
	AntiSpamReady          IntegratedAntiSpamStatus = "AntiSpamReady"
	AntiSpamDisabled       IntegratedAntiSpamStatus = "AntiSpamDisabled"
	AntiSpamNotLicenced    IntegratedAntiSpamStatus = "AntiSpamNotLicenced"
	AntiSpamNotInitialized IntegratedAntiSpamStatus = "AntiSpamNotInitialized"
	AntiSpamNotConnected   IntegratedAntiSpamStatus = "AntiSpamNotConnected"
)

type IntegratedAntiSpamEngine struct {
	Enabled       bool                     `json:"enabled"`
	Score         int                      `json:"score"`         // Spam score(default 10)
	NegativeScore int                      `json:"negativeScore"` // Score for legit messages (default 0)
	SubmitSpam    bool                     `json:"submitSpam"`    // Submit spam samples(default on)
	SubmitLegit   bool                     `json:"submitLegit"`   // Submit legit samples(default on)
	Status        IntegratedAntiSpamStatus `json:"status"`        // [READ-ONLY]
	IsLicensed    bool                     `json:"isLicensed"`    // [READ-ONLY] Is license valid and not expired? If false engine is not running no matter on value in enabled.
}

type AntiSpamSetting struct {
	IsRatingEnabled           bool                     `json:"isRatingEnabled"`      // is spam filter rating enabled?
	IsRatingRelayEnabled      bool                     `json:"isRatingRelayEnabled"` // is rating of messages sent from trustworthy relay agents enabled?
	TagScore                  int                      `json:"tagScore"`
	BlockScore                int                      `json:"blockScore"`
	SubjectPrefix             OptionalString           `json:"subjectPrefix"`           // SPAM is marked with this prefix
	SendBounce                bool                     `json:"sendBounce"`              // send bounce message to the sender of SPAM?
	QuarantineAddress         OptionalString           `json:"quarantineAddress"`       // forward SPAM to a Quarantine address?
	CustomWhiteList           OptionalEntity           `json:"customWhiteList"`         // switchable custom white list IP group
	CustomBlackList           BlackListSetting         `json:"customBlackList"`         // switchable custom blacklist list IP group
	SendBounceCustom          bool                     `json:"sendBounceCustom"`        // send bounce message to the sender if rejection was done by custom rule(s)?
	QuarantineAddressCustom   OptionalString           `json:"quarantineAddressCustom"` // forward custom rules identified SPAM to a Quarantine address?
	UseSurbl                  bool                     `json:"useSurbl"`                // use Spam URI Realtime Block List database?
	FilterStatus              BayesState               `json:"filterStatus"`            // read only: Bayesian filter status
	LearnedAsSpam             int                      `json:"learnedAsSpam"`           // read only: number of messages that Bayesian filter learned as Spam
	LearnedAsNotSpam          int                      `json:"learnedAsNotSpam"`        // read only: number of messages that Bayesian filter learned as NOT a Spam
	IsCustomSigningKey        bool                     `json:"isCustomSigningKey"`      // Custom signing key is used for DKIM validation
	CallerSetting             CallerId                 `json:"callerSetting"`           // Caller ID setting
	CallerUrl                 string                   `json:"callerUrl"`               // read only: Caller ID URL with detailed info
	SpfSetting                Spf                      `json:"spfSetting"`              // Sender Policy Framework setting
	RepellentSetting          Repellent                `json:"repellentSetting"`
	GreylistingStatus         Greylisting              `json:"greylistingStatus"`
	UseCustomRulesInSmtp      bool                     `json:"useCustomRulesInSmtp"`      //
	IntegratedAntiSpamSetting IntegratedAntiSpamEngine `json:"integratedAntiSpamSetting"` // Kerio Anti-Spam
}

type BlackList struct {
	Id          KId          `json:"id"` // global identifier
	Enabled     bool         `json:"enabled"`
	DnsSuffix   string       `json:"dnsSuffix"`
	Description string       `json:"description"`
	Action      BlockOrScore `json:"action"` // what to do if IP address is found on blacklist
	Score       int          `json:"score"`
	AskDirectly bool         `json:"askDirectly"`
}

type BlackListList []BlackList

type CustomRuleKind string

const (
	Header CustomRuleKind = "Header"
	Body   CustomRuleKind = "Body"
)

type CustomRuleType string

const (
	IsEmpty           CustomRuleType = "IsEmpty"
	IsMissing         CustomRuleType = "IsMissing"
	ContainsAddress   CustomRuleType = "ContainsAddress"
	ContainsDomain    CustomRuleType = "ContainsDomain"
	ContainsSubstring CustomRuleType = "ContainsSubstring"
	ContainsBinary    CustomRuleType = "ContainsBinary"
)

type CustomRuleAction string

const (
	TreatAsSpam       CustomRuleAction = "TreatAsSpam"
	TreatAsNotSpam    CustomRuleAction = "TreatAsNotSpam"
	IncreaseSpamScore CustomRuleAction = "IncreaseSpamScore"
)

type CustomRule struct {
	Id          KId              `json:"id"` // global identifier
	Enabled     bool             `json:"enabled"`
	Kind        CustomRuleKind   `json:"kind"`
	Header      string           `json:"header"`
	Content     string           `json:"content"`
	Description string           `json:"description"`
	Type        CustomRuleType   `json:"type"`
	Action      CustomRuleAction `json:"action"`
	Score       int              `json:"score"`
	LastUsed    DistanceOrNull   `json:"lastUsed"`
}

type CustomRuleList []CustomRule

type HourOrDay string

const (
	Hour HourOrDay = "Hour"
	Day  HourOrDay = "Day"
)

// ContentAddBlackLists - Add a blacklist item.
//	items - array of new items
// Return
//	errors - error message list
func (s *ServerConnection) ContentAddBlackLists(items BlackListList) (ErrorList, error) {
	params := struct {
		Items BlackListList `json:"items"`
	}{items}
	data, err := s.CallRaw("Content.addBlackLists", params)
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

// ContentGetAntiSpamSetting - Get antiSPAM settings.
// Return
//	setting - new antivirus filter settings
func (s *ServerConnection) ContentGetAntiSpamSetting() (*AntiSpamSetting, error) {
	data, err := s.CallRaw("Content.getAntiSpamSetting", nil)
	if err != nil {
		return nil, err
	}
	setting := struct {
		Result struct {
			Setting AntiSpamSetting `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return &setting.Result.Setting, err
}

// ContentGetAntivirusSetting - Get antivirus filter settings.
// Return
//	setting - new antivirus filter settings
func (s *ServerConnection) ContentGetAntivirusSetting() (*AntivirusSetting, error) {
	data, err := s.CallRaw("Content.getAntivirusSetting", nil)
	if err != nil {
		return nil, err
	}
	setting := struct {
		Result struct {
			Setting AntivirusSetting `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return &setting.Result.Setting, err
}

// ContentGetAttachmentRules - Get a list of attachment filter rules.
// Return
//	filterRules - attachment filter rules
func (s *ServerConnection) ContentGetAttachmentRules() (AttachmentItemList, error) {
	data, err := s.CallRaw("Content.getAttachmentRules", nil)
	if err != nil {
		return nil, err
	}
	filterRules := struct {
		Result struct {
			FilterRules AttachmentItemList `json:"filterRules"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &filterRules)
	return filterRules.Result.FilterRules, err
}

// ContentGetAttachmentSetting - Obtain attachment filter settings.
// Return
//	setting - current attachment filter settings
func (s *ServerConnection) ContentGetAttachmentSetting() (*AttachmentSetting, error) {
	data, err := s.CallRaw("Content.getAttachmentSetting", nil)
	if err != nil {
		return nil, err
	}
	setting := struct {
		Result struct {
			Setting AttachmentSetting `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return &setting.Result.Setting, err
}

// ContentGetAvailableAttachments - When adding a new attachment rule this can be used to find out available values.
// Return
//	fileNames - list of available file names
//	mimeTypes - list of available MIME types
func (s *ServerConnection) ContentGetAvailableAttachments() (StringList, StringList, error) {
	data, err := s.CallRaw("Content.getAvailableAttachments", nil)
	if err != nil {
		return nil, nil, err
	}
	fileNames := struct {
		Result struct {
			FileNames StringList `json:"fileNames"`
			MimeTypes StringList `json:"mimeTypes"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileNames)
	return fileNames.Result.FileNames, fileNames.Result.MimeTypes, err
}

// ContentGetBlackListList - Obtain all blacklist items.
// Return
//	list - blacklist items
func (s *ServerConnection) ContentGetBlackListList() (BlackListList, error) {
	data, err := s.CallRaw("Content.getBlackListList", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List BlackListList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// ContentGetCustomRuleList - Obtain all custom rules.
//	query - condition and limit definition (orderBy is ignored)
// Return
//	list - custom rules
//  totalItems - amount of rules for given search condition, useful when a limit is defined in search query
func (s *ServerConnection) ContentGetCustomRuleList(query SearchQuery) (CustomRuleList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Content.getCustomRuleList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       CustomRuleList `json:"list"`
			TotalItems int            `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ContentRemoveBlackLists - Remove blacklist items.
//	ids - identifier list of blacklists to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) ContentRemoveBlackLists(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Content.removeBlackLists", params)
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

// ContentRemoveUnusedCustomRules - Remove custom rules which are not used for a specified time.
//	number - how many hours/days is the rule unused
//	unit - which unit is used to measure
func (s *ServerConnection) ContentRemoveUnusedCustomRules(number int, unit HourOrDay) error {
	params := struct {
		Number int       `json:"number"`
		Unit   HourOrDay `json:"unit"`
	}{number, unit}
	_, err := s.CallRaw("Content.removeUnusedCustomRules", params)
	return err
}

// ContentSetAntiSpamSetting - Set antiSPAM filter settings.
//	setting - new antivirus filter settings
func (s *ServerConnection) ContentSetAntiSpamSetting(setting AntiSpamSetting) error {
	params := struct {
		Setting AntiSpamSetting `json:"setting"`
	}{setting}
	_, err := s.CallRaw("Content.setAntiSpamSetting", params)
	return err
}

// ContentSetAntivirusSetting - Set antivirus filter settings.
//	setting - new antivirus filter settingss
// Return
//	errors - error message; Value of inputIndex means type of antivirus (integrated = 0 and external = 1).
func (s *ServerConnection) ContentSetAntivirusSetting(setting AntivirusSetting) (ErrorList, error) {
	params := struct {
		Setting AntivirusSetting `json:"setting"`
	}{setting}
	data, err := s.CallRaw("Content.setAntivirusSetting", params)
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

// ContentSetAttachmentRules - Set list of attachment filter rules.
//	filterRules - attachment filter rules
func (s *ServerConnection) ContentSetAttachmentRules(filterRules AttachmentItemList) error {
	params := struct {
		FilterRules AttachmentItemList `json:"filterRules"`
	}{filterRules}
	_, err := s.CallRaw("Content.setAttachmentRules", params)
	return err
}

// ContentSetAttachmentSetting - Set attachment filter settings.
//	setting - new attachment filter settings
func (s *ServerConnection) ContentSetAttachmentSetting(setting AttachmentSetting) error {
	params := struct {
		Setting AttachmentSetting `json:"setting"`
	}{setting}
	_, err := s.CallRaw("Content.setAttachmentSetting", params)
	return err
}

// ContentSetBlackLists - Set blacklist item.
//	ids - list of blacklist global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) ContentSetBlackLists(ids KIdList, pattern BlackList) (ErrorList, error) {
	params := struct {
		Ids     KIdList   `json:"ids"`
		Pattern BlackList `json:"pattern"`
	}{ids, pattern}
	data, err := s.CallRaw("Content.setBlackLists", params)
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

// ContentSetCustomRuleList - Set custom rules.
//	list - custom rule records
func (s *ServerConnection) ContentSetCustomRuleList(list CustomRuleList) error {
	params := struct {
		List CustomRuleList `json:"list"`
	}{list}
	_, err := s.CallRaw("Content.setCustomRuleList", params)
	return err
}

// ContentTestGreylistConnection - Test connection to the greylisting service. Returns nothing if successful.
func (s *ServerConnection) ContentTestGreylistConnection() error {
	_, err := s.CallRaw("Content.testGreylistConnection", nil)
	return err
}

// ContentTestIntegratedAntiSpamEngine - Test connection to the anti-spam service. Returns nothing if successful.
func (s *ServerConnection) ContentTestIntegratedAntiSpamEngine() (*IntegratedAntiSpamStatus, error) {
	data, err := s.CallRaw("Content.testIntegratedAntiSpamEngine", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status IntegratedAntiSpamStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// ContentUpdateAntivirusStatus - Get progress of antivirus updating.
// Return
//	status - status of the update process
func (s *ServerConnection) ContentUpdateAntivirusStatus() (*IntegratedAvirUpdateStatus, error) {
	data, err := s.CallRaw("Content.updateAntivirusStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status IntegratedAvirUpdateStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// ContentUpdateIntegratedAntivirus - Force update of the integrated antivirus.
func (s *ServerConnection) ContentUpdateIntegratedAntivirus() error {
	_, err := s.CallRaw("Content.updateIntegratedAntivirus", nil)
	return err
}
