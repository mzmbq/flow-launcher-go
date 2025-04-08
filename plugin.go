package flow

import (
	"encoding/json"
	"fmt"
	"os"
)

type QueryHandler func(*Request) *Response
type CtxMenuHandler func(*Request) *Response
type ActionHandler func([]string) *Response

type Plugin struct {
	queryHandler   QueryHandler
	ctxMenuHandler CtxMenuHandler
	actionHandlers map[string]ActionHandler
}

func NewPlugin() *Plugin {
	return &Plugin{
		queryHandler:   nil,
		actionHandlers: make(map[string]ActionHandler),
	}
}

// Handles JSON-RPC and writes the result to stdout
func (p *Plugin) HandleRPC(query string) error {
	req := &Request{}
	err := json.Unmarshal([]byte(query), req)
	if err != nil {
		return err
	}

	var res *Response
	switch {
	case req.Method == "query":
		if p.queryHandler == nil {
			break
		}
		res = p.queryHandler(req)

	case req.Method == "context_menu":
		if p.ctxMenuHandler == nil {
			break
		}
		res = p.ctxMenuHandler(req)

	case p.isAction(req.Method):
		ah := p.actionHandlers[req.Method]
		res = ah(req.Parameters)

	default:
		res = &Response{
			Results:      make([]*Result, 0),
			DebugMessage: "error: invalid method",
		}
	}

	err = json.NewEncoder(os.Stdout).Encode(res)
	if err != nil {
		return nil
	}
	return nil
}

// Sets the query handler
func (p *Plugin) Query(h QueryHandler) {
	p.queryHandler = h
}

// Sets the actions handler for `action`
func (p *Plugin) Action(action string, h ActionHandler) {
	if _, ok := p.actionHandlers[action]; ok {
		panic(fmt.Sprintf("action name must be unique: '%s'", action))
	}
	if action == "query" || action == "context_menu" {
		panic("invalid action names: 'query', 'context_menu'")
	}

	p.actionHandlers[action] = h
}

// Sets the context menu handler
func (p *Plugin) ContextMenu(h CtxMenuHandler) {
	p.ctxMenuHandler = h
}

func ErrorResponse(text string) *Response {
	return &Response{
		Results:      nil,
		DebugMessage: text,
	}
}

func (p *Plugin) isAction(method string) (ok bool) {
	_, ok = p.actionHandlers[method]
	return
}
