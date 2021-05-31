package connect

import "encoding/json"

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
	IsCreated bool               `json:"isCreated"` // True, if backup was successfully created
	Created   UtcDateTime        `json:"created"`   // Time when backup started (always in GTM)
	Size      ByteValueWithUnits `json:"size"`      // Compressed size of a backup
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
	Id          KId        `json:"id"`          // [READ-ONLY]
	IsEnabled   bool       `json:"isEnabled"`   // True if backup schedule is enabled
	Type        BackupType `json:"type"`        // Backup type
	DayType     DayType    `json:"day"`         // Backup schedule day of week
	Time        TimeHMS    `json:"time"`        // Backup schedule start time - days are ignored!
	Description string     `json:"description"` // description of the backup schedule
}

type BackupScheduleList []BackupSchedule

// Backup

// BackupGet - Obtain backup options.
// Return
//	options - current backup options
func (c *ServerConnection) BackupGet() (*BackupOptions, error) {
	data, err := c.CallRaw("Backup.get", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options BackupOptions `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// BackupGetScheduleList - Obtain list of backup scheduling.
// Parameters
//	query - order by, limits
// Return
//	scheduleList
func (c *ServerConnection) BackupGetScheduleList(query SearchQuery) (BackupScheduleList, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Backup.getScheduleList", params)
	if err != nil {
		return nil, err
	}
	scheduleList := struct {
		Result struct {
			ScheduleList BackupScheduleList `json:"scheduleList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &scheduleList)
	return scheduleList.Result.ScheduleList, err
}

// BackupGetStatus - Return current backup status.
// Return
//	status - backup status
func (c *ServerConnection) BackupGetStatus() (*BackupStatus, error) {
	data, err := c.CallRaw("Backup.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status BackupStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// BackupSet - Set backup options.
// Parameters
//	options - backup options
func (c *ServerConnection) BackupSet(options BackupOptions) error {
	params := struct {
		Options BackupOptions `json:"options"`
	}{options}
	_, err := c.CallRaw("Backup.set", params)
	return err
}

// BackupSetScheduleList - Set all backup schedules.
// Parameters
//	scheduleList
func (c *ServerConnection) BackupSetScheduleList(scheduleList BackupScheduleList) error {
	params := struct {
		ScheduleList BackupScheduleList `json:"scheduleList"`
	}{scheduleList}
	_, err := c.CallRaw("Backup.setScheduleList", params)
	return err
}

// BackupStart - Start backup according to current settings.
// Parameters
//	backupType - backup type
func (c *ServerConnection) BackupStart(backupType BackupType) error {
	params := struct {
		BackupType BackupType `json:"backupType"`
	}{backupType}
	_, err := c.CallRaw("Backup.start", params)
	return err
}
