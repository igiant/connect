package connect

import "encoding/json"

type TypeExpStatistics string

const (
	expStatShort TypeExpStatistics = "expStatShort"
	expStatFull  TypeExpStatistics = "expStatFull"
)

type OccupiedStorage struct {
	Total      ByteValueWithUnits `json:"total"`      // total space on disc
	Occupied   ByteValueWithUnits `json:"occupied"`   // occupied space on disc
	Percentage string             `json:"percentage"` // how many per cent is occupied
}

type MessageThroughput struct {
	Count      string             `json:"count"`      // how many messages
	Volume     ByteValueWithUnits `json:"volume"`     // how much space is occupied by messages
	Recipients string             `json:"recipients"` // how many recipients in messages
}

type FailureAndBounce struct {
	TransientFailures string `json:"transientFailures"` // transient delivery failures
	PermanentFailures string `json:"permanentFailures"` // permanent delivery failures
}

type Notifications struct {
	Success string `json:"success"` // how many sent success notifications
	Delay   string `json:"delay"`   // how many sent delay notifications
	Failure string `json:"failure"` // how many sent failure notifications
}

type AntivirusStats struct {
	CheckedAttachments string `json:"checkedAttachments"` // how many checked attachments
	FoundViruses       string `json:"foundViruses"`       // how many found viruses
	ProhibitedTypes    string `json:"prohibitedTypes"`    // how many found prohibited filenames/MIME types
}

type SpamStats struct {
	Checked         string `json:"checked"`         // how many checked messages
	Tagged          string `json:"tagged"`          // how many tagged messages
	Rejected        string `json:"rejected"`        // how many rejected messages
	MarkedAsSpam    string `json:"markedAsSpam"`    // how many messages were marked as spam by users
	MarkedAsNotSpam string `json:"markedAsNotSpam"` // how many messages were marked as NOT spam by users
}

type OtherStats struct {
	Largest ByteValueWithUnits `json:"largest"` // the largest messages received by server
	Loops   string             `json:"loops"`   // how many detected message loops
}

type SmtpServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
	LostConnections          string `json:"lostConnections"`
	RejectedByBlacklist      string `json:"rejectedByBlacklist"`
	AuthenticationAttempts   string `json:"authenticationAttempts"`
	AuthenticationFailures   string `json:"authenticationFailures"`
	RejectedRelays           string `json:"rejectedRelays"`
	AcceptedMessages         string `json:"acceptedMessages"`
}

type SmtpClientStats struct {
	ConnectionAttempts string `json:"connectionAttempts"`
	DnsFailures        string `json:"dnsFailures"`
	ConnectionFailures string `json:"connectionFailures"`
	ConnectionLosses   string `json:"connectionLosses"`
}

type Pop3ServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
	AuthenticationFailures   string `json:"authenticationFailures"`
	SentMessages             string `json:"sentMessages"`
}

type Pop3ClientStats struct {
	ConnectionAttempts     string `json:"connectionAttempts"`
	ConnectionFailures     string `json:"connectionFailures"`
	AuthenticationFailures string `json:"authenticationFailures"`
	TotalDownloads         string `json:"totalDownloads"`
}

type ImapServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
	AuthenticationFailures   string `json:"authenticationFailures"`
}

type LdapServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
	AuthenticationFailures   string `json:"authenticationFailures"`
	TotalSearchRequests      string `json:"totalSearchRequests"`
}

type WebServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
}

type XmppServerStats struct {
	TotalIncomingConnections string `json:"totalIncomingConnections"`
	AuthenticationFailures   string `json:"authenticationFailures"`
}

type DnsResolverStats struct {
	HostnameQueries       string `json:"hostnameQueries"`
	CachedHostnameQueries string `json:"cachedHostnameQueries"`
	MxQueries             string `json:"mxQueries"`
	CachedMxQueries       string `json:"cachedMxQueries"`
}

type AntibombingStats struct {
	RejectedConnections    string `json:"rejectedConnections"`
	RejectedMessages       string `json:"rejectedMessages"`
	RejectedHarvestAttacks string `json:"rejectedHarvestAttacks"`
}

type GreylistingStats struct {
	MessagesAccepted string `json:"messagesAccepted"`
	MessagesDelayed  string `json:"messagesDelayed"`
	MessagesSkipped  string `json:"messagesSkipped"`
}

type ServerStatistics struct {
	Start             DateTimeStamp     `json:"start"`
	Uptime            Distance          `json:"uptime"` // server uptime
	Storage           OccupiedStorage   `json:"storage"`
	Received          MessageThroughput `json:"received"`          // messages received by server
	StoredInQueue     MessageThroughput `json:"storedInQueue"`     // messages stored in queue
	Transmitted       MessageThroughput `json:"transmitted"`       // messages transmitted by server
	DeliveredToLocals MessageThroughput `json:"deliveredToLocals"` // messages delivered to local domains
	Mx                MessageThroughput `json:"mx"`                // messages sent to remote MX servers
	Relay             MessageThroughput `json:"relay"`             // messages sent to relay server
	Failures          FailureAndBounce  `json:"failures"`
	DeliveryStatus    Notifications     `json:"deliveryStatus"`
	Antivirus         AntivirusStats    `json:"antivirus"`
	Spam              SpamStats         `json:"spam"`
	Other             OtherStats        `json:"other"`
	SmtpServer        SmtpServerStats   `json:"smtpServer"`
	SmtpClient        SmtpClientStats   `json:"smtpClient"`
	Pop3Server        Pop3ServerStats   `json:"pop3Server"`
	Pop3Client        Pop3ClientStats   `json:"pop3Client"`
	ImapServer        ImapServerStats   `json:"imapServer"`
	LdapServer        LdapServerStats   `json:"ldapServer"`
	WebServer         WebServerStats    `json:"webServer"`
	XmppServer        XmppServerStats   `json:"xmppServer"`
	DnsResolver       DnsResolverStats  `json:"dnsResolver"`
	Antibombing       AntibombingStats  `json:"antibombing"`
	Greylisting       GreylistingStats  `json:"greylisting"`
}

type Scale struct {
	Id         int `json:"id"`
	ScaleTime  int `json:"scaleTime"`  // The time scale
	SampleTime int `json:"sampleTime"` // The sample scale
}

type ScaleList []Scale

// Chart - Descriptions of charts graph
type Chart struct {
	Classname  string    `json:"classname"`  // A class name of chart
	Name       string    `json:"name"`       // A chart name
	Xtype      string    `json:"xtype"`      // An x scale type
	Ytype      string    `json:"ytype"`      // An y scale type
	ScaleCount int       `json:"scaleCount"` // A count of scales
	Scale      ScaleList `json:"scale"`      // List of scales
}

type ChartList []Chart

type ChartValueList []int

type ChartRowNamesList []string

type ChartRowValuesList []ChartValueList

// ChartData - Values of charts graph
type ChartData struct {
	XName       string             `json:"xName"`       // Name of X axis
	XValues     ChartValueList     `json:"xValues"`     // Values of X axis
	CountValues int                `json:"countValues"` // A count of values in X axis
	CountRows   int                `json:"countRows"`   // A count of rows
	RowNames    ChartRowNamesList  `json:"rowNames"`    // Array of names of rows
	RowValues   ChartRowValuesList `json:"rowValues"`   // Array of values of rows
}

// StatisticsExportToHtml - Export server statistics to HTML format.
//	type - export type
//	lang - export language
// Return
//	fileDownload - description of output file
func (s *ServerConnection) StatisticsExportToHtml(typeExpStatistics TypeExpStatistics, lang string) (*Download, error) {
	params := struct {
		Type TypeExpStatistics `json:"type"`
		Lang string            `json:"lang"`
	}{typeExpStatistics, lang}
	data, err := s.CallRaw("Statistics.exportToHtml", params)
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

// StatisticsGet - Obtain overall server statistics.
func (s *ServerConnection) StatisticsGet() (*ServerStatistics, error) {
	data, err := s.CallRaw("Statistics.get", nil)
	if err != nil {
		return nil, err
	}
	statistics := struct {
		Result struct {
			Statistics ServerStatistics `json:"statistics"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &statistics)
	return &statistics.Result.Statistics, err
}

// StatisticsGetCharts - Obtain descriptions of charts graphs.
// Return
//	chartList - list of descriptions of charts graphs
func (s *ServerConnection) StatisticsGetCharts() (ChartList, error) {
	data, err := s.CallRaw("Statistics.getCharts", nil)
	if err != nil {
		return nil, err
	}
	chartList := struct {
		Result struct {
			ChartList ChartList `json:"chartList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &chartList)
	return chartList.Result.ChartList, err
}

// StatisticsGetChartData - Obtain values of charts graph.
//	classname - classname of charts graph
//	name - name of charts graph
//	scaleId - ID of scale
// Return
//	chartData - values of charts graph
func (s *ServerConnection) StatisticsGetChartData(classname string, name string, scaleId int) (*ChartData, error) {
	params := struct {
		Classname string `json:"classname"`
		Name      string `json:"name"`
		ScaleId   int    `json:"scaleId"`
	}{classname, name, scaleId}
	data, err := s.CallRaw("Statistics.getChartData", params)
	if err != nil {
		return nil, err
	}
	chartData := struct {
		Result struct {
			ChartData ChartData `json:"chartData"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &chartData)
	return &chartData.Result.ChartData, err
}

// StatisticsReset - Reset statistics data.
func (s *ServerConnection) StatisticsReset() error {
	_, err := s.CallRaw("Statistics.reset", nil)
	return err
}
