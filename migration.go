package connect

import "encoding/json"

type MigrationStatusEnum string

const (
	migNotStarted            MigrationStatusEnum = "migNotStarted"            // Migration planed but not started
	migCompressionStarted    MigrationStatusEnum = "migCompressionStarted"    // Migration started - packing of the mailbox and sending through network
	migCompressionFinished   MigrationStatusEnum = "migCompressionFinished"   // Source server actions finished
	migTransferStarted       MigrationStatusEnum = "migTransferStarted"       // Downloading packed mailbox
	migTransferFinished      MigrationStatusEnum = "migTransferFinished"      // Download completed
	migDecompressionStarted  MigrationStatusEnum = "migDecompressionStarted"  // Target server actions - unpacking of the mailbox
	migDecompressionFinished MigrationStatusEnum = "migDecompressionFinished" // Mailbox successfully decompressed
	migFinished              MigrationStatusEnum = "migFinished"              // Migration finished
	migCanceled              MigrationStatusEnum = "migCanceled"              // Migration canceled by user
	migError                 MigrationStatusEnum = "migError"                 // Migration failed - detailed error description can be found in logs
)

// MigrationStatus - Status of the migration task and progress of migration in percents.
type MigrationStatus struct {
	MigrationStatus    MigrationStatusEnum `json:"migrationStatus"`    // current status of the migration
	ProgressInPercents int                 `json:"progressInPercents"` // number from 0 to 100
	ErrorMessage       string              `json:"errorMessage"`       // If migrationStatus is migError, errorMessage optionaly contains detailed info about error in english
}

// This structure contains information about:
// - user which will be/is migrated
// - current homeserver (sourceServer)
// In get methods is returned also status of user's migration.

// MigrationTask - Whole structure is READ-ONLY
type MigrationTask struct {
	Id           KId        `json:"id"`           // id of migration task
	UserId       KId        `json:"userId"`       // id of migrated user
	UserName     string     `json:"userName"`     // name of migrated user
	SourceServer HomeServer `json:"sourceServer"` // source homeserver
	// HomeServer targetServer;   // in current implementation it's always current server -
	Status MigrationStatus `json:"status"` // status of the migration
}

// MigrationTaskList - List of migration tasks
type MigrationTaskList []MigrationTask

// MigrationCancel - Cancel planned or running migration tasks.
// Parameters
//	taskIdList - Identifiers of migration tasks which should be canceled
// Return
//	errors - error message list
func (c *ServerConnection) MigrationCancel(taskIdList KIdList) (ErrorList, error) {
	params := struct {
		TaskIdList KIdList `json:"taskIdList"`
	}{taskIdList}
	data, err := c.CallRaw("Migration.cancel", params)
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

// MigrationGet - Obtain list of migration tasks.
// Parameters
//	query - query attributes and limits
// Return
//	list - migration tasks
func (c *ServerConnection) MigrationGet(query SearchQuery) (MigrationTaskList, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Migration.get", params)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List MigrationTaskList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// MigrationGetCurrentHomeServer - Note: This method should be moved to DistributedDomain
// Return
//	homeServer - homeserver attributes
func (c *ServerConnection) MigrationGetCurrentHomeServer() (*HomeServer, error) {
	data, err := c.CallRaw("Migration.getCurrentHomeServer", nil)
	if err != nil {
		return nil, err
	}
	homeServer := struct {
		Result struct {
			HomeServer HomeServer `json:"homeServer"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &homeServer)
	return &homeServer.Result.HomeServer, err
}

// MigrationGetCurrentStatus - Obtain status of currently running migration task.
// Return
//	taskId - migration task identifier
//	status - migration task status
func (c *ServerConnection) MigrationGetCurrentStatus() (*KId, *MigrationStatus, error) {
	data, err := c.CallRaw("Migration.getCurrentStatus", nil)
	if err != nil {
		return nil, nil, err
	}
	taskId := struct {
		Result struct {
			TaskId KId             `json:"taskId"`
			Status MigrationStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &taskId)
	return &taskId.Result.TaskId, &taskId.Result.Status, err
}

// MigrationGetStatus - Obtain status of migration task specified by the task ID.
// Parameters
//	taskId - migration task identifier
// Return
//	status - migration task status
func (c *ServerConnection) MigrationGetStatus(taskId KId) (*MigrationStatus, error) {
	params := struct {
		TaskId KId `json:"taskId"`
	}{taskId}
	data, err := c.CallRaw("Migration.getStatus", params)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status MigrationStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// MigrationIsInProgress - Note: This method may fail if caller does not have full admin rights.
// Return
//	isInProgress - is there any migration task running?
func (c *ServerConnection) MigrationIsInProgress() (bool, error) {
	data, err := c.CallRaw("Migration.isInProgress", nil)
	if err != nil {
		return false, err
	}
	isInProgress := struct {
		Result struct {
			IsInProgress bool `json:"isInProgress"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &isInProgress)
	return isInProgress.Result.IsInProgress, err
}

// MigrationStart - Start a new migration task.
// Parameters
//	userIds - users to be migrated
// Return
//	errors - error message list
//	result
func (c *ServerConnection) MigrationStart(userIds KIdList) (ErrorList, CreateResultList, error) {
	params := struct {
		UserIds KIdList `json:"userIds"`
	}{userIds}
	data, err := c.CallRaw("Migration.start", params)
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
