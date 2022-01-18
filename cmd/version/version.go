package version

import (
	"fmt"
	"github.com/harunnryd/btrade/internal/app/entity"
	"os"

	"github.com/harunnryd/btrade/cmd/root"
	"github.com/harunnryd/btrade/internal/pkg/utils/atexit"
)

// Start ...
func Start() {
	code := 0

	defer func() {
		os.Exit(code)
	}()

	defer func() {
		atexit.AtExit()
	}()

	atexit.Add(Shutdown)

	fmt.Println("App\t\t:", root.AppName)
	fmt.Println("Desc\t\t:", root.AppDescLong)
	fmt.Println("Build Date\t:", entity.BuildDate)
	fmt.Println("Git Commit\t:", entity.GitCommit)
	fmt.Println("Version\t\t:", entity.Version)
	fmt.Println("Environment\t:", entity.Environment)
	fmt.Println("Go Version\t:", entity.GoVersion)
	fmt.Println("OS / Arch\t:", entity.OsArch)
}

// Shutdown ...
func Shutdown() {
	fmt.Println("\n-#-")
}
