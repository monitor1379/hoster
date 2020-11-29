package cmds

/*
 * @Date: 2020-11-29 22:17:45
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 22:55:20
 */

import (
	"fmt"
	"os"

	"github.com/monitor1379/hoster"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hoster",
		Short: "Hoster is a cross-platform operating system host file management library written in Go.",
	}

	sysHostFilePath, err := hoster.GetSysHostFilePath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s", err.Error())
	}

	rootCmd.PersistentFlags().StringP("file", "f", sysHostFilePath, "specify a host file path")

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newLookupCommand())
	rootCmd.AddCommand(newSetCommand())

	return rootCmd
}
