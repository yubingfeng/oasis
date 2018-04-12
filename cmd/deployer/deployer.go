/*
Copyright 2018 hill
*/

package main

import  (
        "fmt"
        "os"
        "github.com/yubingfeng/oasis/cmd/deployer/app"

)

func main() {

        if err := app.Run(); err != nil {
                fmt.Fprintf(os.Stderr, "Run Deployer Error: %s \n", err)
                os.Exit(1)
        }
	os.Exit(0)
}
