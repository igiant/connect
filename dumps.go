package connect

import "encoding/json"

type CoreDump struct {
	Size      ByteValueWithUnits `json:"size"`
	Timestamp DateTimeStamp      `json:"timestamp"`
}

type DumpList []CoreDump

// DumpsGet - Obtain list of available crash dumps
// Return
//	dumps - list of all available crash dumps
func (c *Connection) DumpsGet() (DumpList, error) {
	data, err := c.CallRaw("Dumps.get", nil)
	if err != nil {
		return nil, err
	}
	dumps := struct {
		Result struct {
			Dumps DumpList `json:"dumps"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dumps)
	return dumps.Result.Dumps, err
}

// DumpsRemove - Remove all crash dumps from server disk
func (c *Connection) DumpsRemove() error {
	_, err := c.CallRaw("Dumps.remove", nil)
	return err
}

// DumpsSend - Upload last available crash dump to Kerio.
// Parameters
//	description - plain text information to be sent with crash dump
//	email - contact information to be sent with crash dump
func (c *Connection) DumpsSend(description string, email string) error {
	params := struct {
		Description string `json:"description"`
		Email       string `json:"email"`
	}{description, email}
	_, err := c.CallRaw("Dumps.send", params)
	return err
}
