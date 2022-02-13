package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Author  = "MiguelNdeCarvalho"
	Version = "0.1.0"
)

func getVersion() string {
	return fmt.Sprintf(`DeskPi-CLI v%s
Platform: %s/%s
GoVersion: %s
Made by: %s`, Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), Author)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
	Short: "Show version info",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
