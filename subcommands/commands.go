package subcommands

import (
	"fmt"
	"manager/common"
	"manager/docker"
	"os"
)

var subCommands = map[string]Cmd{
	"show":  show,
	"exit":  exit,
	"quit":  exit,
	"help":  help,
	"set":   set,
	"print": list,
	"save":  save,
	"add":   add,
	"del":   del,
}

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

// save the changes to the docker config file
func save(c *docker.Config) error {
	return nil
}

// print config file
func list(c *docker.Config) error {
	m := c.GetAll()
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
	return nil
}

// reset docker config path
// TODO 执行 set 命令后，需要将新的docker配置文件路径更新到 manager.conf
// TODO 在该函数中增加修改配置文件的代码，将配置更新到 manager.conf 文件中
func set(c *docker.Config) error {

	l, err := common.Scanner("Input the docker config path: ")
	if err != nil {
		return err
	}

	path, err := c.SetPath(l)
	if err != nil {
		return err
	}

	fmt.Printf("reset config path successful, new path: %s\n", path)

	// TODO 在该位置更新 manager.conf

	return nil
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

// TODO 学习切片的相关操作，增减、删除元素等，排序等算法
// add a docker insecure registry
func add(c *docker.Config) error {
	return nil
}

// delete a docker insecure registry
func del(c *docker.Config) error {
	return nil
}

// change value for docker config
func change(c *docker.Config) error {
	return nil
}
