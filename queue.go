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
// Parameters
//	query - search conditions
// Return
//	list - awaiting messages
//	volume - space occupied by messages in the queue
func (c *ServerConnection) QueueGet(query SearchQuery) (MessageInQueueList, *ByteValueWithUnits, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Queue.get", params)
	if err != nil {
		return nil, nil, err
	}
	list := struct {
		Result struct {
			List   MessageInQueueList `json:"list"`
			Volume ByteValueWithUnits `json:"volume"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, &list.Result.Volume, err
}

// QueueGetProcessed - List messages that are being processed by the server.
// Parameters
//	query - search conditions
// Return
//	list - processed messages
func (c *ServerConnection) QueueGetProcessed(query SearchQuery) (MessageInProcessList, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Queue.getProcessed", params)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List MessageInProcessList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// QueueRemove - Remove selected messages from the message queue.
// Parameters
//	messageIds - identifiers of messages to be deleted
func (c *ServerConnection) QueueRemove(messageIds KIdList) (int, error) {
	params := struct {
		MessageIds KIdList `json:"messageIds"`
	}{messageIds}
	data, err := c.CallRaw("Queue.remove", params)
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
func (c *ServerConnection) QueueRemoveAll() (int, error) {
	data, err := c.CallRaw("Queue.removeAll", nil)
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
// Parameters
//	senderPattern - sender pattern with wildcards
//	recipientPattern - recipient pattern with wildcards
func (c *ServerConnection) QueueRemoveMatching(senderPattern string, recipientPattern string) (int, error) {
	params := struct {
		SenderPattern    string `json:"senderPattern"`
		RecipientPattern string `json:"recipientPattern"`
	}{senderPattern, recipientPattern}
	data, err := c.CallRaw("Queue.removeMatching", params)
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
func (c *ServerConnection) QueueRun() error {
	_, err := c.CallRaw("Queue.run", nil)
	return err
}

// QueueTryToSend - Try to send selected messages.
// Parameters
//	messageIds - identifiers of messages to be sent immediately
func (c *ServerConnection) QueueTryToSend(messageIds KIdList) error {
	params := struct {
		MessageIds KIdList `json:"messageIds"`
	}{messageIds}
	_, err := c.CallRaw("Queue.tryToSend", params)
	return err
}
