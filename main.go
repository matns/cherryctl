/*
Copyright © 2022 Cherry Severs <support@cherryservers.com>
*/
package main

import (
	"os"

	"github.com/cherryservers/cherryctl/cmd"
)

func main() {
	cli, err := cmd.NewCli()
	if err != nil {
		os.Exit(1)
	}
	if err := cli.MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
