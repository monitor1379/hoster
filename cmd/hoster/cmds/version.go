package cmds

/*
 * @Date: 2020-11-29 22:20:45
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 22:31:47
 */

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.1.0"
)

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of hoster",
		Run:   versionCmdRun,
	}
	return versionCmd
}

func versionCmdRun(c *cobra.Command, args []string) {
	fmt.Printf("hoster version %s", version)
}
