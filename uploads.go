package connect

import "encoding/json"

// Upload management

// UploadsRemove - Remove uploaded file.
// Parameters
//	id - identifier of uploaded file
func (c *ServerConnection) UploadsRemove(id string) error {
	params := struct {
		Id string `json:"id"`
	}{id}
	_, err := c.CallRaw("Uploads.remove", params)
	return err
}

// UploadsRemoveList - Remove uploaded files.
// Parameters
//	ids - identifiers of uploaded files
// Return
//	errors - list of errors
func (c *ServerConnection) UploadsRemoveList(ids StringList) (ErrorList, error) {
	params := struct {
		Ids StringList `json:"ids"`
	}{ids}
	data, err := c.CallRaw("Uploads.removeList", params)
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
