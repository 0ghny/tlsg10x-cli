package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/0ghny/tlsg10x"
	"github.com/spf13/cobra"
)

const (
	varBackupOutputLong  = "output"
	varBackupOutputShort = "o"
)

var (
	backupOutput string
)

func BackupCmd() *cobra.Command {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Creates a configuration backup",
		Long:  `Creates a configuration backup to a selected destination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := tlsg10x.New(host, username, password, nil)

			// Retrieve backup from switch
			backupContent, err := client.CreateBackup()
			if err != nil {
				return fmt.Errorf("failed retrieving backup from device: %v", err)
			}

			// Check output parameter
			if backupOutput == "" {
				file, err := ioutil.TempFile("", "tlsg-backup")
				if err != nil {
					return fmt.Errorf("failed creating temp file to store backup: %v", err)
				}
				backupOutput = file.Name()
			}

			// Save backup to disk
			err = os.WriteFile(backupOutput, backupContent, 0644)
			if err != nil {
				return fmt.Errorf("failed writing backup file to disk: %v", err)
			}

			fmt.Printf("Backup saved to disk at %s\n", backupOutput)

			return nil
		},
	}
	addBackupCmdFlags(backupCmd)
	return backupCmd
}

func addBackupCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().
		StringVarP(&backupOutput, varBackupOutputLong, varBackupOutputShort, "", "Backup file destination")
}
