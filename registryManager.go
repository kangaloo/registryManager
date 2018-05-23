package main

import (
	"flag"
	"log"
	"manager/docker"
	"fmt"
	"manager/config"
	"manager/subcommands"
)

func main() {

	//defaultConf := "/etc/docker/daemon.json"

	var Config *config.Conf
	conf := flag.String("c", "", "manager config file")
	flag.Parse()

	var err error


	// 命令行指定的配置文件优先，没有指定尝试读取默认配置，
	// 默认配置不存在则要求输入docker配置文件的路径，并保存在默认配置中
	if *conf != "" {
		Config, err = config.New(*conf)
	}

	if err != nil || *conf == "" {
		Config, err = config.NewDefault()
	}

	if err != nil {
		Config, err = config.Create()
	}



	/*
	if *conf == "" {
		Config, err = config.NewDefault()
	} else {
		Config, err = config.New(*conf)
	}

	*/

	if err != nil {
		log.Fatalln(err)
	}

	path := Config.ReadPath()

	dockerConf, err := docker.New(path)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(dockerConf)

	fmt.Printf("All commands are: %s\n", subcommands.AllCmds)

	err = subcommands.CmdScaner(dockerConf)
	if err != nil {
		log.Fatalln(err)
	}

	//subcommands.List(dockerConf)

}
