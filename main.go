//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	packageVersion "windows-os-info/version"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "windows-os-info",
		Usage:   "CLI tool to echo windows os info scripts",
		Version: packageVersion.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "action",
				Usage: "support action: cpu_architecture / os_version / logicalDriveStrings / cpu_caption / cpu_NumberOfCores / cpu_NumberOfLogicalProcessors / memorychip_capacity / computersystem_totalphysicalmemory",
			},
		},
		Action: func(c *cli.Context) error {
			action := c.String("action")
			if len(action) == 0 {
				return fmt.Errorf("invalid value for --action: %s. Valid options are 'os_arch', 'os_version', 'cpu_Num'", action)
			}

			if action == "os_arch" {

			} else if action == "os_version" {

			} else if action == "cpu_Num" {
			} else {
				return fmt.Errorf("invalid value for --action: %s. Valid options are 'os_arch', 'os_version', 'cpu_Num'", action)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
