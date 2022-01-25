package olymptrade

// Options ...
type Options struct {
	Host                 string `deepcopier:"field:Host"`
	HeaderUserAgent      string `deepcopier:"field:HeaderUserAgent"`
	HeaderAccept         string `deepcopier:"field:HeaderAccept"`
	HeaderCookie         string `deepcopier:"field:HeaderCookie"`
	HeaderAcceptLanguage string `deepcopier:"field:HeaderAcceptLanguage"`
	HeaderAcceptEncoding string `deepcopier:"field:HeaderAcceptEncoding"`
	HeaderOrigin         string `deepcopier:"field:HeaderOrigin"`
	HeaderSecFetchDest   string `deepcopier:"field:HeaderSecFetchDest"`
	HeaderSecFetchMode   string `deepcopier:"field:HeaderSecFetchMode"`
	HeaderSecFetchSite   string `deepcopier:"field:HeaderSecFetchSite"`
	HeaderPragma         string `deepcopier:"field:HeaderPragma"`
	HeaderCacheControl   string `deepcopier:"field:HeaderCacheControl"`
}
