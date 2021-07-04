package connect

import "encoding/json"

// ValidType - Certificate Time properties info
type ValidType string

const (
	Valid       ValidType = "Valid"
	NotValidYet ValidType = "NotValidYet"
	ExpireSoon  ValidType = "ExpireSoon"
	Expired     ValidType = "Expired"
)

// ValidPeriod - Certificate Time properties
type ValidPeriod struct {
	ValidFromDate Date      `json:"validFromDate"` // @see SharedStructures.idl shared in lib
	ValidFromTime Time      `json:"validFromTime"` // @see SharedStructures.idl shared in lib
	ValidToDate   Date      `json:"validToDate"`   // @see SharedStructures.idl shared in lib
	ValidToTime   Time      `json:"validToTime"`   // @see SharedStructures.idl shared in lib
	ValidType     ValidType `json:"validType"`
}

type CertificateType string

const (
	ActiveCertificate   CertificateType = "ActiveCertificate"
	InactiveCertificate CertificateType = "InactiveCertificate"
	CertificateRequest  CertificateType = "CertificateRequest"
	Authority           CertificateType = "Authority"
	LocalAuthority      CertificateType = "LocalAuthority"
	BuiltInAuthority    CertificateType = "BuiltInAuthority"
	ServerCertificate   CertificateType = "ServerCertificate"
)

// Certificate properties
// issuer & subject valid names:
//  hostname;        // max 127 bytes
//  organizationName;    // max 127 bytes
//  organizationalUnitName; // max 127 bytes
//  city;          // max 127 bytes
//  state;          // max 127 bytes
//  country;         // ISO 3166 code
// Certificate -  emailAddress;      // max 255 bytes
type Certificate struct {
	Id                         KId                 `json:"id"`
	Status                     StoreStatus         `json:"status"`
	Name                       string              `json:"name"`
	Issuer                     NamedValueList      `json:"issuer"`
	Subject                    NamedValueList      `json:"subject"`
	SubjectAlternativeNameList NamedMultiValueList `json:"subjectAlternativeNameList"`
	Fingerprint                string              `json:"fingerprint"`       // 128-bit MD5, i.e. 16 hexa values separated by colons
	FingerprintSha1            string              `json:"fingerprintSha1"`   // 160-bit SHA1, i.e. 20 hexa values separated by colons
	FingerprintSha256          string              `json:"fingerprintSha256"` // 512-bit SHA256, i.e. 64 hexa values separated by colons
	ValidPeriod                ValidPeriod         `json:"validPeriod"`
	Valid                      bool                `json:"valid"` // exists and valid content
	Type                       CertificateType     `json:"type"`
	IsUntrusted                bool                `json:"isUntrusted"`
	VerificationMessage        string              `json:"verificationMessage"`
	ChainInfo                  StringList          `json:"chainInfo"`
	IsSelfSigned               bool                `json:"isSelfSigned"`
}

type CertificateList []Certificate

// Manager of Certificates

// CertificatesGet - Obtain a list of certificates
//	query - conditions and limits. Included from weblib.
// Return
//	certificates - current list of certificates
//  totalItems - count of all services on server (before the start/limit applied)
func (s *ServerConnection) CertificatesGet(query SearchQuery) (CertificateList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Certificates.get", params)
	if err != nil {
		return nil, 0, err
	}
	certificates := struct {
		Result struct {
			Certificates CertificateList `json:"certificates"`
			TotalItems   int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &certificates)
	return certificates.Result.Certificates, certificates.Result.TotalItems, err
}

// CertificatesSetName - Renames certificate
//	id - ID of certificate to rename
//	name - new name of the certificate
func (s *ServerConnection) CertificatesSetName(id KId, name string) error {
	params := struct {
		Id   KId    `json:"id"`
		Name string `json:"name"`
	}{id, name}
	_, err := s.CallRaw("Certificates.setName", params)
	return err
}

// CertificatesRemove - Remove list of certificate records
//	ids - list of identifiers of deleted user templates
// Return
//	errors - error message list
func (s *ServerConnection) CertificatesRemove(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Certificates.remove", params)
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

// CertificatesGenerate - Generate certificate.
//	subject - properties specified by user
//	name - name of the new certificate
//	certificateType - type of certificate to be generated, valid input is one of: InactiveCertificate/CertificateRequest/LocalAuthority
//	period - time properties specified by user, not relevant for CertificateRequest
// Return
//	id - ID of generated certificate
func (s *ServerConnection) CertificatesGenerate(subject NamedValueList, name string, certificateType CertificateType, period ValidPeriod) (KId, error) {
	params := struct {
		Subject NamedValueList  `json:"subject"`
		Name    string          `json:"name"`
		Type    CertificateType `json:"type"`
		Period  ValidPeriod     `json:"period"`
	}{subject, name, certificateType, period}
	data, err := s.CallRaw("Certificates.generate", params)
	if err != nil {
		return "", err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return id.Result.Id, err
}

// CertificatesGetCountryList - Get a list of countries.
// Return
//	countries - list of countries (name and ISO 3166 code)
func (s *ServerConnection) CertificatesGetCountryList() (NamedValueList, error) {
	data, err := s.CallRaw("Certificates.getCountryList", nil)
	if err != nil {
		return nil, err
	}
	countries := struct {
		Result struct {
			Countries NamedValueList `json:"countries"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &countries)
	return countries.Result.Countries, err
}

// CertificatesImportCertificate - Import certificate in PEM format
//	keyId - ID assigned to imported private key, @see importPrivateKey
//	fileId - id of uploaded file
//	name - name of the new certificate
//	certificateType - type of certificate to be imported, valid input is one of: InactiveCertificate/Authority/LocalAuthority
// Return
//	id - ID of generated certificate
func (s *ServerConnection) CertificatesImportCertificate(keyId KId, fileId string, name string, certificateType CertificateType) (KId, error) {
	params := struct {
		KeyId  KId             `json:"keyId"`
		FileId string          `json:"fileId"`
		Name   string          `json:"name"`
		Type   CertificateType `json:"type"`
	}{keyId, fileId, name, certificateType}
	data, err := s.CallRaw("Certificates.importCertificate", params)
	if err != nil {
		return "", err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return id.Result.Id, err
}

// CertificatesImportPrivateKey - Import private key. It generates ID, so it can be linked to Certificate content imported later, @see importCertificate
//	fileId - id of uploaded file
// Return
//	keyId - generated ID for new key
//	needPassword - true if private key is encrypted with password
func (s *ServerConnection) CertificatesImportPrivateKey(fileId string) (KId, bool, error) {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	data, err := s.CallRaw("Certificates.importPrivateKey", params)
	if err != nil {
		return "", false, err
	}
	keyId := struct {
		Result struct {
			KeyId        KId  `json:"keyId"`
			NeedPassword bool `json:"needPassword"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &keyId)
	return keyId.Result.KeyId, keyId.Result.NeedPassword, err
}

// CertificatesUnlockPrivateKey - Try to parse imported private key. Need to be called, when @importPrivateKey returns needPassword == true.
//	keyId - ID assigned to imported private key, @see importPrivateKey
//	password - certificate password
func (s *ServerConnection) CertificatesUnlockPrivateKey(keyId KId, password string) error {
	params := struct {
		KeyId    KId    `json:"keyId"`
		Password string `json:"password"`
	}{keyId, password}
	_, err := s.CallRaw("Certificates.unlockPrivateKey", params)
	return err
}

// CertificatesExportCertificate - Export of certificate or certificate request
//	id - ID of the certificate or certificate request
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) CertificatesExportCertificate(id KId) (*Download, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("Certificates.exportCertificate", params)
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

// CertificatesExportPrivateKey - Export of certificate or request privatekey
//	id - ID of the certificate or certificate request
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) CertificatesExportPrivateKey(id KId) (*Download, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("Certificates.exportPrivateKey", params)
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

// CertificatesToSource - Obtain source (plain-text representation) of the certificate
//	id - global identifier
// Return
//	source - certificate in plain text
func (s *ServerConnection) CertificatesToSource(id KId) (string, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("Certificates.toSource", params)
	if err != nil {
		return "", err
	}
	source := struct {
		Result struct {
			Source string `json:"source"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &source)
	return source.Result.Source, err
}
