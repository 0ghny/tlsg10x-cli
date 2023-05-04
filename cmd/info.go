package cmd

import (
	"fmt"
	"os"

	"github.com/0ghny/tlsg10x"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func InfoCmd() *cobra.Command {
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Prints Switch System information",
		Long:  "Prints Switch System ",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := tlsg10x.New(host, username, password, nil)
			sysInfo, err := client.SystemInfo()
			if err != nil {
				return fmt.Errorf("failed retrieving system info: %v", err)
			}

			t := table.NewWriter()
			t.SetTitle("System Information")
			t.SetOutputMirror(os.Stdout)
			t.AppendRow([]interface{}{"Device Description", sysInfo.Name})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"MAC Address", sysInfo.MacAddress})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"IP Address", sysInfo.IPAddress})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"SUbnet Mask", sysInfo.NetMask})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"Default Gateway", sysInfo.Gateway})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"Firmware Version", sysInfo.Firmware})
			t.AppendSeparator()
			t.AppendRow([]interface{}{"Hardware Version", sysInfo.Hardware})
			t.Render()

			return nil
		},
	}
	return infoCmd
}
