package cmds

/*
 * @Date: 2020-11-29 22:51:12
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 23:00:29
 */

import (
	"errors"
	"os"

	"github.com/monitor1379/hoster"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newLookupCommand() *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:   "lookup",
		Short: "Lookup by address or host",
		RunE:  lookupCmdRunE,
	}
	lookupCmd.PersistentFlags().String("host", "", "host name")
	lookupCmd.PersistentFlags().String("address", "", "address of ip")

	return lookupCmd
}

func lookupCmdRunE(c *cobra.Command, args []string) error {
	hostFilePath := c.Parent().PersistentFlags().Lookup("file").Value.String()
	hm, err := hoster.New(hostFilePath)
	if err != nil {
		return err
	}

	address := c.PersistentFlags().Lookup("address").Value.String()
	host := c.PersistentFlags().Lookup("host").Value.String()

	var (
		mapping *hoster.Mapping
		ok      bool
	)
	if address != "" {
		mapping, ok = hm.LookupByAddress(address)
	} else if host != "" {
		mapping, ok = hm.LookupByHost(host)
	} else {
		return errors.New("must specify address or host")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Address", "Host"})

	if ok {
		for _, host := range mapping.Hosts {
			table.Append([]string{mapping.Address, host})
		}
	}

	table.Render()

	return nil
}
