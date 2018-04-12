/*
  Copyright 2018 hill
*/

package cmd

import (

        "io"
)

func NewDeployerCommand(_ io.Reader, out, err io.Writer)  *cobra.Command {

        cmds := &cobra.Command{
                Use:     "deployer",
		Short:   "deployer: deploy Oasis applications",
                Long:    "deployer: deploy Oasis applications",
        }

        cmds.ResetFlags()
	cmds.SetGlobalNormalizationFunc(flag.WarnWordSepNormailizeFunc)
	cmds.AddCommand(NewCmdVersion(out))

        return cmds
}
