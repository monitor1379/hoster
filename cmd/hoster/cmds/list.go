package cmds

/*
 * @Date: 2020-11-29 22:23:25
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 22:42:17
 */

import (
	"os"

	"github.com/monitor1379/hoster"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List address-host mappings",
		RunE:  listCmdRunE,
	}
	return listCmd
}

func listCmdRunE(c *cobra.Command, args []string) error {
	hostFilePath := c.Parent().PersistentFlags().Lookup("file").Value.String()
	hm, err := hoster.New(hostFilePath)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address", "Host"})

	mappings := hm.Mappings()
	for _, mapping := range mappings {
		if mapping.IsEmptyLine() || mapping.IsOnlyComment() {
			continue
		}
		for _, host := range mapping.Hosts {
			table.Append([]string{mapping.Address, host})
		}
	}

	table.Render()
	return nil
}
