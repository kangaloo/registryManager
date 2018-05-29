package subcommands

import (
	"errors"
	"fmt"
	"github.com/c-bata/go-prompt"
	"manager/docker"
	"strings"
)

type Cmd func(config *docker.Config) error

// check if the command exist
func CmdChecker(c string) (Cmd, error) {

	if c == "" {
		return nil, nil
	}

	cmd, ok := subCommands[c]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s command not exist", c))
	}
	return cmd, nil
}

// TODO 针对registries实现 增加、删除、修改、查看四个功能
func CmdScanner(conf *docker.Config) error {

	for {
		l := prompt.Input(">> ", completer)
		l = strings.TrimSpace(l)
		cmd, err := CmdChecker(l)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if cmd == nil {
			continue
		}

		err = cmd(conf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
}
