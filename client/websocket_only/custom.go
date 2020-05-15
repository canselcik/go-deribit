package websocket_only

import "github.com/go-openapi/runtime"

func (a *Client) GetTransport() runtime.ClientTransport {
	return a.transport
}