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
				Usage: "support action: os_arch / os_version / cpu_num",
			},
		},
		Action: func(c *cli.Context) error {
			action := c.String("action")
			if len(action) == 0 {
				return fmt.Errorf("invalid value for --action: %s. Valid options are 'os_arch', 'os_version', 'cpu_num'", action)
			}

			if action == "os_arch" {
				arch, err := getOSArch()
				if err != nil {
					return err
				}
				fmt.Print(arch)
			} else if action == "os_version" {
				osVersion, err := GetOSVersion()
				if err != nil {
					return err
				}
				fmt.Print(osVersion)
			} else if action == "cpu_num" {
				cpuNum, err := getCPUNum()
				if err != nil {
					return err
				} else {
					fmt.Print(cpuNum)
				}
			} else {
				return fmt.Errorf("invalid value for --action: %s. Valid options are 'os_arch', 'os_version', 'cpu_num'", action)
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
