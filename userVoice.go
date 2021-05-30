package connect

import "encoding/json"

// UserVoiceSettings - Settings of UserVoice
type UserVoiceSettings struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserVoice accounts management

// UserVoiceGet - Obtain settings of User Voice.
// Return
//	settings - structure with settings
func (c *ServerConnection) UserVoiceGet() (*UserVoiceSettings, error) {
	data, err := c.CallRaw("UserVoice.get", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings UserVoiceSettings `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// UserVoiceGetStatus - Get status of registration of user voice.
// Return
//	isSet - true if user voice is set
func (c *ServerConnection) UserVoiceGetStatus() (bool, error) {
	data, err := c.CallRaw("UserVoice.getStatus", nil)
	if err != nil {
		return false, err
	}
	isSet := struct {
		Result struct {
			IsSet bool `json:"isSet"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &isSet)
	return isSet.Result.IsSet, err
}

// UserVoiceGetUrl - Parameters name and email can be empty strings.
// Parameters
//	name - user displayname for userVoice
//	email - user email address for userVoice
// Return
//	url - URL to userVoice with single sign on token
func (c *ServerConnection) UserVoiceGetUrl(name string, email string) (string, error) {
	params := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{name, email}
	data, err := c.CallRaw("UserVoice.getUrl", params)
	if err != nil {
		return "", err
	}
	url := struct {
		Result struct {
			Url string `json:"url"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &url)
	return url.Result.Url, err
}

// UserVoiceSet - Set settings of User Voice.
// Parameters
//	settings - structure with settings
func (c *ServerConnection) UserVoiceSet(settings UserVoiceSettings) error {
	params := struct {
		Settings UserVoiceSettings `json:"settings"`
	}{settings}
	_, err := c.CallRaw("UserVoice.set", params)
	return err
}
