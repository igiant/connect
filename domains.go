package connect

import "encoding/json"

type SearchQuery struct {
	Fields     []string          `json:"fields"`
	Conditions []string          `json:"conditions"`
	Start      int               `json:"start"`
	Limit      int               `json:"limit"`
	OrderBy    map[string]string `json:"orderBy,omitempty"`
}

type PlaceHolderList []Constant

type DomainInfo struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	IsPrimary            bool   `json:"isPrimary"`
	UserMaxCount         int    `json:"userMaxCount"`
	OutgoingMessageLimit `json:"outgoingMessageLimit"`
	KeepForRecovery      `json:"keepForRecovery"`
	AliasList            []interface{} `json:"aliasList"`
	ForwardingOptions    `json:"forwardingOptions"`
	KerberosRealm        string `json:"kerberosRealm"`
	WinNtName            string `json:"winNtName"`
	PamRealm             string `json:"pamRealm"`
	IPAddressBind        struct {
		Enabled bool   `json:"enabled"`
		Value   string `json:"value"`
	} `json:"ipAddressBind"`
	RenameInfo    `json:"renameInfo"`
	IsDistributed bool `json:"isDistributed"`
}

type DomainList []DomainInfo

type DomainSetting struct {
	Hostname               string `json:"hostname"`
	PublicFoldersPerDomain bool   `json:"publicFoldersPerDomain"`
	ServerID               string `json:"serverId"`
}

// DomainsGet obtains a list of domains.
func (c *Connection) DomainsGet(query SearchQuery) (DomainList, error) {
	q := struct {
		SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Domains.get", q)
	if err != nil {
		return nil, err
	}
	domainList := struct {
		Result struct {
			DomainList DomainList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &domainList)
	return domainList.Result.DomainList, err
}

// DomainsGetSettings obtains settings common in all domains.
func (c *Connection) DomainsGetSettings() (*DomainSetting, error) {
	data, err := c.CallRaw("Domains.getSettings", nil)
	if err != nil {
		return nil, err
	}
	domainSetting := struct {
		Result struct {
			DomainSetting `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &domainSetting)
	return &domainSetting.Result.DomainSetting, err
}

// DomainsCheckPublicFoldersIntegrity checks integrity of all public folders.
// If corrupted folder is found, try to fix it.
func (c *Connection) DomainsCheckPublicFoldersIntegrity() error {
	_, err := c.CallRaw("Domains.checkPublicFoldersIntegrity", nil)
	return err
}

// DomainsGetDomainFooterPlaceholders returns all supported placeholders for domain footer.
func (c *Connection) DomainsGetDomainFooterPlaceholders() (PlaceHolderList, error) {
	data, err := c.CallRaw("Domains.getDomainFooterPlaceholders", nil)
	if err != nil {
		return nil, err
	}
	placeHolderList := struct {
		Result struct {
			PlaceHolderList `json:"placeholders"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &placeHolderList)
	return placeHolderList.Result.PlaceHolderList, err
}

// DomainsSaveFooterImage - save a new footer's image.
func (c *Connection) DomainsSaveFooterImage(fileID string) (string, error) {
	params := struct {
		FileID string `json:"fileId"`
	}{fileID}
	data, err := c.CallRaw("Domains.saveFooterImage", &params)
	if err != nil {
		return "", err
	}
	imgUrl := struct {
		Result struct {
			ImgUrl string `json:"imgUrl"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &imgUrl)
	return imgUrl.Result.ImgUrl, err
}
