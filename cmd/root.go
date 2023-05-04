package cmd

import (
	"github.com/spf13/cobra"
)

const (
	varUsernameLong  = "username"
	varUsernameShort = "u"
	varPasswordLong  = "password"
	varPasswordShort = "p"
	varHostLong      = "host"
	varHostShort     = "z"
)

var (
	username string
	password string
	host     string
)

func RootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "tlsg-cli",
		Version:      version,
		SilenceUsage: true,
		Short:        "TL SG Switches UnOfficial CLI",
		Long:         `Interacts with your TLSG switch from command line`,
	}

	addRootCmdFlags(rootCmd)

	// Subcommands
	rootCmd.AddCommand(InfoCmd(), PortSettingsCmd(), StatisticsCmd(), BackupCmd())
	return rootCmd
}

func addRootCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().
		StringVarP(&host, varHostLong, varHostShort, "localhost", "Switch hostname or IP")
	cmd.PersistentFlags().
		StringVarP(&username, varUsernameLong, varUsernameShort, "admin", "Switch login username")
	cmd.PersistentFlags().
		StringVarP(&password, varPasswordLong, varPasswordShort, "admin", "Switch login password")
}
