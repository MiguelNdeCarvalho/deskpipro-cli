package cmd

import (
	"fmt"
	"runtime"

	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-latest"
)

var (
	Author     = "MiguelNdeCarvalho"
	Repository = "deskpi-cli"
	Version    = "0.1.0"
)

func getVersion() string {
	versionDialog := fmt.Sprintf("DeskPi-CLI: v%s\nPlatform: %s/%s\nGoVersion: %s",
		Version, runtime.GOOS, runtime.GOARCH, runtime.Version(),
	)

	currentVersion, err := version.NewVersion(Version)

	if err != nil {
		if err := rootCmd.Execute(); err != nil {
			fmt.Print(err)
		}
	}

	latesVersion, err := version.NewVersion(getLatestVersion())

	if err != nil {
		if err := rootCmd.Execute(); err != nil {
			fmt.Print(err)
		}
	}

	if currentVersion.LessThan(latesVersion) {
		versionDialog += fmt.Sprintf("\n\n\033[33mYour version of Deskpi-CLI is outdated.\n"+
			"Please update to the latest: v%s by downloading from: \nhttps://github.com/%s"+
			"/%s/releases\033[0m", latesVersion.Original(), Author, Repository,
		)
	}

	return versionDialog
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

func getLatestVersion() string {
	githubTag := &latest.GithubTag{
		Owner:             Author,
		Repository:        Repository,
		FixVersionStrFunc: latest.DeleteFrontV(),
	}
	res, _ := latest.Check(githubTag, Version)
	return res.Current
}
