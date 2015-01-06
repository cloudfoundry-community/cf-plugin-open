package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	plugin.Start(&OpenPlugin{})
}

//OpenPlugin empty struct for plugin
type OpenPlugin struct{}

//Run of seeder plugin
func (plugin OpenPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput("app", args[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, "error: app does not exist")
		os.Exit(1)
	}
	var url string
	for _, line := range output {
		splitLine := strings.Split(strings.TrimSpace(line), " ")
		if splitLine[0] == "urls:" {
			url = "http://" + strings.Trim(splitLine[1], ",")
			url = strings.TrimSpace(url)
		}
	}
	open.Run(url)
}

//GetMetadata of plugin
func (OpenPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "open",
		Commands: []plugin.Command{
			{
				Name:     "open",
				HelpText: "open app url in browser",
			},
		},
	}
}
