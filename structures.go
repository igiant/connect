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
