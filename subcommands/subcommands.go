package subcommands

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/c-bata/go-prompt"
	"manager/common"
	"manager/docker"
	"os"
	"strings"
)

const AllCmds = `show set del exit`

const helpStr = `	help: show this help
	exit: exit the manager
	quit: same as exit
	set:  set the docker config path
	show: show registries
	`

func help(c *docker.Config) error {

	fmt.Println(helpStr)

	return nil
}

func exit(c *docker.Config) error {
	os.Exit(0)
	return nil
}

/*
const SubCommands  map[string]func() = map[string]func(){
	"save": save,
}

*/
func save() {

}

// set docker config path
func set(c *docker.Config) error {
	l, err := common.Scanner("Input the docker config path: ")
	if err != nil {
		return err
	}
	return c.SetPath(l)
}

// print docker insecure registries
func show(c *docker.Config) error {

	reg := c.Get("insecure-registries")

	if v, ok := reg.([]interface{}); ok {
		for _, i := range v {
			fmt.Println(i)
		}
	} else {
		fmt.Println("not string slice .")
		fmt.Println(reg)
	}

	return nil
}

// add a docker insecure registry
func add() {
	fmt.Print("> ")
	//iterm := readStdin()

}

// delete a docker insecure registry
func del() {

}

type Cmd func(config *docker.Config) error

var subCommands = map[string]Cmd{
	"show": show,
	"exit": exit,
	"quit": exit,
	"help": help,
	"set":  set,
	//"save": save,
	//"add":  add,
	//"del":  del,
	//"exit": exit,
}

/*
func CMDParser() {
	r := bufio.NewReader(os.Stdin)
	c, _, _ := r.ReadLine()

	cmd := string(c)

	commands := subCommands[cmd]
	commands()

}
*/

func readStdin() string {
	r := bufio.NewReader(os.Stdin)
	l, _, _ := r.ReadLine()
	return string(l)
}

// check if the command exist
func CmdChecker(c string) (Cmd, error) {
	cmd, ok := subCommands[c]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s command not exist", c))
	}
	return cmd, nil
}

// TODO 针对registries实现 增加、删除、修改、查看四个功能
func CmdScanner(conf *docker.Config) error {

	// TODO 对每个输入进行检查、修改
	// TODO 每个输入都是不可靠的，需要去掉开头、结尾的空格、换行符等

	//r := bufio.NewReader(os.Stdin)

	for {

		l := prompt.Input(">> ", completer)
		l = strings.TrimSpace(l)

		/*
			fmt.Print(">> ")
			l, _, err := r.ReadLine()
			if err != nil {
				// TODO 此处应设置一个错误处理方式，使得出错后尽量不影响程序运行
				log.Fatalln(err)
			}
		*/

		cmd, err := CmdChecker(l)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = cmd(conf)
		if err != nil {
			fmt.Println(err)
		}
	}
}
