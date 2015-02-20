package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"errors"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/skratchdot/open-golang/open"
)

type serviceInstanceResponse struct {
	Entity serviceInstance `json:"entity"`
}

type serviceInstance struct {
	DashboardURL string `json:"dashboard_url"`
}

func main() {
	plugin.Start(&OpenPlugin{})
}

// OpenPlugin empty struct for plugin
type OpenPlugin struct{}

// Run of seeder plugin
func (plugin OpenPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	err := checkArgs(cliConnection, args)
	if err != nil {
		os.Exit(1);
	}
	if args[0] == "open" {
		plugin.runAppOpen(cliConnection, args)
	} else if args[0] == "service-open" {
		plugin.runServiceOpen(cliConnection, args)
	}
}

// GetMetadata of plugin
func (OpenPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    "open",
		Version: plugin.VersionType{Major: 1, Minor: 0},
		Commands: []plugin.Command{
			{
				Name:     "open",
				HelpText: "open app url in browser",
				UsageDetails: plugin.Usage{
					Usage: "open <appname>",
				},
			},
			{
				Name:     "service-open",
				HelpText: "open service instance dashboard in browser",
				UsageDetails: plugin.Usage{
					Usage: "service-open <servicename>",
				},
			},
		},
	}
}

func (plugin OpenPlugin) runAppOpen(cliConnection plugin.CliConnection, args []string) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput("app", args[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, "error: app does not exist")
		os.Exit(1)
	}
	var url string
	for _, line := range output {
		splitLine := strings.Split(strings.TrimSpace(line), " ")
		if splitLine[0] == "urls:" {
			if len(splitLine) > 1 {
				url = "http://" + strings.Trim(splitLine[1], ",")
				url = strings.TrimSpace(url)
			}
		}
	}
	err = checkRoutes(url)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	open.Run(url)

}

func (plugin OpenPlugin) runServiceOpen(cliConnection plugin.CliConnection, args []string) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput("service", args[1], "--guid")
	if err != nil {
		fmt.Fprintln(os.Stdout, "error: service does not exist")
		os.Exit(1)
	}
	serviceInstanceGUID := strings.TrimSpace(output[0])

	output, err = cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v2/service_instances/%s", serviceInstanceGUID))
	if err != nil {
		fmt.Fprintln(os.Stdout, "error: service does not exist")
		os.Exit(1)
	}
	jsonStr := ""
	for _, line := range output {
		jsonStr += line + "\n"
	}

	response := serviceInstanceResponse{}
	json.Unmarshal([]byte(jsonStr), &response)

	url := response.Entity.DashboardURL
	if url == "" {
		fmt.Println("No dashboard available")
	} else {
		open.Run(url)
	}
}

func checkArgs(cliConnection plugin.CliConnection, args []string) error{
	if len(args) < 2  {
		if args[0] == "open" {
			cliConnection.CliCommand(args[0], "-h")
			return errors.New("Appname is needed")
		} else if args[0] == "service-open" {
			cliConnection.CliCommand(args[0], "-h")
			return errors.New("Appname is needed")
		}
	}
	return nil
}

func checkRoutes(url string) error{
	if len(url) > 0 {
		return nil
	}
	return errors.New("error: app does not have any routes")
}
