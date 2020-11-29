package cmds

/*
 * @Date: 2020-11-29 22:32:10
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 22:50:36
 */

import (
	"fmt"

	"github.com/monitor1379/hoster"
	"github.com/spf13/cobra"
)

func newSetCommand() *cobra.Command {
	setCommand := &cobra.Command{
		Use:   "set",
		Short: "Add a address-host mapping",
		RunE:  setCommandRunE,
	}
	setCommand.PersistentFlags().String("host", "", "host name")
	setCommand.PersistentFlags().String("address", "", "address of ip")
	setCommand.PersistentFlags().String("comment", "", "comment")

	return setCommand
}

func setCommandRunE(c *cobra.Command, args []string) error {
	hostFilePath := c.Parent().PersistentFlags().Lookup("file").Value.String()
	hm, err := hoster.New(hostFilePath)
	if err != nil {
		return err
	}

	address := c.PersistentFlags().Lookup("address").Value.String()
	host := c.PersistentFlags().Lookup("host").Value.String()
	comment := c.PersistentFlags().Lookup("comment").Value.String()
	if comment != "" && []rune(comment)[0] != '#' {
		comment = fmt.Sprintf("#%s", comment)
	}

	err = hm.Set(host, address, comment)
	if err != nil {
		return err
	}

	return nil
}
