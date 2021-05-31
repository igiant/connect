package connect

import "encoding/json"

type BuildType string

const (
	Alpha BuildType = "Alpha"
	Beta  BuildType = "Beta"
	Rc    BuildType = "Rc"
	Final BuildType = "Final"
	Patch BuildType = "Patch"
)

// ProductVersion - Product version, machine- and human-readable.
type ProductVersion struct {
	ProductName   string    `json:"productName"`   // e.g. "Kerio Connect"
	Major         int       `json:"major"`         // e.g. 7
	Minor         int       `json:"minor"`         // e.g. 4
	Revision      int       `json:"revision"`      // e.g. 0
	Build         int       `json:"build"`         // e.g. 4528
	Order         int       `json:"order"`         // e.g. 1 for alpha/beta/rc/patch 1
	ReleaseType   BuildType `json:"releaseType"`   // e.g. Patch
	DisplayNumber string    `json:"displayNumber"` // e.g. "7.4.0 patch 1"
}

// Number, bumped with each API change. E.g. version 7.3.1 has a

// ApiVersion - higher or equal number as version 7.3.0 - equal if there's no change in API.
type ApiVersion int

// Informs about product version and API version.

// VersionGetProductVersion - Get product version.
func (s *ServerConnection) VersionGetProductVersion() (*ProductVersion, error) {
	data, err := s.CallRaw("Version.getProductVersion", nil)
	if err != nil {
		return nil, err
	}
	productVersion := struct {
		Result struct {
			ProductVersion ProductVersion `json:"productVersion"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &productVersion)
	return &productVersion.Result.ProductVersion, err
}

// VersionGetApiVersion - Get version of Administration API.
func (s *ServerConnection) VersionGetApiVersion() (*ApiVersion, error) {
	data, err := s.CallRaw("Version.getApiVersion", nil)
	if err != nil {
		return nil, err
	}
	apiVersion := struct {
		Result struct {
			ApiVersion ApiVersion `json:"apiVersion"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &apiVersion)
	return &apiVersion.Result.ApiVersion, err
}
