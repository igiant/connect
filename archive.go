package connect

type DayWeekMonthPeriod string

const (
	periodDay   DayWeekMonthPeriod = "periodDay"
	periodWeek  DayWeekMonthPeriod = "periodWeek"
	periodMonth DayWeekMonthPeriod = "periodMonth"
)

type ArchiveOptions struct {
	Paths                     Directories        `json:"paths"`                     // Paths to store/archive/backup
	IsEnabled                 bool               `json:"isEnabled"`                 // Enable mail archiving
	RemoteArchive             OptionalString     `json:"remoteArchive"`             // Archive to remote email address
	ArchiveToLocalFolder      bool               `json:"archiveToLocalFolder"`      // Archive to local folder
	ArchiveFoldersInterval    DayWeekMonthPeriod `json:"archiveFoldersInterval"`    // Interval used for creating of new archive folders (in days/weeks/months)
	CompressOldArchiveFolders bool               `json:"compressOldArchiveFolders"` // Compress old archive folders
	CompressionStartTime      Time               `json:"compressionStartTime"`      // Time in the day when an archive compression shall start
	ArchiveLocalMessages      bool               `json:"archiveLocalMessages"`      // Local messages (local sender, local recipient)
	ArchiveIncomingMessages   bool               `json:"archiveIncomingMessages"`   // Incoming messages (remote sender, local recipient)
	ArchiveOutgoingMessages   bool               `json:"archiveOutgoingMessages"`   // Outgoing messages (local sender, remote recipient)
	ArchiveRelayedMessages    bool               `json:"archiveRelayedMessages"`    // Relayed messages (remote sender, remote recipient)
	ArchiveBeforeFilter       bool               `json:"archiveBeforeFilter"`       // Archive messages before content filter check (viruses and spams will be stored intact in the archive folders)
	IsXmppEnabled             bool               `json:"isXmppEnabled"`             // Enable archiving for instant messaging
	IsEnabledPerDomain        bool               `json:"isEnabledPerDomain"`        // Enable custom per domain settings
}

// TODO Add methods
