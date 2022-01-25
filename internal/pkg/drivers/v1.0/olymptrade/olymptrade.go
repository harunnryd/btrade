package olymptrade

import (
	"github.com/sacOO7/gowebsocket"
)

// NewOlymptrade ...
func NewOlymptrade(options Options) (websocket gowebsocket.Socket, err error) {
	websocket = gowebsocket.New(options.Host)
	websocket.RequestHeader.Set("User-Agent", options.HeaderUserAgent)
	websocket.RequestHeader.Set("Accept", options.HeaderAccept)
	websocket.RequestHeader.Set("Accept-Language", options.HeaderAcceptLanguage)
	websocket.RequestHeader.Set("Accept-Encoding", options.HeaderAcceptEncoding)
	websocket.RequestHeader.Set("Origin", options.HeaderOrigin)
	websocket.RequestHeader.Set("Sec-Fetch-Dest", options.HeaderSecFetchDest)
	websocket.RequestHeader.Set("Sec-Fetch-Mode", options.HeaderSecFetchMode)
	websocket.RequestHeader.Set("Sec-Fetch-Site", options.HeaderSecFetchSite)
	websocket.RequestHeader.Set("Pragma", options.HeaderPragma)
	websocket.RequestHeader.Set("Cache-Control", options.HeaderCacheControl)
	websocket.RequestHeader.Set("Cookie", options.HeaderCookie)

	return
}
