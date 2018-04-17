/*
Copyright 2018 hill
*/


package app

import (
        "os"
	"github.com/yubingfeng/oasis/cmd/deployer/app/cmd"
        "github.com/spf13/pflag"
)


func Run()  error {

        pflag.CommandLine.MarkHidden("version")
	cmd :=cmd.NewDeployerCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()

}

