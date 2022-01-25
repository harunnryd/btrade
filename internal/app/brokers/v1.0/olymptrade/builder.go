package olymptrade

import (
	"github.com/chuckpreslar/emission"
	"github.com/sacOO7/gowebsocket"
)

// Builder ...
type Builder struct {
	option  Options
	builder WebsocketServiceClient
}

// SetBuilder ...
func (b *Builder) SetBuilder(builder WebsocketServiceClient) {
	b.builder = builder
}

// SetPair ...
func (b *Builder) SetPair(pair string) {
	b.option.pair = pair
}

// SetAccountMode ...
func (b *Builder) SetAccountMode(accountMode string) {
	b.option.accountMode = accountMode
}

// SetAccountID ...
func (b *Builder) SetAccountID(accountID string) {
	b.option.accountID = accountID
}

// SetEmitter ...
func (b *Builder) SetEmitter(emitter *emission.Emitter) {
	b.option.emitter = emitter
}

// SetWebsocket ...
func (b *Builder) SetWebsocket(websocket gowebsocket.Socket) {
	b.option.websocket = websocket
}

// Build ...
func (b *Builder) Build() WebsocketServiceClient {
	b.builder.SetOptions(b.option)
	return b.builder
}
