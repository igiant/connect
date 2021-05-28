package connect

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
