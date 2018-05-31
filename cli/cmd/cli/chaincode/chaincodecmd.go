/*
Copyright hill. All Rights Reserved.
*/

package chaincode

import (
	cliconfig "github.com/yubingfeng/oasis/cli/cmd/cli/config"
	"github.com/spf13/cobra"
)

var chaincodeCmd = &cobra.Command{
	Use:   "chaincode",
	Short: "Chaincode commands",
	Long:  "Chaincode commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd returns the chaincode command
func Cmd() *cobra.Command {
	cliconfig.InitChannelID(chaincodeCmd.Flags())

	chaincodeCmd.AddCommand(getInstallCmd())
	chaincodeCmd.AddCommand(getInstantiateCmd())
	chaincodeCmd.AddCommand(getInvokeCmd())
	chaincodeCmd.AddCommand(getQueryCmd())
	chaincodeCmd.AddCommand(getGetInfoCmd())
	chaincodeCmd.AddCommand(getUpgradeCmd())

	return chaincodeCmd
}
