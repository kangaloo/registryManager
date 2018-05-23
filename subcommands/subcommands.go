package subcommands

import (
	"bufio"
	"fmt"
	"log"
	"manager/common"
	"manager/docker"
	"os"
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
	fmt.Printf("Input the docker config path: ")
	l, err := common.Scaner()
	if err != nil {
		return err
	}
	return c.SetPath(l)
}

func show(c *docker.Config) error {

	reg := c.Get("insecure-registries")
	for _, v := range reg {
		fmt.Println(v)
	}

	return nil

}

func add() {
	fmt.Print("> ")
	//iterm := readStdin()

}

func del() {

}

type cmd func(config *docker.Config) error

var subCommands = map[string]cmd{
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

// TODO 针对registries实现 增加、删除、修改、查看四个功能
func CmdScanner(conf *docker.Config) error {

	// TODO 对每个输入进行检查、修改
	// TODO 每个输入都是不可靠的，需要去掉开头、结尾的空格、换行符等

	r := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">> ")
		l, _, err := r.ReadLine()
		if err != nil {
			// todo 此处应设置一个错误处理方式，使得出错后尽量不影响程序运行
			log.Fatalln(err)
		}

		cmd, ok := subCommands[string(l)]

		if !ok {
			fmt.Printf("command %s not exist\n", string(l))
			continue
		}

		err = cmd(conf)
		if err != nil {
			fmt.Println(err)
		}

	}

}
