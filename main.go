package main

import (
	"embed"
	"fmt"

	"github.com/Tarocch1/file-admin/common"
)

var (
	Version   = ""
	GoVersion = ""
	BuildTime = ""
	CommitID  = ""
)

//go:embed static
var static embed.FS

func main() {
	common.ParseFlag()

	if common.FlagVersion {
		fmt.Println("Version:", Version)
		fmt.Println("Go Version:", GoVersion)
		fmt.Println("Build Time:", BuildTime)
		fmt.Println("Git Commit ID:", CommitID)
		return
	}

	common.GetRootDir()

	startServer()
}
