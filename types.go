package flow

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// `jsonrpc` is not used by Flow Launcher
type JsonRpcBase struct {
	ID  int   `json:"id"`
	Err Error `json:"error"`
}

type Request struct {
	JsonRpcBase
	Method     string   `json:"method"`
	Parameters []string `json:"parameters"`
	Settings   map[string]any
}

type Response struct {
	JsonRpcBase
	Results        []*Result `json:"result"`
	SettingsChange map[string]any
	DebugMessage   string `json:",omitempty"`
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

func NewResponse(req *Request) *Response {
	res := &Response{
		JsonRpcBase: JsonRpcBase{
			ID: req.ID,
		},
		Results: make([]*Result, 0),
	}

	return res
}

func (resp *Response) AddResult(r *Result) {
	resp.Results = append(resp.Results, r)
}
