package connect

import "encoding/json"

// MlPermission - Mailing List (=ML) Action Access Right Type
type MlPermission string

const (
	Allowed   MlPermission = "Allowed"   // certain action is allowed
	Moderated MlPermission = "Moderated" // certain action must be approved by moderator
	Denied    MlPermission = "Denied"    // certain action is denied
)

// ModeratorPermission - Moderator Posting Access Right Type
type ModeratorPermission string

const (
	PostAllowed             ModeratorPermission = "PostAllowed"             // certain action is allowed
	PostModerated           ModeratorPermission = "PostModerated"           // certain action must be approved by moderator
	PostAccordingMembership ModeratorPermission = "PostAccordingMembership" // certain action is ruled according modarator membership
)

// MlReplyTo - ML Addressee Type
type MlReplyTo string

const (
	Sender         MlReplyTo = "Sender"         // email address of the sender
	ThisList       MlReplyTo = "ThisList"       // email address of the ML
	OtherAddress   MlReplyTo = "OtherAddress"   // the address of the original sender will be substituted by a user defined email address
	SenderThisList MlReplyTo = "SenderThisList" // Sender + ThisList
)

// MlMembership - Type of Mailing List Membership
type MlMembership string

const (
	Member    MlMembership = "Member"    // obtains contributions
	Moderator MlMembership = "Moderator" // can manipulate with members but does not obtain contributions
)

// UserOrEmail - ML member
type UserOrEmail struct {
	HasId        bool         `json:"hasId"`        // is a real user or email address only?
	UserId       KId          `json:"userId"`       // global user identification
	EmailAddress string       `json:"emailAddress"` // email address, filled only if hasId is false
	FullName     string       `json:"fullName"`     // fullName of user or associated email
	Kind         MlMembership `json:"kind"`         // a kind of membership
}

// MLMemberImportee - A ML member being imported from CSV file.
type MLMemberImportee struct {
	Member       UserOrEmail `json:"member"`       // ML member data
	IsImportable bool        `json:"isImportable"` // ML member can be imported
	Message      string      `json:"message"`      // error message if ML member is not importable
}

type MLMemberImporteeList []MLMemberImportee

// UserOrEmailList - List of ML members
type UserOrEmailList []UserOrEmail

// TrusteeKind - type indicator
type TrusteeKind string

const (
	TrusteeUser  TrusteeKind = "TrusteeUser"  // the type is the user
	TrusteeGroup TrusteeKind = "TrusteeGroup" // the type is the group
)

// Trustee - Entities that have access rights to read ML archive
type Trustee struct {
	Kind          TrusteeKind `json:"kind"`          // is user or group?
	ReaderId      KId         `json:"readerId"`      // group or user KId
	DisplayString string      `json:"displayString"` // login name or group name with domain name
	IsEnabled     bool        `json:"isEnabled"`     // true if user account is enabled
	ItemSource    DataSource  `json:"itemSource"`    // internal/LDAP
}

// TrusteeList - List of entities that have access rights to read ML archive
type TrusteeList []Trustee

// TrusteeTarget - Trustee target can be user or group
type TrusteeTarget struct {
	Id          KId         `json:"id"`          // unique identifier
	Type        TrusteeKind `json:"type"`        // is user or group?
	Name        string      `json:"name"`        // loginName for the User, name in square brackets for the Group
	FullName    string      `json:"fullName"`    // fullname for the User, empty string for the Group
	Description string      `json:"description"` // description of User/Group
	IsEnabled   bool        `json:"isEnabled"`   // is the User/Group enabled?
	ItemSource  DataSource  `json:"itemSource"`  // is the User/Group stored internally or by LDAP?
	HomeServer  HomeServer  `json:"homeServer"`  // id of users homeserver if server is in Cluster; groups haven't homeserver
}

// TrusteeTargetList - List of trustee targets
type TrusteeTargetList []TrusteeTarget

// SubscriptionPolicy - Rules for subscription
type SubscriptionPolicy struct {
	Type                  MlPermission `json:"type"`
	ModeratorReview       bool         `json:"moderatorReview"`
	ModeratorNotification bool         `json:"moderatorNotification"`
}

// PostingPolicy - Rules for posting
type PostingPolicy struct {
	MemberPosting           MlPermission        `json:"memberPosting"`           // posting policy for ML member(s)
	NonMemberPosting        MlPermission        `json:"nonMemberPosting"`        // posting policy for ML non-member(s)
	ModeratorPosting        ModeratorPermission `json:"moderatorPosting"`        // posting policy for ML moderator(s)
	UserPostingNotification bool                `json:"userPostingNotification"` // notify user that the posting will be reviewed by a moderator
	SendErrorsToModerator   bool                `json:"sendErrorsToModerator"`   // send delivery errors to moderator(s)
}

// ArchiveSettings - How is the archive organized?
type ArchiveSettings struct {
	KeepArchive          bool        `json:"keepArchive"`          // maintain archive
	ArchiveOnlyForLogged bool        `json:"archiveOnlyForLogged"` // the archive is available for logged users only
	ArchiveReaderList    TrusteeList `json:"archiveReaderList"`    // list of archive readers, can be either user or group, meaningful only if onlyForLogged is true
}

// Ml - Mailing List Structure
type Ml struct {
	Id                KId                `json:"id"`                // [READ-ONLY] global identification of ML
	DomainId          KId                `json:"domainId"`          // [REQUIRED FOR CREATE] [WRITE-ONCE] identification in which domain ML exists
	Name              string             `json:"name"`              // [REQUIRED FOR CREATE] [WRITE-ONCE] ML name, name@domain is email address
	Description       string             `json:"description"`       // description
	LanguageId        KId                `json:"languageId"`        // language to be spoken withing mailing list
	WelcomeString     string             `json:"welcomeString"`     // string to be sent as welcome of a new memeber
	FooterString      string             `json:"footerString"`      // string to be sent as footer of each contribution
	Subscription      SubscriptionPolicy `json:"subscription"`      // type of ML subscription policy
	Posting           PostingPolicy      `json:"posting"`           // type of ML posting policy
	ReplyTo           MlReplyTo          `json:"replyTo"`           // how should be replied to
	OtherAddress      string             `json:"otherAddress"`      // if replyTo is OtherAddress, it contains email address
	SubjectPrefix     string             `json:"subjectPrefix"`     // prefix for each subject
	HideSenderAddress bool               `json:"hideSenderAddress"` // replace sender's email address by ML address
	AllowEmptySubject bool               `json:"allowEmptySubject"` // allow posting with empty subject
	Archive           ArchiveSettings    `json:"archive"`           // archive settings
	MembersCount      int                `json:"membersCount"`      // [READ-ONLY] Number of members.
	HomeServer        HomeServer         `json:"homeServer"`        // [READ-ONLY] Id of users homeserver if server is in Cluster
}

// MlList - List of mailing lists
type MlList []Ml

// MailingListsAddMlUserList - Add one or more members/moderators to a mailing list.
// Parameters
//	members - ML members and/or moderators
//	mlId - unique ML identifier
// Return
//	errors - appropriate error messages
func (s *ServerConnection) MailingListsAddMlUserList(members UserOrEmailList, mlId KId) (ErrorList, error) {
	params := struct {
		Members UserOrEmailList `json:"members"`
		MlId    KId             `json:"mlId"`
	}{members, mlId}
	data, err := s.CallRaw("MailingLists.addMlUserList", params)
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

// MailingListsCreate - Create new mailing lists.
// Parameters
//	mailingLists - mailing list entities
// Return
//	errors - error message list
//	result - list of IDs of created mailing lists
func (s *ServerConnection) MailingListsCreate(mailingLists MlList) (ErrorList, CreateResultList, error) {
	params := struct {
		MailingLists MlList `json:"mailingLists"`
	}{mailingLists}
	data, err := s.CallRaw("MailingLists.create", params)
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

// MailingListsExportMlUsersToCsv - Export of mailing list users of specified membership type.
// Parameters
//	kind - membership type
//	mlId - unique ML identifier
// Return
//	fileDownload - description of output file
func (s *ServerConnection) MailingListsExportMlUsersToCsv(kind MlMembership, mlId KId) (*Download, error) {
	params := struct {
		Kind MlMembership `json:"kind"`
		MlId KId          `json:"mlId"`
	}{kind, mlId}
	data, err := s.CallRaw("MailingLists.exportMlUsersToCsv", params)
	if err != nil {
		return nil, err
	}
	fileDownload := struct {
		Result struct {
			FileDownload Download `json:"fileDownload"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileDownload)
	return &fileDownload.Result.FileDownload, err
}

// MailingListsGet - Obtain a list of mailing lists.
// Parameters
//	query - query conditions and limits
// Return
//	list - mailing lists
//  totalItems - amount of MLs for given search condition, useful when a limit is defined in SearchQuery
func (s *ServerConnection) MailingListsGet(query SearchQuery, domainId KId) (MlList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("MailingLists.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       MlList `json:"list"`
			TotalItems int    `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// MailingListsGetMlUserList - Obtain list of mailing list users including membership type.
// Parameters
//	query - orderBy definition (conditions and limit are ignored)
//	mlId - unique ML identifier
// Return
//	list - mailing list members and/or moderators
//  totalItems - amount of MLs members for given search condition, useful when a limit is defined in search query
func (s *ServerConnection) MailingListsGetMlUserList(query SearchQuery, mlId KId) (UserOrEmailList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
		MlId  KId         `json:"mlId"`
	}{query, mlId}
	data, err := s.CallRaw("MailingLists.getMlUserList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       UserOrEmailList `json:"list"`
			TotalItems int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// MailingListsGetMlUserListFromCsv - Parse CSV file in format 'Email, FullName' and return list of members.
// Parameters
//	fileId - ID of the uploaded file
//	mlToImport - unique ML identifier or empty string if XML does not exist yet
// Return
//	members - ML members and/or moderators
func (s *ServerConnection) MailingListsGetMlUserListFromCsv(fileId string, mlToImport KId) (MLMemberImporteeList, error) {
	params := struct {
		FileId     string `json:"fileId"`
		MlToImport KId    `json:"mlToImport"`
	}{fileId, mlToImport}
	data, err := s.CallRaw("MailingLists.getMlUserListFromCsv", params)
	if err != nil {
		return nil, err
	}
	members := struct {
		Result struct {
			Members MLMemberImporteeList `json:"members"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &members)
	return members.Result.Members, err
}

// MailingListsGetSuffixes - processing of special commands of mailing list.
// Return
//	suffixes - list of suffixes
func (s *ServerConnection) MailingListsGetSuffixes() (StringList, error) {
	data, err := s.CallRaw("MailingLists.getSuffixes", nil)
	if err != nil {
		return nil, err
	}
	suffixes := struct {
		Result struct {
			Suffixes StringList `json:"suffixes"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &suffixes)
	return suffixes.Result.Suffixes, err
}

// MailingListsGetTrusteeTargetList - Obtain a list of potential mailing list archive rights targets.
// Parameters
//	query - query attributes and limits
// Return
//	list - trustee targets
//  totalItems - amount of trustee targets, useful when a limit is defined in SearchQuery
func (s *ServerConnection) MailingListsGetTrusteeTargetList(query SearchQuery, domainId KId) (TrusteeTargetList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("MailingLists.getTrusteeTargetList", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       TrusteeTargetList `json:"list"`
			TotalItems int               `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// MailingListsRemove - Remove mailing lists.
// Parameters
//	mlIds - list of global identifiers of MLs to be deleted
// Return
//	errors - error message list
func (s *ServerConnection) MailingListsRemove(mlIds KIdList) (ErrorList, error) {
	params := struct {
		MlIds KIdList `json:"mlIds"`
	}{mlIds}
	data, err := s.CallRaw("MailingLists.remove", params)
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

// MailingListsRemoveMlUserList - Remove member(s)/moderator(s) from a mailing list.
// Parameters
//	members - ML members and/or moderators
//	mlId - unique ML identifier
// Return
//	errors - appropriate error messages
func (s *ServerConnection) MailingListsRemoveMlUserList(members UserOrEmailList, mlId KId) (ErrorList, error) {
	params := struct {
		Members UserOrEmailList `json:"members"`
		MlId    KId             `json:"mlId"`
	}{members, mlId}
	data, err := s.CallRaw("MailingLists.removeMlUserList", params)
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

// MailingListsSet - Create a new mailing list.
// Parameters
//	mlIds - ML global identifiers
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) MailingListsSet(mlIds KIdList, pattern Ml) (ErrorList, error) {
	params := struct {
		MlIds   KIdList `json:"mlIds"`
		Pattern Ml      `json:"pattern"`
	}{mlIds, pattern}
	data, err := s.CallRaw("MailingLists.set", params)
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
