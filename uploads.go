package connect

import "encoding/json"

// Upload management

// UploadsRemove - Remove uploaded file.
//	id - identifier of uploaded file
func (s *ServerConnection) UploadsRemove(id string) error {
	params := struct {
		Id string `json:"id"`
	}{id}
	_, err := s.CallRaw("Uploads.remove", params)
	return err
}

// UploadsRemoveList - Remove uploaded files.
//	ids - identifiers of uploaded files
// Return
//	errors - list of errors
func (s *ServerConnection) UploadsRemoveList(ids StringList) (ErrorList, error) {
	params := struct {
		Ids StringList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Uploads.removeList", params)
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
