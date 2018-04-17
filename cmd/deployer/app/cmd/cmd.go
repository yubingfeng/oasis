/*
  Copyright 2018 hill
*/

package cmd

import (

        "io"
	"github.com/spf13/cobra"
        "github.com/yubingfeng/oasis/cmd/deployer/app/util/flag"
)

func NewDeployerCommand(_ io.Reader, out, err io.Writer)  *cobra.Command {

        cmds := &cobra.Command{
                Use:     "deployer",
		Short:   "deployer: deploy Oasis applications",
                Long:    "deployer: deploy Oasis applications",
        }

        cmds.ResetFlags()
	cmds.SetGlobalNormalizationFunc(flag.WarnWordSepNormalizeFunc)
	cmds.AddCommand(NewCmdVersion(out))
        cmds.AddCommand(NewCmdChaincode(out))

        return cmds
}
