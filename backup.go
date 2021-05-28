package connect

type LastBackupStatus string

const (
	backupStatusNone       LastBackupStatus = "backupStatusNone"
	backupStatusSuccessful LastBackupStatus = "backupStatusSuccessful"
	backupStatusFailed     LastBackupStatus = "backupStatusFailed"
)

type BackupType string

const (
	backupTypeFull         BackupType = "backupTypeFull"
	backupTypeDifferential BackupType = "backupTypeDifferential"
	backupTypeMirror       BackupType = "backupTypeMirror"
)

// BackupInfo - [READ-ONLY]
type BackupInfo struct {
	IsCreated bool `json:"isCreated"` // True, if backup was successfully created
	//Created UtcDateTime `json:"created"` // Time when backup started (always in GTM) TODO
	Size ByteValueWithUnits `json:"size"` // Compressed size of a backup
}

// BackupStatus - [READ-ONLY]
type BackupStatus struct {
	BackupInProgress bool             `json:"backupInProgress"` // True, if backup is in progress; otherwise, false
	Percents         int              `json:"percents"`         // Backup progress in percents (form 0 to 100)
	LastBackupStatus LastBackupStatus `json:"lastBackupStatus"` // Status of the last backup run
	ElapsedTime      Distance         `json:"elapsedTime"`      // Time from last started backup
	RemainingTime    Distance         `json:"remainingTime"`    // Approximated time to end of current backup
	LastFull         BackupInfo       `json:"lastFull"`         // Information about last full backup
	LastDifferential BackupInfo       `json:"lastDifferential"` // Information about last differential backup
	LastMirror       BackupInfo       `json:"lastMirror"`       // Information about last mirror backup
}

type BackupOptions struct {
	Paths                    Directories  `json:"paths"`                    // Paths to store/archive/backup, this field is used in both, archive and backup, options
	IsEnabled                bool         `json:"isEnabled"`                // Enable message store and configuration recovery backup
	Status                   BackupStatus `json:"status"`                   // Current backup status
	SplitSizeLimit           int          `json:"splitSizeLimit"`           // Split backup files if size reaches 'splitSizeLimit' (MB)
	RotationLimit            int          `json:"rotationLimit"`            // Keep at most 'rotationLimit' complete backups
	NetworkDiskUserName      string       `json:"networkDiskUserName"`      // If the backup directory is on the network disk, you may need to specify user name
	NetworkDiskPassword      string       `json:"networkDiskPassword"`      // ... and password
	NotificationEmailAddress string       `json:"notificationEmailAddress"` // An email address of person that will be notified when backup is completed or if any problems arise
}

type BackupSchedule struct {
	Id        KId        `json:"id"`        // [READ-ONLY]
	IsEnabled bool       `json:"isEnabled"` // True if backup schedule is enabled
	Type      BackupType `json:"type"`      // Backup type
	//Day Day `json:"day"` // Backup schedule day of week TODO
	Time        TimeHMS `json:"time"`        // Backup schedule start time - days are ignored!
	Description string  `json:"description"` // description of the backup schedule
}

type BackupScheduleList []BackupSchedule

//TODO Add Methods
