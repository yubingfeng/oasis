/*

Copyright 2018 hill

*/

package util

import   (

	"fmt"
	"github.com/spf13/cobra"

)


//SubcmdRunE returns a function that handles a case where a subcommand must be specified

func SubCmdRunE(name string) func(*cobra.Command, []string) error{

	return func(_ *cobra.Command, args  []string) error{

		if len(args) < 1 {
			return fmt.Errorf("missing subcommand, %q is not mean to run its own", name)
		}

		return fmt.Errorf("invalid subcommand: %q", args[0])
        }
}
