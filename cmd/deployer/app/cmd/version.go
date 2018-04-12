/*
Copyright 2018 hill
*/

package cmd

import (
        "fmt"
	"io"
	"os"
        "github.com/golang/glog"
)


const (
        VERSION =  "1.0"
)

func NewCmdVersion(out io.Writer) *cobra.Command {

        cmd := &cobra.Command{
                Use:    "Version"
                Short:  "Show the version of deployer"
                Run:    func(cmd *cobra.Command, args []string){
                        err := RunVersion(out, cmd)
			if err != nil {
				fmt.Fprintf(out, "run version error %s \n", err)
				os.Exit(1)
			}
        }

}

func RunVersion(out io.Writer,  cmd *cobra.Command) error {
        glog.V(1).Infoln("[version] retrieving version info")
        fmt.Fprintf("Deployer version: %s", VERSION)
        return nil
}
