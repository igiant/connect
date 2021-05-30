package connect

import "encoding/json"

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

// Archive

// ArchiveGet - Obtain archive options.
// Return
//	options - current archive options
func (c *ServerConnection) ArchiveGet() (*ArchiveOptions, error) {
	data, err := c.CallRaw("Archive.get", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options ArchiveOptions `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return &options.Result.Options, err
}

// ArchiveSet - Set archive options.
// Parameters
//	options - archive options
func (c *ServerConnection) ArchiveSet(options ArchiveOptions) error {
	params := struct {
		Options ArchiveOptions `json:"options"`
	}{options}
	_, err := c.CallRaw("Archive.set", params)
	return err
}

// ArchiveGetXmppArchiveFiles - Returns links to available Instant Messaging archive files
func (c *ServerConnection) ArchiveGetXmppArchiveFiles() (DownloadList, error) {
	data, err := c.CallRaw("Archive.getXmppArchiveFiles", nil)
	if err != nil {
		return nil, err
	}
	fileList := struct {
		Result struct {
			FileList DownloadList `json:"fileList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileList)
	return fileList.Result.FileList, err
}

// ArchiveGetImArchiveFile - Returns link to IM archive file in given period
func (c *ServerConnection) ArchiveGetImArchiveFile(fromDate Date, toDate Date) (*Download, error) {
	params := struct {
		FromDate Date `json:"fromDate"`
		ToDate   Date `json:"toDate"`
	}{fromDate, toDate}
	data, err := c.CallRaw("Archive.getImArchiveFile", params)
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
