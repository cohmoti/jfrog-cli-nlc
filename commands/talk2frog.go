package commands

import (
	"errors"
	"fmt"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	result, err := doTranslate(conf)
	if err != nil {
		log.Error("Failed doTranslate() err: ", err)
		return err
	}
	log.Output(result)
	return nil
}

func doTranslate(c *doConfiguration) (string, error) {
	log.Debug("Got this input string:", c.nlCommand)

	homePath := os.Getenv("TALK2FROG_MODEL_HOME")
	if homePath == "" {
		return "", errors.New(`missing "TALK2FROG_MODEL_HOME" environment variable`)
	}
	scriptPath := filepath.Join(homePath, "query_script.py")
	if output, err := exec.Command("python", scriptPath, c.nlCommand).Output(); err != nil {
		return "", fmt.Errorf("an error occurred during python model execution: %v", err)
	} else {
		// result := "jfrog xr ago --watches \"watch1\""
		result := strings.TrimSuffix(string(output), "\n")
		log.Debug("Got this output string from python model:", result)
		return result, nil
	}
}
