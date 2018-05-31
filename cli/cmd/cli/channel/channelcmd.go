/*
Copyright Hill. All Rights Reserved.
*/

package channel

import (
	cliconfig "github.com/yubingfeng/oasis/cli/cmd/cli/config"
	"github.com/spf13/cobra"
)

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Channel commands",
	Long:  "Channel commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd returns the channel command
func Cmd() *cobra.Command {
	cliconfig.InitChannelID(channelCmd.Flags())

	channelCmd.AddCommand(getChannelCreateCmd())
	channelCmd.AddCommand(getChannelJoinCmd())

	return channelCmd
}
