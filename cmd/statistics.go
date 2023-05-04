package cmd

import (
	"fmt"
	"os"

	"github.com/0ghny/tlsg10x"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func StatisticsCmd() *cobra.Command {
	statisticsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Retrieve Port Statistics",
		Long:  `Retrieve Port Statistics`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := tlsg10x.New(host, username, password, nil)
			stats, err := client.PortsStats()
			if err != nil {
				return fmt.Errorf("failed retrieving port stats: %v", err)
			}

			tablePrinter := table.NewWriter()
			tablePrinter.SetTitle("Ports Statistics")
			tablePrinter.SetOutputMirror(os.Stdout)
			tablePrinter.AppendHeader(table.Row{"Port", "Status", "Link Status", "TxGoodPkt", "TxBadPkt", "RxGoodPkt", "RxBadPkt"})
			for _, portStats := range stats {
				tablePrinter.AppendRow([]interface{}{
					portStats.Name, portStats.State.String(), portStats.LinkStatus.String(),
					portStats.TxGoodPkt, portStats.TxBadPkt,
					portStats.RxGoodPkt, portStats.RxBadPkt,
				})
				tablePrinter.AppendSeparator()
			}
			tablePrinter.Render()
			return nil
		},
	}
	return statisticsCmd
}
