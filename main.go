package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/go-template/startup"
)

var services = execution_context.Get().Services

func main() {
	services.Version.Name = "GoLang Template"
	services.Version.Author = "Carlos Lapao"
	services.Version.License = "MIT"

	services.Version.Major = 0
	services.Version.Minor = 0
	services.Version.Build = 0
	services.Version.Rev = 1

	getVersion := helper.GetFlagSwitch("version", false)
	if getVersion {
		format := helper.GetFlagValue("o", "json")
		switch strings.ToLower(format) {
		case "json":
			fmt.Println(services.Version.PrintVersion(int(version.JSON)))
		case "yaml":
			fmt.Println(services.Version.PrintVersion(int(version.JSON)))
		default:
			fmt.Println("Please choose a valid format, this can be either json or yaml")
		}
		os.Exit(0)
	}

	services.Version.PrintAnsiHeader()

	configFile := helper.GetFlagValue("config", "")
	if configFile != "" {
		services.Logger.Command("Loading configuration from " + configFile)
		services.Configuration.LoadFromFile(configFile)
	}

	defer func() {
	}()

	startup.Init()
}
