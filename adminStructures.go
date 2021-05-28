package connect

// ActionAfterDays Message clean out setting Note: all fields must be assigned if used in set methods
type ActionAfterDays struct {
	IsEnabled bool `json:"isEnabled"` // is action on/off?
	Days      int  `json:"days"`      // after how many days is an action performed?
}

// Distance Note: all fields must be assigned if used in set methods
type Distance struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type TimeHMS Distance

type DistanceOrNull struct {
	Type     string   `json:"type"`
	TimeSpan Distance `json:"timeSpan"`
}

type Directories struct {
	StorePath   string `json:"storePath"`   // Path to the store directory.
	ArchivePath string `json:"archivePath"` // Path to the archive directory.
	BackupPath  string `json:"backupPath"`  // Path to the backup directory.
}
