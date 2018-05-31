/*
Copyright Hill. All Rights Reserved.

*/

package cmd

import (
	"os"

	"github.com/yubingfeng/oasis/cli/cmd/cli/chaincode"
	"github.com/yubingfeng/oasis/cli/cmd/cli/channel"
	cliconfig "github.com/yubingfeng/oasis/cli/cmd/cli/config"
	"github.com/yubingfeng/oasis/cli/cmd/cli/event"
	"github.com/yubingfeng/oasis/cli/cmd/cli/query"
	"github.com/spf13/cobra"
)

func newFabricCLICmd() *cobra.Command {

	mainCmd := &cobra.Command{
		Use: "fabric-cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	flags := mainCmd.PersistentFlags()
	cliconfig.InitConfigFile(flags)
	cliconfig.InitLoggingLevel(flags)
	cliconfig.InitUserName(flags)
	cliconfig.InitUserPassword(flags)
	cliconfig.InitOrdererTLSCertificate(flags)
	cliconfig.InitPrintFormat(flags)
	cliconfig.InitWriter(flags)
	cliconfig.InitBase64(flags)
	cliconfig.InitOrgIDs(flags)

	mainCmd.AddCommand(chaincode.Cmd())
	mainCmd.AddCommand(query.Cmd())
	mainCmd.AddCommand(channel.Cmd())
	mainCmd.AddCommand(event.Cmd())

	return mainCmd
}

func Execute() {
	if newFabricCLICmd().Execute() != nil {
		os.Exit(1)
	}
}
