package connect

import "encoding/json"

type CompanyContact struct {
	Id       KId    `json:"id"`
	Name     string `json:"name"` // name of company contact (caption of item in list of contacts)
	Company  string `json:"company"`
	Street   string `json:"street"`
	Locality string `json:"locality"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Url      string `json:"url"`
	Phone    string `json:"phone"`
	Fax      string `json:"fax"`
	DomainId KId    `json:"domainId"` // id of domain associated with company contact
}

// CompanyContactList - List of company contacts
type CompanyContactList []CompanyContact

// Company contacts management

// CompanyContactsCreate - Create new company contacts.
// Return
//	errors - error message list
//	result - particular results for all items
func (s *ServerConnection) CompanyContactsCreate(companyContacts CompanyContactList) (ErrorList, CreateResultList, error) {
	params := struct {
		CompanyContacts CompanyContactList `json:"companyContacts"`
	}{companyContacts}
	data, err := s.CallRaw("CompanyContacts.create", params)
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

// CompanyContactsGet - Obtain a list of company contacts.
//	query - query conditions and limits
// Return
//	list - list of company contacts
//  totalItems - amount of company contacts for given search condition, useful when limit is defined in SearchQuery
func (s *ServerConnection) CompanyContactsGet(query SearchQuery) (CompanyContactList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("CompanyContacts.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       CompanyContactList `json:"list"`
			TotalItems int                `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// CompanyContactsGetAvailable - - Only company contacts for given domain and global company contacts are listed.
//	domainId - Only company contacts for given domain and global company contacts are listed.
// Return
//	companyContactList - list of user templates
func (s *ServerConnection) CompanyContactsGetAvailable(domainId KId) (CompanyContactList, error) {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	data, err := s.CallRaw("CompanyContacts.getAvailable", params)
	if err != nil {
		return nil, err
	}
	companyContactList := struct {
		Result struct {
			CompanyContactList CompanyContactList `json:"companyContactList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &companyContactList)
	return companyContactList.Result.CompanyContactList, err
}

// CompanyContactsRemove - Remove company contacts.
// Return
//	errors - error message list
func (s *ServerConnection) CompanyContactsRemove(companyContactsIds KIdList) (ErrorList, error) {
	params := struct {
		CompanyContactsIds KIdList `json:"companyContactsIds"`
	}{companyContactsIds}
	data, err := s.CallRaw("CompanyContacts.remove", params)
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

// CompanyContactsSet - Set existing company contacts to given pattern.
//	companyContactsIds - list of the company contacts's global identifier(s)
//	pattern - pattern to use for new values
// Return
//	errors - error message list
func (s *ServerConnection) CompanyContactsSet(companyContactsIds KIdList, pattern CompanyContact) (ErrorList, error) {
	params := struct {
		CompanyContactsIds KIdList        `json:"companyContactsIds"`
		Pattern            CompanyContact `json:"pattern"`
	}{companyContactsIds, pattern}
	data, err := s.CallRaw("CompanyContacts.set", params)
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
