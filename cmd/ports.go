package cmd

import (
	"fmt"
	"os"

	"github.com/0ghny/tlsg10x"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func PortSettingsCmd() *cobra.Command {
	portSettingsCmd := &cobra.Command{
		Use:   "ports",
		Short: "Display Port Settings",
		Long:  `Display Port Settings`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := tlsg10x.New(host, username, password, nil)
			stats, err := client.PortsSettings()
			if err != nil {
				return fmt.Errorf("failed retrieving port settings: %v", err)
			}

			t := table.NewWriter()
			t.SetTitle("Ports Settings")
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Port", "Status", "Speed Config", "Speed Actual", "Flow Config", "Flow Actual"})
			for _, settings := range stats {
				t.AppendRow([]interface{}{
					settings.Name, settings.State.String(),
					settings.SpeedCfg.String(), settings.SpeedAct.String(),
					settings.FlowCtrlCfg.String(), settings.FlowCtrlAct.String(),
				})
				t.AppendSeparator()
			}
			t.Render()

			return nil
		},
	}
	return portSettingsCmd
}
