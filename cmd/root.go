package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deskpipro-cli",
	Short: "DeskPiPro-CLI is an Unofficial tool to control DeskPiPro devices",
	Long:  "DeskPiPro-CLI is an Unofficial tool to control DeskPiPro devices",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	// Main Commands
	rootCmd.AddCommand(fanCLI)
	rootCmd.AddCommand(safeShutdownCLI)

	// Fan Commands and Flags
	fanCLI.AddCommand(fanCLIDaemon)
	fanCLI.AddCommand(fanCLISet)
	fanCLI.AddCommand(fanCLIStop)
}

var fanCLI = &cobra.Command{
	Use:   "fan",
	Short: "Control CPU fan",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var safeShutdownCLI = &cobra.Command{
	Use:   "fan",
	Short: "Control CPU fan",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var fanCLIDaemon = &cobra.Command{
	Use:   "daemon",
	Short: "Daemon to control fan speed automatically",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Fan Control Daemon")
		return nil
	},
}

var fanCLISet = &cobra.Command{
	Use:   "set",
	Short: "Set fan speed in Percentage(%)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing arg for fan speed")
		} else if len(args) > 1 {
			return fmt.Errorf("accepts 1 arg, received %d arg(s)", len(args))
		}
		arg, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if arg < 0 || arg > 100 {
			return fmt.Errorf("arg should need to be between 0 and 100 (%%)")
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Fan Control Daemon")

		fanSet(args[0])
		return nil
	},
}

var fanCLIStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop fan speed",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Stop fan")
		return nil
	},
}
