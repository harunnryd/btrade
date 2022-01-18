package root

import (
	"fmt"

	"github.com/harunnryd/btrade/internal/app/config"
	uConfig "github.com/harunnryd/btrade/internal/pkg/utils/config"
	"github.com/harunnryd/btrade/internal/pkg/utils/loghelper"
	"github.com/joho/godotenv"
)

// InitConfig ...
func InitConfig() {
	_ = godotenv.Load("params/.env")
	var cfg config.Config

	if err := uConfig.LoadConfigs(ConfigPaths[:], &cfg); err != nil {
		loghelper.Panic(App.Ctx, "failed to init config", err)
	}

	_ = App.InitConfig(cfg)
}

// Start ...
func Start() {
}

// Shutdown ...
func Shutdown() {
	fmt.Println("\n-#-")
}
