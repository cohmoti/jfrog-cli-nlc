package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
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

	_, err := doTranslate(conf)
	if err != nil {
		log.Error("Failed doTranslate() err: ", err)
		return err
	}
	return nil
}

func doTranslate(c *doConfiguration) (string, error) {
	log.Debug("Got this input string:", c.nlCommand)

	homePath := os.Getenv("TALK2FROG_MODEL_HOME")
	if homePath == "" {
		return "", errors.New(`missing "TALK2FROG_MODEL_HOME" environment variable`)
	}
	log.Output("Translating to command ...")
	cmd := exec.Command("python", "main.py", "--mode", "single", "--sentence", c.nlCommand, "--model_dir", filepath.Join(homePath, "src/model/run"), "--model_file", "model_step_2500.pt")
	cmd.Dir = homePath
	if output, err := cmd.Output(); err != nil {
		log.Error(output)
		return "", fmt.Errorf("an error occurred during python model execution: %v", err)
	} else {
		lines := strings.Split(string(output), "\n")
		result := ""
		for _, s := range lines {
			if strings.HasPrefix(s, "Result=") {
				result = s[len("Result="):]
				break
			}
		}

		log.Output("Result: " + result)

		if strings.HasPrefix(result, "jfrog ") {
			if coreutils.AskYesNo("Would you like to execute it now?", false) {
				result, err = execJfrogCli(result)
				if err != nil {
					log.Output(result)
				}
			}
		}
		return result, nil
	}
}

func execJfrogCli(resultCmd string) (string, error) {
	cmdArr := strings.Split(resultCmd, " ")
	command := exec.Command(cmdArr[0], cmdArr[1:]...)
	command.Stderr = os.Stderr
	//command.Stdout = os.Stdout
	command.Env = os.Environ()
	output, err := command.Output()
	return string(output), err
}
