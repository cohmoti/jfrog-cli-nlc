package commands

import (
	"errors"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func GetDoCommand() components.Command {
	return components.Command{
		Name:        "do",
		Description: "Translate English prompt to jfrog cli command. For instance, \"Audit the Go project at the current directory using the watch1 watch defined in Xray\" ==> jfrog xr ago --watches \"watch1\"",
		Aliases:     []string{"do"},
		Action: func(c *components.Context) error {
			return doCmd(c)
		},
	}
}

type doConfiguration struct {
	nlCommand string
}

func doCmd(c *components.Context) error {

	if len(c.Arguments) != 1 {
		return errors.New("wrong number of arguments. Expecting quoted English command description")
	}
	var conf = new(doConfiguration)
	//conf.nlCommand = strings.Join(c.Arguments, " ")
	conf.nlCommand = c.Arguments[0]

	log.Output(doTranslate(conf))
	return nil
}

func doTranslate(c *doConfiguration) string {
	log.Debug("Got this input string:", c.nlCommand)
	result := "jfrog xr ago --watches \"watch1\""
	return result
}
