package root

import "github.com/harunnryd/btrade/internal/app/appcontext"

const (
	// AppName ...
	AppName string = "btrade"

	// AppDescShort ...
	AppDescShort string = "bot auto trade"

	// AppDescLong ...
	AppDescLong string = `bot auto trade in binary option`
)

var (
	// App ...
	App *appcontext.AppContext

	// ConfigPaths ...
	ConfigPaths = [...]string{
		"./*.toml",
		"./params/*.toml",
		"/opt/" + AppName + "/params/*.toml",
	}
)
