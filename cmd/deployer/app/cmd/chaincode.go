/*
Copyright 2018 hill
*/

package cmd

import (

        "fmt"
	"io"
	"github.com/golang/glog"
        "github.com/spf13/cobra"
        cmdutil "github.com/yubingfeng/oasis/cmd/deployer/app/cmd/util"	
)

var (

        shortDescription = "Run this to install/init/query/ chaincode on specific channel"
	longDescription = "chaincode install/init/query channelname"
)

func NewCmdChaincode(out io.Writer) *cobra.Command{

	chaincodeCmd := &cobra.Command{
                Use:      "chaincode",
 		Short:    shortDescription,
	        Long:     longDescription,
                RunE:     cmdutil.SubCmdRunE("chaincode"), 
        }

	var channel string
	listCmd :=  &cobra.Command{
                Use:     "list",
		Short:   "List chaincode on specific channel",
		Long:    "List chaincode on specific channel",
		Run:     func(cmd *cobra.Command, args []string){
			glog.V(1).Infoln("list chaincode from channel")
                        fmt.Fprintln(out, "list chaincode %s", args)
		},
        }
	listCmd.Flags().StringVar(&channel, "channel", channel, "channel name of chain")
	chaincodeCmd.AddCommand(listCmd)

	return chaincodeCmd
}

