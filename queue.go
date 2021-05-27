package connect

import "encoding/json"

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

// QueueRemove - Remove selected messages from the message queue.
//  Parameters
//      messageIDs	- identifiers of messages to be deleted
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
