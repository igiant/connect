package connect

import "encoding/json"

// HomeServer - User's home server in a distributed domain.
type HomeServer struct {
	Id   KId    `json:"id"`   // server's id
	Name string `json:"name"` // server's Internet hostname
}

type HomeServerList []HomeServer

type ClusterErrorType string

const (
	clSuccess                 ClusterErrorType = "clSuccess"
	clError                   ClusterErrorType = "clError"                   // Generic cluster error, see ClusterError.errorMessage for details
	clSelfConnectError        ClusterErrorType = "clSelfConnectError"        // The master cannot be the same as slave
	clConnectToSlaveError     ClusterErrorType = "clConnectToSlaveError"     // Connection to slave is not allowed
	clInaccessibleHost        ClusterErrorType = "clInaccessibleHost"        // Cannot connect to the specified host
	clInvalidUserOrPassword   ClusterErrorType = "clInvalidUserOrPassword"   // User name or password are invalid or has insufficient rights
	clIncorrectClusterVersion ClusterErrorType = "clIncorrectClusterVersion" // Remote server has incompatible implementation of cluster services
	clDataConflict            ClusterErrorType = "clDataConflict"            // There are multiple resources/aliases or mailing lists with the same name, server cannot be connected to cluster
	clDirServiceRemoteEmpty   ClusterErrorType = "clDirServiceRemoteEmpty"   // Specified distributed domain has no Directory Service configured on remote distributed domain host
	clDirServiceLocalEmpty    ClusterErrorType = "clDirServiceLocalEmpty"    // Specified distributed domain has no Directory Service configured on local distributed domain host
	clDirServiceDifferent     ClusterErrorType = "clDirServiceDifferent"     // Specified distributed domain has different Directory Service configured on local and remote distributed domain host
)

type ClusterConflictTarget string

const (
	clResource    ClusterConflictTarget = "clResource"
	clAlias       ClusterConflictTarget = "clAlias"
	clMailingList ClusterConflictTarget = "clMailingList"
	clDomainAlias ClusterConflictTarget = "clDomainAlias"
	clDomain      ClusterConflictTarget = "clDomain"
)

type ClusterConflict struct {
	Type       ClusterConflictTarget `json:"type"`
	Name       string                `json:"name"`
	Domain     string                `json:"domain"`
	HomeServer string                `json:"homeServer"`
}

type ClusterConflictList []ClusterConflict

type ClusterError struct {
	Type         ClusterErrorType    `json:"type"`
	ErrorMessage LocalizableMessage  `json:"errorMessage"` // is assigned if type is clError
	ConflictList ClusterConflictList `json:"conflictList"` // List of Resources/Aliases/MLists which are already defined in cluster. The conflictList is empty if type is different from dataConflict.
}

// ClusterRole - Role of the server in cluster
type ClusterRole string

const (
	clStandalone ClusterRole = "clStandalone"
	clMaster     ClusterRole = "clMaster"
	clSlave      ClusterRole = "clSlave"
)

type ClusterStatus string

const (
	csReady ClusterStatus = "csReady" // Server in claster work well.
	csError ClusterStatus = "csError" // Server in claster don't work with some error, see errorMessages for details
)

type LocalizableMessageList []LocalizableMessage

type ClusterDomainStatus string

const (
	csDomainNotChecked   ClusterDomainStatus = "csDomainNotChecked"
	csDomainExists       ClusterDomainStatus = "csDomainExists"
	csDomainDoesNotExist ClusterDomainStatus = "csDomainDoesNotExist"
)

type ClusterServer struct {
	Hostname      string                 `json:"hostname"`
	IsPrimary     bool                   `json:"isPrimary"`
	IsLocal       bool                   `json:"isLocal"`
	Status        ClusterStatus          `json:"status"`
	ErrorMessages LocalizableMessageList `json:"errorMessages"` // is assigned if type is clError
	DomainStatus  ClusterDomainStatus    `json:"domainStatus"`
}

type ClusterAuthentication struct {
	HostName  string `json:"hostName"`
	AdminUser string `json:"adminUser"`
	Password  string `json:"password"`
}

type ClusterServerList []ClusterServer

// DistributedDomainConnect - Connect server to cluster as slave.
// Parameters
//	hostName - name of the master server
//	adminUser - username of administrator on the master server
//	password - administrator's password
// Return
//	result - if ClusterErrorType is not clSuccess, error argument contains additional error info
func (c *Connection) DistributedDomainConnect(hostName string, adminUser string, password string) (*ClusterError, error) {
	params := struct {
		HostName  string `json:"hostName"`
		AdminUser string `json:"adminUser"`
		Password  string `json:"password"`
	}{hostName, adminUser, password}
	data, err := c.CallRaw("DistributedDomain.connect", params)
	if err != nil {
		return nil, err
	}
	result := struct {
		Result struct {
			Result ClusterError `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &result)
	return &result.Result.Result, err
}

// DistributedDomainCopy - Copy domain from the master server.
// Parameters
//	domainName - name of the domain on the master server that you want to copy. Name can be obtained by using method getDomainsFromServer.
func (c *Connection) DistributedDomainCopy(domainName string) error {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	_, err := c.CallRaw("DistributedDomain.copy", params)
	return err
}

// DistributedDomainDisconnect - Disconnect server from the cluster.
func (c *Connection) DistributedDomainDisconnect() error {
	_, err := c.CallRaw("DistributedDomain.disconnect", nil)
	return err
}

// DistributedDomainGetDistributable - Retrieve domains, which can be distributed, from the master server as a standalone server.
// Parameters
//	connected - true means the caller is connected to cluster
//	authentication - Structure with a credential. Credential will be used when connected is false.
// Return
//	domainNames - List of domains which can be distributed (they have a directory service set).
func (c *Connection) DistributedDomainGetDistributable(authentication ClusterAuthentication, connected bool) (StringList, error) {
	params := struct {
		Authentication ClusterAuthentication `json:"authentication"`
		Connected      bool                  `json:"connected"`
	}{authentication, connected}
	data, err := c.CallRaw("DistributedDomain.getDistributable", params)
	if err != nil {
		return nil, err
	}
	domainNames := struct {
		Result struct {
			DomainNames StringList `json:"domainNames"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &domainNames)
	return domainNames.Result.DomainNames, err
}

// DistributedDomainGetRole - Return server role in the cluster.
func (c *Connection) DistributedDomainGetRole() (*ClusterRole, bool, error) {
	data, err := c.CallRaw("DistributedDomain.getRole", nil)
	if err != nil {
		return nil, false, err
	}
	role := struct {
		Result struct {
			Role          ClusterRole `json:"role"`
			IsMultiServer bool        `json:"isMultiServer"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &role)
	return &role.Result.Role, role.Result.IsMultiServer, err
}

// DistributedDomainGetServerList - Retrieve information about servers in the cluster.
// Return
//	servers - List of all servers in cluster.
func (c *Connection) DistributedDomainGetServerList() (ClusterServerList, error) {
	data, err := c.CallRaw("DistributedDomain.getServerList", nil)
	if err != nil {
		return nil, err
	}
	servers := struct {
		Result struct {
			Servers ClusterServerList `json:"servers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &servers)
	return servers.Result.Servers, err
}

// DistributedDomainGetHomeServerList - Retrieve information about servers in the cluster.
// Return
//	servers - List of all servers in cluster.
func (c *Connection) DistributedDomainGetHomeServerList() (HomeServerList, error) {
	data, err := c.CallRaw("DistributedDomain.getHomeServerList", nil)
	if err != nil {
		return nil, err
	}
	servers := struct {
		Result struct {
			Servers HomeServerList `json:"servers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &servers)
	return servers.Result.Servers, err
}

// DistributedDomainGetStatus - Notes: This method fails if caller has not admin rights; This method fails if there is no cluster
// Return
//	isInCluster true if server is not standalone
//	isError status of error in cluster
func (c *Connection) DistributedDomainGetStatus() (bool, bool, error) {
	data, err := c.CallRaw("DistributedDomain.getStatus", nil)
	if err != nil {
		return false, false, err
	}
	isInCluster := struct {
		Result struct {
			IsInCluster bool `json:"isInCluster"`
			IsError     bool `json:"isError"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &isInCluster)
	return isInCluster.Result.IsInCluster, isInCluster.Result.IsError, err
}
