package connect

type parameters struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int         `json:"id"`
	Token   *string     `json:"token,omitempty"`
	Params  interface{} `json:"params,omitempty"`
}

type Config struct {
	url string
	id  int
}

type Application struct {
	Name    string
	Vendor  string
	Version string
}

type loginStruct struct {
	UserName    string      `json:"userName"`
	Password    string      `json:"password"`
	Application Application `json:"application"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LocalizableMessageParameters struct {
	PositionalParameters []string `json:"positionalParameters"` //additional strings to replace the placeholders in message (first string replaces %1 etc.)
	Plurality            int      `json:"plurality"`            //count of items, used to distinguish among singular/paucal/plural; 1 for messages with no counted items
}

type Error struct {
	InputIndex        int                          `json:"inputIndex"`        //0-based index to input array, e.g. 3 means that the relates to the 4th element of the input parameter array
	Code              int                          `json:"code"`              //-32767..-1 (JSON-RPC) or 1..32767 (application)
	Message           string                       `json:"message"`           //text with placeholders %1, %2, etc., e.g. "User %1 cannot be deleted."
	MessageParameters LocalizableMessageParameters `json:"messageParameters"` //strings to replace placeholders in message, and message plurality.
}

type ErrorList []Error

// Details about a particular item created.
type CreateResult struct {
	InputIndex int    `json:"inputIndex"` //0-based index to input array, e.g. 3 means that the relates to the 4th element of the input parameter array
	ID         string `json:"id"`         //ID of created item
}

type CreateResultList []CreateResult
