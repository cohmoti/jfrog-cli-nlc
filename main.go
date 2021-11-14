package main

import (
	"github.com/cohmoti/jfrog-cli-nlc/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "talk2frog"
	app.Description = "talk2frog brings the power of natural language processing to the JFrog command line. It transforms descriptions of command line tasks in English to their jfrog CLI syntax."
	app.Version = "v1.0.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetDoCommand()}
}
