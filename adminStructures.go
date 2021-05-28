package connect

// Type of authorization
type UserAuthType string

const (
	UInternalAuth  UserAuthType = "UInternalAuth"  // Internal authorization
	UWindowsNTAuth UserAuthType = "UWindowsNTAuth" // Windows NT domain authorization
	UPamAuth       UserAuthType = "UPamAuth"       // Authorization for linux
	UKerberosAuth  UserAuthType = "UKerberosAuth"  // Kerberos authorization
	UAppleAuth     UserAuthType = "UAppleAuth"     // Apple authorization
	ULDAPAuth      UserAuthType = "ULDAPAuth"      // LDAP authorization
)

// Type for returning authorization type supported by server
type AuthTypeList []UserAuthType

// Type of deleting a user
type DirectoryServiceDeleteMode string

const (
	DSModeDeactivate DirectoryServiceDeleteMode = "DSModeDeactivate" // User is deactivated but not deleted
	DSModeDelete     DirectoryServiceDeleteMode = "DSModeDelete"     // User is deleted
)

//  Enum type for determine which web component is in use in session
type WebComponent string

const (
	WebComponentWEBMAIL WebComponent = "WebComponentWEBMAIL" // WebMail
	WebComponentADMIN   WebComponent = "WebComponentADMIN"   // Web Administration
	WebComponentMINI    WebComponent = "WebComponentMINI"    // WebMail mini
)

// Message clean out setting
// Note: all fields must be assigned if used in set methods
type ActionAfterDays struct {
	IsEnabled bool `json:"isEnabled"` // is action on/off?
	Days      int  `json:"days"`      // after how many days is an action performed?
}

// Note: all fields must be assigned if used in set methods
type Distance struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type DistanceType string

const (
	dtNull     DistanceType = "dtNull"
	dtTimeSpan DistanceType = "dtTimeSpan"
)

type DistanceOrNull struct {
	Type     DistanceType `json:"type"`
	TimeSpan Distance     `json:"timeSpan"`
}

type Directories struct {
	StorePath   string `json:"storePath"`   // Path to the store directory
	ArchivePath string `json:"archivePath"` // Path to the archive directory
	BackupPath  string `json:"backupPath"`  // Path to the backup directory
}

// source of user data
type DataSource string

const (
	DSInternalSource DataSource = "DSInternalSource" // internal source of user data
	DSLDAPSource     DataSource = "DSLDAPSource"     // LDAP source of user data
)

//
// Note: all fields must be assigned if used in set methods
type TimeHMS struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

type TimeUnit string

const (
	Minutes TimeUnit = "Minutes"
	Hours   TimeUnit = "Hours"
	Days    TimeUnit = "Days"
	Weeks   TimeUnit = "Weeks"
)

//
// Note: all fields must be assigned if used in set methods
type TimeLimit struct {
	Value int      `json:"value"` // how many
	Units TimeUnit `json:"units"` // in which units
}

// A way how to say client that the server has a constant
type NamedConstant struct {
	Name  string `json:"name"`  // constant name
	Value string `json:"value"` // a value of constant
}

type NamedConstantList []NamedConstant

// Except getAboutInfo() all methods are available for non-authenticated users
type DirectoryAccessResult string

const (
	directoryExists            DirectoryAccessResult = "directoryExists"            // Directory exist, read/write allowed
	directoryDoesNotExist      DirectoryAccessResult = "directoryDoesNotExist"      // Directory does not exist (or unable to create)
	directoryExistAccessDenied DirectoryAccessResult = "directoryExistAccessDenied" // Directory exist, read or write permission not granted
	directoryUnaccessible      DirectoryAccessResult = "directoryUnaccessible"      // Unable to connect network device
)

// Information about directory
type Directory struct {
	Name            string `json:"name"`
	HasSubdirectory bool   `json:"hasSubdirectory"`
}

// List of restrictions
type DirectoryList []Directory

type BuildType string

const (
	Alpha BuildType = "Alpha"
	Beta  BuildType = "Beta"
	Rc    BuildType = "Rc"
	Final BuildType = "Final"
	Patch BuildType = "Patch"
)

type DeployedType string

const (
	DeployedStandalone DeployedType = "DeployedStandalone" // Normal instalation
	DeployedCloud      DeployedType = "DeployedCloud"      // Kerio Connect is running in a cloud
	DeployedKerioVA    DeployedType = "DeployedKerioVA"    // Kerio Connect VMWare Virtual Appliance
)

// Operating System Family
type ServerOs string

const (
	Windows ServerOs = "Windows"
	MacOs   ServerOs = "MacOs"
	Linux   ServerOs = "Linux"
)

type UpdateCheckerStatus string

const (
	updNoUpdate   UpdateCheckerStatus = "updNoUpdate"   // Update status: No update
	updNewVersion UpdateCheckerStatus = "updNewVersion" // Update status: New version
	updError      UpdateCheckerStatus = "updError"      // Update status: Error
)

type UpdateInfo struct {
	Result      UpdateCheckerStatus `json:"result"`
	Description string              `json:"description"`
	DownloadUrl string              `json:"downloadUrl"`
	InfoUrl     string              `json:"infoUrl"`
}

type ProductInfo struct {
	ProductName  string       `json:"productName"`
	Version      string       `json:"version"`
	BuildNumber  string       `json:"buildNumber"`
	OsName       string       `json:"osName"`
	Os           ServerOs     `json:"os"`
	ReleaseType  BuildType    `json:"releaseType"`
	DeployedType DeployedType `json:"deployedType"`
	UpdateInfo   UpdateInfo   `json:"updateInfo"`
}

type CustomImage struct {
	IsEnabled bool   `json:"isEnabled"` // Is used
	Url       string `json:"url"`       // [READ ONLY]
	Id        string `json:"id"`        // [WRITE ONCE] Id of uploaded image.
}

type NotificationType string

const (
	notifyOnce  NotificationType = "notifyOnce"
	notifyEvery NotificationType = "notifyEvery"
)

type QuotaNotification struct {
	Type   NotificationType `json:"type"`
	Period TimeLimit        `json:"period"`
}
