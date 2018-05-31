/*
Copyright Hill. All Rights Reserved.

*/

package query

import (
	cliconfig "github.com/yubingfeng/oasis/cli/cmd/cli/config"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query commands",
	Long:  "Query commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd returns the query command
func Cmd() *cobra.Command {
	cliconfig.InitChannelID(queryCmd.Flags())

	queryCmd.AddCommand(getQueryBlockCmd())
	queryCmd.AddCommand(getQueryInfoCmd())
	queryCmd.AddCommand(getQueryTXCmd())
	queryCmd.AddCommand(getQueryChannelsCmd())
	queryCmd.AddCommand(getQueryInstalledCmd())

	return queryCmd
}
