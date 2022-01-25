package worker

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
	// CmdUse ...
	CmdUse = "worker"

	// CmdAliases ...
	CmdAliases = []string{"w"}

	// App ...
	App *appcontext.AppContext

	// ConfigPaths ...
	ConfigPaths = [...]string{
		"./*.toml",
		"./params/*.toml",
		"/opt/" + AppName + "/params/*.toml",
	}
)
