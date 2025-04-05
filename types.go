package flow

type Request struct {
	Method     string   `json:"method"`
	Parameters []string `json:"parameters"`
	Settings   map[string]string
	ID         int
	Error      string
}

type Result struct {
	Title     string
	SubTitle  string
	IcoPath   string
	RpcAction *JsonRpcAction `json:"JsonRPCAction,omitempty"`
}

type JsonRpcAction struct {
	Method     string   `json:"method"`
	Parameters []string `json:"parameters"`
}

type Response struct {
	Results      []Result `json:"result"`
	DebugMessage string   `json:",omitempty"`
}
