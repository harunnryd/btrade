package olymptrade

import (
	driver "github.com/harunnryd/btrade/internal/pkg/drivers/v1.0/olymptrade"
	"github.com/ulule/deepcopier"
)

// Olymptrade ...
type Olymptrade struct {
	IsEnabled            bool   `mapstructure:"is_enabled"`
	Host                 string `mapstructure:"host"`
	HeaderUserAgent      string `mapstructure:"header_user_agent"`
	HeaderAccept         string `mapstructure:"header_accept"`
	HeaderCookie         string `mapstructure:"header_cookie"`
	HeaderAcceptLanguage string `mapstructure:"header_accept_language"`
	HeaderAcceptEncoding string `mapstructure:"header_accept_encoding"`
	HeaderOrigin         string `mapstructure:"header_origin"`
	HeaderSecFetchDest   string `mapstructure:"header_sec_fetch_dest"`
	HeaderSecFetchMode   string `mapstructure:"header_sec_fetch_mode"`
	HeaderSecFetchSite   string `mapstructure:"header_sec_fetch_site"`
	HeaderPragma         string `mapstructure:"header_pragma"`
	HeaderCacheControl   string `mapstructure:"header_cache_control"`
	AccountMode          string `mapstructure:"account_mode"`
	AccountID            string `mapstructure:"account_id"`
	Pair                 string `mapstructure:"pair"`
}

// Options ...
func (cfg *Olymptrade) Options() (rd driver.Options) {
	_ = deepcopier.Copy(cfg).To(&rd)
	return
}
