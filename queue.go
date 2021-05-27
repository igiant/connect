package connect

import "encoding/json"

// TODO QueueGet
// QueueGet obtains a list of queued messages.
// Parameters
//      list	- awaiting messages
//      totalItems	- number of listed messages
//      volume	- space occupied by messages in the queue
//      query	- search conditions
func (c *Connection) QueueGet(query SearchQuery) {
	//params := struct {
	//	SearchQuery `json:"query"`
	//}{query}
	//data, err := c.CallRaw("Queue.get", params)
	//if err != nil {
}

// QueueRun - try to process message queue immediately.
func (c *Connection) QueueRun() error {
	_, err := c.CallRaw("Queue.run", nil)
	return err
}

// QueueGetProcessed - list messages that are being processed by the server.
//  Parameters
//      query	- search conditions
func (c *Connection) QueueGetProcessed(query SearchQuery) ([]string, error) {
	params := struct {
		SearchQuery `json:"query"`
	}{query}
	data, err := c.CallRaw("Queue.getProcessed", params)
	if err != nil {
		return nil, err
	}
	messageInProcessList := struct {
		Result struct {
			MessageInProcessList []string `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &messageInProcessList)
	return messageInProcessList.Result.MessageInProcessList, err
}

// QueueRemove - remove selected messages from the message queue.
//  Parameters
//      messageIDs	- identifiers of messages to be deleted
// Return number of items deleted from the message queue
func (c *Connection) QueueRemove(messageIDs []string) (int, error) {
	params := struct {
		MessageIDs []string `json:"messageIds"`
	}{messageIDs}
	data, err := c.CallRaw("Queue.remove", params)
	if err != nil {
		return 0, err
	}
	deletedItems := struct {
		Result struct {
			DeletedItems int `json:"deletedItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deletedItems)
	return deletedItems.Result.DeletedItems, err
}

// QueueRemoveAll - remove all message from the queue.
// Return number of items deleted from the message queue
func (c *Connection) QueueRemoveAll() (int, error) {
	data, err := c.CallRaw("Queue.removeAll", nil)
	if err != nil {
		return 0, err
	}
	deletedItems := struct {
		Result struct {
			DeletedItems int `json:"deletedItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deletedItems)
	return deletedItems.Result.DeletedItems, err
}

// QueueRemoveMatching - remove all messages matching a pattern from the message queue.
//  Parameters
//      senderPattern	    - sender pattern with wildcards
//      recipientPattern    - recipient pattern with wildcards
// Return number of items deleted from the message queue
func (c *Connection) QueueRemoveMatching(senderPattern, recipientPattern string) (int, error) {
	params := struct {
		SenderPattern    string `json:"senderPattern"`
		RecipientPattern string `json:"recipientPattern"`
	}{senderPattern, recipientPattern}
	data, err := c.CallRaw("Queue.removeMatching", params)
	if err != nil {
		return 0, err
	}
	deletedItems := struct {
		Result struct {
			DeletedItems int `json:"deletedItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &deletedItems)
	return deletedItems.Result.DeletedItems, err
}

// QueueTryToSend - try to send selected messages.
//  Parameters
//      messageIDs	- identifiers of messages to be sent immediately.
func (c *Connection) QueueTryToSend(messageIDs []string) error {
	params := struct {
		MessageIDs []string `json:"messageIds"`
	}{messageIDs}
	_, err := c.CallRaw("Queue.tryToSend", params)
	return err
}
