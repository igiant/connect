package connect

import "encoding/json"

type MessageStatus string

const (
	msExecuting        MessageStatus = "msExecuting"
	msBackup           MessageStatus = "msBackup"
	msContentFiltering MessageStatus = "msContentFiltering"
	msAntivirusControl MessageStatus = "msAntivirusControl" // shouldn't it be Antivirus Check
	msLocalDelivering  MessageStatus = "msLocalDelivering"
	msSmtpDelivering   MessageStatus = "msSmtpDelivering"
	msFinishing        MessageStatus = "msFinishing" // Terminating in AC manual vs. Finishing in AC
)

// MessageInQueue - Message waiting in the queue
type MessageInQueue struct {
	Id           KId                `json:"id"`           // Queue ID
	CreationTime string             `json:"creationTime"` // when the message was created
	NextTry      string             `json:"nextTry"`      // when to try send message
	MessageSize  ByteValueWithUnits `json:"messageSize"`  // message size in appropriate units
	From         string             `json:"from"`         // sender email address
	To           string             `json:"to"`           // recipient email address
	Status       string             `json:"status"`       // message status
	AuthSender   string             `json:"authSender"`   // email address of authenticated sender
	SenderIp     IpAddress          `json:"senderIp"`     // IP address of authenticated sender
}

// MessageInProcess - Message being processed by server
type MessageInProcess struct {
	Id          KId                `json:"id"`          // Queue ID
	MessageSize ByteValueWithUnits `json:"messageSize"` // message size in appropriate units
	From        string             `json:"from"`        // sender email address
	To          string             `json:"to"`          // recipient email address
	Status      MessageStatus      `json:"status"`      // message status
	Percentage  int                `json:"percentage"`  // only for processing: completed percentage
	Server      string             `json:"server"`      // server name or IP
	Time        string             `json:"time"`        // time in process
}

type MessageInQueueList []MessageInQueue

type MessageInProcessList []MessageInProcess

// QueueGet - Obtain a list of queued messages.
//	query - search conditions
// Return
//	list - awaiting messages
//  totalItems - number of listed messages
//	volume - space occupied by messages in the queue
func (s *ServerConnection) QueueGet(query SearchQuery) (MessageInQueueList, int, *ByteValueWithUnits, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Queue.get", params)
	if err != nil {
		return nil, 0, nil, err
	}
	list := struct {
		Result struct {
			List       MessageInQueueList `json:"list"`
			Volume     ByteValueWithUnits `json:"volume"`
			TotalItems int                `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, &list.Result.Volume, err
}

// QueueGetProcessed - List messages that are being processed by the server.
//	query - search conditions
// Return
//	list - processed messages
//  totalItems - number of processed messages
func (s *ServerConnection) QueueGetProcessed(query SearchQuery) (MessageInProcessList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Queue.getProcessed", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       MessageInProcessList `json:"list"`
			TotalItems int                  `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// QueueRemove - Remove selected messages from the message queue.
//	messageIds - identifiers of messages to be deleted
func (s *ServerConnection) QueueRemove(messageIds KIdList) (int, error) {
	params := struct {
		MessageIds KIdList `json:"messageIds"`
	}{messageIds}
	data, err := s.CallRaw("Queue.remove", params)
	if err != nil {
		return 0, err
	}
	deleteItems := struct {
		Result struct {
			DeleteItems int `json:"deleteItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deleteItems)
	return deleteItems.Result.DeleteItems, err
}

// QueueRemoveAll - Remove all message from the queue.
func (s *ServerConnection) QueueRemoveAll() (int, error) {
	data, err := s.CallRaw("Queue.removeAll", nil)
	if err != nil {
		return 0, err
	}
	deleteItems := struct {
		Result struct {
			DeleteItems int `json:"deleteItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deleteItems)
	return deleteItems.Result.DeleteItems, err
}

// QueueRemoveMatching - Remove all messages matching a pattern from the message queue.
//	senderPattern - sender pattern with wildcards
//	recipientPattern - recipient pattern with wildcards
func (s *ServerConnection) QueueRemoveMatching(senderPattern string, recipientPattern string) (int, error) {
	params := struct {
		SenderPattern    string `json:"senderPattern"`
		RecipientPattern string `json:"recipientPattern"`
	}{senderPattern, recipientPattern}
	data, err := s.CallRaw("Queue.removeMatching", params)
	if err != nil {
		return 0, err
	}
	deleteItems := struct {
		Result struct {
			DeleteItems int `json:"deleteItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deleteItems)
	return deleteItems.Result.DeleteItems, err
}

// QueueRun - Try to process message queue immediately.
func (s *ServerConnection) QueueRun() error {
	_, err := s.CallRaw("Queue.run", nil)
	return err
}

// QueueTryToSend - Try to send selected messages.
//	messageIds - identifiers of messages to be sent immediately
func (s *ServerConnection) QueueTryToSend(messageIds KIdList) error {
	params := struct {
		MessageIds KIdList `json:"messageIds"`
	}{messageIds}
	_, err := s.CallRaw("Queue.tryToSend", params)
	return err
}
