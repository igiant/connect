package connect

import "encoding/json"

type TriggerType string

const (
	Every TriggerType = "Every"
	At    TriggerType = "At"
)

// TiedTimeUnit - fewer posibilities than TimeUnit
type TiedTimeUnit string

const (
	TMinutes TiedTimeUnit = "TMinutes"
	THours   TiedTimeUnit = "THours"
)

// TimeCondition - Note: all fields must be assigned if used in set methods
type TimeCondition struct {
	Type      TriggerType  `json:"type"`
	Number    int          `json:"number"`    // for type "every"
	Units     TiedTimeUnit `json:"units"`     // for type "every"
	Minutes   int          `json:"minutes"`   // for type "at"
	Hours     int          `json:"hours"`     // for type "at"
	IsLimited bool         `json:"isLimited"` // is trigger limited to specified time range?
	GroupId   KId          `json:"groupId"`   // time range identifier; (ID should be ok because should show all available time ranges anyway)
}

// ScheduledActionAction - Note: all fields must be assigned if used in set methods
type ScheduledActionAction struct {
	SendFromQueue bool `json:"sendFromQueue"` // send messages from outgoing queue
	Pop3Download  bool `json:"pop3Download"`  // download messages from POP3 mailboxes
	SendEtrn      bool `json:"sendEtrn"`      // send ETRN command to invoke mail transfer
}

type ScheduledAction struct {
	Id          KId                   `json:"id"`
	IsActive    bool                  `json:"isActive"` // record is active
	Description string                `json:"description"`
	Condition   TimeCondition         `json:"condition"`
	AllowDialUp bool                  `json:"allowDialUp"` // allow to establish dial-up connection if necessary
	Action      ScheduledActionAction `json:"action"`
}

type ScheduledActionList []ScheduledAction

type EtrnDownload struct {
	Id                    KId    `json:"id"`
	IsActive              bool   `json:"isActive"`
	Server                string `json:"server"`  // server URL
	Domains               string `json:"domains"` // semicolon separated list of domain names
	Description           string `json:"description"`
	RequireAuthentication bool   `json:"requireAuthentication"` // Is authentication required?
	UserName              string `json:"userName"`              // make sense only if authentication is required
	Password              string `json:"password"`              // make sense only if authentication is required
}

type EtrnDownloadList []EtrnDownload

type SslMode string

const (
	NoSsl       SslMode = "NoSsl"
	SpecialPort SslMode = "SpecialPort"
	StlsCommand SslMode = "StlsCommand"
)

type Pop3Authentication string

const (
	PlainPop3 Pop3Authentication = "PlainPop3"
	Apop      Pop3Authentication = "Apop"
)

// If leaveOnServer ID enabled, messages are left on the server and
// the removeAfterPeriod option is used; Otherwise, messages are deleted immediately
// and removeAfterPeriod is ignored.

// LeaveOnServer - If removeAfterPeriod is enabled, messages are deleted after specified period (in days).
type LeaveOnServer struct {
	Enabled           bool         `json:"enabled"`
	RemoveAfterPeriod OptionalLong `json:"removeAfterPeriod"`
}

type Pop3Account struct {
	Id              KId                `json:"id"`
	IsActive        bool               `json:"isActive"`
	Server          string             `json:"server"`   // POP3 server name
	UserName        string             `json:"userName"` // username on POP3 server
	Password        string             `json:"password"` // password appropriate to username
	Description     string             `json:"description"`
	DeliveryAddress string             `json:"deliveryAddress"`
	UseSortingRules bool               `json:"useSortingRules;//"` // If value is true sortType will save otherwise deliveryAddress will save. Default is false.
	SortType        string             `json:"sortType"`
	DropDuplicates  bool               `json:"dropDuplicates"` // drop duplicate messages?
	Mode            SslMode            `json:"mode"`
	Port            int                `json:"port"`
	Authentication  Pop3Authentication `json:"authentication"`
	MessageLimit    ByteValueWithUnits `json:"messageLimit"` // per session download limit - total message size
	MaxCount        int                `json:"maxCount"`     // per session download limit - maximum message count
	LeaveOnServer   LeaveOnServer      `json:"leaveOnServer"`
}

type Pop3AccountList []Pop3Account

type Pop3Sorting struct {
	Id          KId    `json:"id"`
	IsActive    bool   `json:"isActive"`
	SortAddress string `json:"sortAddress"`
	DeliverTo   string `json:"deliverTo"`
	Description string `json:"description"`
}

type Pop3SortingList []Pop3Sorting

type InternetConnection string

const (
	Permanent      InternetConnection = "Permanent"      // permanent Internet connection
	Triggered      InternetConnection = "Triggered"      // connection is established by scheduler
	TriggeredOnRas InternetConnection = "TriggeredOnRas" // Remote Access Service - Windows only option
)

type InternetSettings struct {
	Type                 InternetConnection `json:"type"`                 // type of Internet settings connection
	RasLine              string             `json:"rasLine"`              // name of RAS line
	UseSystemCredentials bool               `json:"useSystemCredentials"` // use username and password defined in system
	RasUser              string             `json:"rasUser"`              // RAS username
	RasPassword          string             `json:"rasPassword"`          // write only; password to RAS
	DialOnHigh           bool               `json:"dialOnHigh"`           // enable dial-up on high priority message
}

// DeliveryAddEtrnDownloadList - Add new ETRN downloads.
//	downloads - new ETRN download records
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryAddEtrnDownloadList(downloads EtrnDownloadList) (ErrorList, error) {
	params := struct {
		Downloads EtrnDownloadList `json:"downloads"`
	}{downloads}
	data, err := s.CallRaw("Delivery.addEtrnDownloadList", params)
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

// DeliveryAddPop3AccountList - Add new POP3 accounts.
//	accounts - new POP3 account records
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryAddPop3AccountList(accounts Pop3AccountList) (ErrorList, error) {
	params := struct {
		Accounts Pop3AccountList `json:"accounts"`
	}{accounts}
	data, err := s.CallRaw("Delivery.addPop3AccountList", params)
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

// DeliveryAddPop3SortingList - Add new POP3 sorting rules.
//	sortings - new POP3 sorting records
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryAddPop3SortingList(sortings Pop3SortingList) (ErrorList, error) {
	params := struct {
		Sortings Pop3SortingList `json:"sortings"`
	}{sortings}
	data, err := s.CallRaw("Delivery.addPop3SortingList", params)
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

// DeliveryAddScheduledActionList - Add scheduled actions.
//	actions - new scheduler actions
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryAddScheduledActionList(actions ScheduledActionList) (ErrorList, error) {
	params := struct {
		Actions ScheduledActionList `json:"actions"`
	}{actions}
	data, err := s.CallRaw("Delivery.addScheduledActionList", params)
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

// DeliveryDownloadEtrn - Start ETRN downloads.
func (s *ServerConnection) DeliveryDownloadEtrn() error {
	_, err := s.CallRaw("Delivery.downloadEtrn", nil)
	return err
}

// DeliveryGetEtrnDownloadList - Obtain list of ETRN download items.
//	query - query conditions and limits
// Return
//	list - ETRN download records
//  totalItems - amount of records for given search condition, useful when a limit is defined in the query
func (s *ServerConnection) DeliveryGetEtrnDownloadList(query SearchQuery) (EtrnDownloadList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Delivery.getEtrnDownloadList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       EtrnDownloadList `json:"list"`
			TotalItems int              `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DeliveryGetEtrnTimeout - Get timeout for ETRN reply on dial-up line
// Return
//	seconds - number of seconds for ETRN timeout
func (s *ServerConnection) DeliveryGetEtrnTimeout() (int, error) {
	data, err := s.CallRaw("Delivery.getEtrnTimeout", nil)
	if err != nil {
		return 0, err
	}
	seconds := struct {
		Result struct {
			Seconds int `json:"seconds"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &seconds)
	return seconds.Result.Seconds, err
}

// DeliveryGetInternetSettings - Obtain Internet connection settings.
// Return
//	settings - Internet connection settings
func (s *ServerConnection) DeliveryGetInternetSettings() (*InternetSettings, error) {
	data, err := s.CallRaw("Delivery.getInternetSettings", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings InternetSettings `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// DeliveryGetPop3AccountList - Obtain list of POP3 accounts.
//	query - query conditions and limits
// Return
//	list - POP3 accounts
//  totalItems - amount of accounts for given search condition, useful when a limit is defined in the query
func (s *ServerConnection) DeliveryGetPop3AccountList(query SearchQuery) (Pop3AccountList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Delivery.getPop3AccountList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       Pop3AccountList `json:"list"`
			TotalItems int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DeliveryGetPop3SortingList - Obtain list of POP3 sorting rules
//	query - query conditions and limits
// Return
//	list - POP3 sorting records
//  totalItems - amount of records for given search condition, useful when a limit is defined in the query
func (s *ServerConnection) DeliveryGetPop3SortingList(query SearchQuery) (Pop3SortingList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Delivery.getPop3SortingList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       Pop3SortingList `json:"list"`
			TotalItems int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DeliveryGetRasNames - Obtain Remote Access Service. Valid information available on Windows only.
// Return
//	names - list of available RAS names
func (s *ServerConnection) DeliveryGetRasNames() (StringList, error) {
	data, err := s.CallRaw("Delivery.getRasNames", nil)
	if err != nil {
		return nil, err
	}
	names := struct {
		Result struct {
			Names StringList `json:"names"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &names)
	return names.Result.Names, err
}

// DeliveryGetScheduledActionList - Obtain a list of scheduler actions.
//	query - query conditions and limits
// Return
//	list - scheduler actions
//  totalItems - amount of actions for given search condition, useful when limit is defined in query
func (s *ServerConnection) DeliveryGetScheduledActionList(query SearchQuery) (ScheduledActionList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Delivery.getScheduledActionList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ScheduledActionList `json:"list"`
			TotalItems int                 `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DeliveryRemoveEtrnDownloadList - Remove ETRN download items.
//	ids - identifier list of ETRN download records to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) DeliveryRemoveEtrnDownloadList(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Delivery.removeEtrnDownloadList", params)
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

// DeliveryRemovePop3AccountList - Remove POP3 accounts.
//	ids - identifier list of POP3 account records to be deleted
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryRemovePop3AccountList(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Delivery.removePop3AccountList", params)
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

// DeliveryRemovePop3SortingList - Remove POP3 sorting rules.
//	ids - identifier list of POP3 sorting records to be deleted
// Return
//	errors - list of error messages
func (s *ServerConnection) DeliveryRemovePop3SortingList(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Delivery.removePop3SortingList", params)
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

// DeliveryRemoveScheduledActionList - Remove scheduled actions.
//	ids - identifier list of scheduler actions to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) DeliveryRemoveScheduledActionList(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Delivery.removeScheduledActionList", params)
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

// DeliveryRunPop3Downloads - Proceed all POP3 downloads right now.
func (s *ServerConnection) DeliveryRunPop3Downloads() error {
	_, err := s.CallRaw("Delivery.runPop3Downloads", nil)
	return err
}

// DeliverySetEtrnDownload - Set 1 ETRN download item.
//	downloadId - updated ETRN download identifier
//	download - new ETRN download record
func (s *ServerConnection) DeliverySetEtrnDownload(downloadId KId, download EtrnDownload) error {
	params := struct {
		DownloadId KId          `json:"downloadId"`
		Download   EtrnDownload `json:"download"`
	}{downloadId, download}
	_, err := s.CallRaw("Delivery.setEtrnDownload", params)
	return err
}

// DeliverySetEtrnTimeout - Set timeout for ETRN reply on dial-up line.
//	seconds - number of seconds for ETRN timeout
func (s *ServerConnection) DeliverySetEtrnTimeout(seconds int) error {
	params := struct {
		Seconds int `json:"seconds"`
	}{seconds}
	_, err := s.CallRaw("Delivery.setEtrnTimeout", params)
	return err
}

// DeliverySetInternetSettings - Set Internet connection settings.
//	settings - Internet connection settings
func (s *ServerConnection) DeliverySetInternetSettings(settings InternetSettings) error {
	params := struct {
		Settings InternetSettings `json:"settings"`
	}{settings}
	_, err := s.CallRaw("Delivery.setInternetSettings", params)
	return err
}

// DeliverySetPop3Account - Set POP3 account.
//	accountId - updated POP3 account identifier
//	account - new POP3 account record
func (s *ServerConnection) DeliverySetPop3Account(accountId KId, account Pop3Account) error {
	params := struct {
		AccountId KId         `json:"accountId"`
		Account   Pop3Account `json:"account"`
	}{accountId, account}
	_, err := s.CallRaw("Delivery.setPop3Account", params)
	return err
}

// DeliverySetPop3Sorting - Set POP3 sorting rule.
//	sortingId - updated POP3 sorting identifier
//	sorting - new POP3 sorting record
func (s *ServerConnection) DeliverySetPop3Sorting(sortingId KId, sorting Pop3Sorting) error {
	params := struct {
		SortingId KId         `json:"sortingId"`
		Sorting   Pop3Sorting `json:"sorting"`
	}{sortingId, sorting}
	_, err := s.CallRaw("Delivery.setPop3Sorting", params)
	return err
}

// DeliverySetScheduledAction - Set a scheduled action.
//	actionId - updated action identifier
//	action - new scheduler actions
func (s *ServerConnection) DeliverySetScheduledAction(actionId KId, action ScheduledAction) error {
	params := struct {
		ActionId KId             `json:"actionId"`
		Action   ScheduledAction `json:"action"`
	}{actionId, action}
	_, err := s.CallRaw("Delivery.setScheduledAction", params)
	return err
}
