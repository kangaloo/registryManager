package docker

import (
	"errors"
	"io/ioutil"
	"manager/common"
	"os"
	"strconv"
	"time"
)

//const DefaultPath ="/etc/docker/daemon.json"

type Config struct {
	path     string
	backPath string
	isBack   bool
	config   map[string]interface{}
}

func New(c string) (*Config, error) {

	_, err := os.Stat(c)
	if err != nil {
		return nil, err
	}

	var conf = &Config{path: c}
	err = conf.load()
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// read string from file and assigns the json to c.config
func (c *Config) load() error {

	f, err := os.Open(c.path)
	defer f.Close()
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	c.config, err = common.Json2map(bs)
	if err != nil {
		return err
	}

	return nil
}

// update c.config from c.path
func (c *Config) ReLoad() error {

	c.setBackPath()
	c.isBack = false
	err := c.load()

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) SetPath(path string) (string, error) {

	_, err := os.Stat(path)
	if err != nil {
		return "", errors.New("path not exist")
	}

	c.path = path
	err = c.ReLoad()
	if err != nil {
		return "", errors.New("")
	}

	return c.path, nil
}

func (c *Config) setBackPath() {

	t := int(time.Now().Unix())
	s := strconv.Itoa(t)
	c.backPath = c.path + "." + s

}

func (c *Config) Back() error {

	c.setBackPath()

	_, err := common.CopyFile(c.backPath, c.path)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) DelBack() error {

	return nil
}

/*
func loadConf(s string) (conf, error) {




	return
}

func dumpConf(c conf) error {





}


func (c conf) test()  {

}

*/

//func (c Config) Load()  {
//
//}

func (c *Config) Dump() error {

	s, err := common.Map2json(c.config)

	if err != nil {
		return err
	}

	// TODO truncate文件，危险操作，应在该操作前备份文件
	f, err := os.OpenFile(c.path, os.O_WRONLY|os.O_TRUNC, 0)
	defer f.Close()

	if err != nil {
		return err
	}
	f.Write(s)

	return nil

}

// get element from the map c.config
func (c *Config) Get(s string) interface{} {
	i, ok := c.config[s]
	if !ok {
		return nil
	}
	return i
}

// get all elements from the map c.config
func (c *Config) GetAll() map[string]interface{} {
	return c.config
}

// return c.path
func (c *Config) GetPath() string {
	return c.path
}

/*
func (c *Config) Append(s, iterm string) error {

	i, ok := c.config[s]

	if !ok {
		return errors.New(s + "not exist.")
	}


	c.config[s] = append([]byte(i), iterm)

}
*/
