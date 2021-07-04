package connect

import "encoding/json"

// ConnectCertificateExportCertificate - Note: "export" is a keyword in C++, so name of the method must be changed: exportCertificate
//	id - ID of the certificate or certificate request
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) ConnectCertificateExportCertificate(id KId) (*Download, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("ConnectCertificate.exportCertificate", params)
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

// ConnectCertificateExportPrivateKey - Note: "export" is a keyword in C++, so the name of the method must be changed: exportPrivateKey
//	id - ID of the certificate or certificate request
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) ConnectCertificateExportPrivateKey(id KId) (*Download, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("ConnectCertificate.exportPrivateKey", params)
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

// ConnectCertificateGet - Obtain a list of certificates
// Return
//	certificates - current list of certificates
func (s *ServerConnection) ConnectCertificateGet() (CertificateList, error) {
	data, err := s.CallRaw("ConnectCertificate.get", nil)
	if err != nil {
		return nil, err
	}
	certificates := struct {
		Result struct {
			Certificates CertificateList `json:"certificates"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &certificates)
	return certificates.Result.Certificates, err
}

// ConnectCertificateGenerate - Generate a self-signed certificate
//	subject - information about subject
//	valid - length of the certificate's validity (in years, max value is 10)
// Return
//	id - ID of the new generated certificate
func (s *ServerConnection) ConnectCertificateGenerate(subject NamedValueList, valid int) (*KId, error) {
	params := struct {
		Subject NamedValueList `json:"subject"`
		Valid   int            `json:"valid"`
	}{subject, valid}
	data, err := s.CallRaw("ConnectCertificate.generate", params)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// ConnectCertificateGenerateRequest - Generate certificate request
//	subject - information about subject
// Return
//	id - ID of the new generated certificate request
func (s *ServerConnection) ConnectCertificateGenerateRequest(subject NamedValueList) (*KId, error) {
	params := struct {
		Subject NamedValueList `json:"subject"`
	}{subject}
	data, err := s.CallRaw("ConnectCertificate.generateRequest", params)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// ConnectCertificateImportCertificate - Import private key for the new certificate.
//	id - ID of private key or certificate request which belongs to the certificate
//	fileId - ID of the uploaded file
//	password - certificate password, if it is set (use empty string if password is not set)
func (s *ServerConnection) ConnectCertificateImportCertificate(id KId, fileId string, password string) error {
	params := struct {
		Id       KId    `json:"id"`
		FileId   string `json:"fileId"`
		Password string `json:"password"`
	}{id, fileId, password}
	_, err := s.CallRaw("ConnectCertificate.importCertificate", params)
	return err
}

// ConnectCertificateImportPrivateKey - Import private key for the new certificate.
//	fileId - ID of the uploaded file
// Return
//	needPassword - true if private key is encrypted with password
//	id - temporary ID to assign certificate to private key
func (s *ServerConnection) ConnectCertificateImportPrivateKey(fileId string) (bool, *KId, error) {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	data, err := s.CallRaw("ConnectCertificate.importPrivateKey", params)
	if err != nil {
		return false, nil, err
	}
	needPassword := struct {
		Result struct {
			NeedPassword bool `json:"needPassword"`
			Id           KId  `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &needPassword)
	return needPassword.Result.NeedPassword, &needPassword.Result.Id, err
}

// ConnectCertificateRemove - Remove list of certificate records
// Return
//	errors - error message list
func (s *ServerConnection) ConnectCertificateRemove(certificateIds KIdList) (ErrorList, error) {
	params := struct {
		CertificateIds KIdList `json:"certificateIds"`
	}{certificateIds}
	data, err := s.CallRaw("ConnectCertificate.remove", params)
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

// ConnectCertificateSetActive - Set active certificate
//	id - ID of the new active certificate
func (s *ServerConnection) ConnectCertificateSetActive(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("ConnectCertificate.setActive", params)
	return err
}

// ConnectCertificateToSource - Obtain source (plain-text representation) of the certificate
//	id - global identifier
// Return
//	source - certificate in plain text
func (s *ServerConnection) ConnectCertificateToSource(id KId) (string, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("ConnectCertificate.toSource", params)
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
