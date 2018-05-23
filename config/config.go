package config

import (
	"bufio"
	"errors"
	"io/ioutil"
	"manager/common"
	"os"
	"fmt"
)

// var DefaultConf string

var defaultConf = "manager.conf"

type Conf struct {
	path    string
	content []byte
}

/*
func init()  {

	// TODO 该init函数有待验证
	// TODO common.GetCurrentPath() 函数有待验证
	pwd, err := common.GetCurrentPath()
	if err != nil {
		DefaultConf = ""
	}
	DefaultConf = pwd + string(os.PathSeparator) + "manager.conf"
}

func LoadDef()  {



}
*/

func (c *Conf) ReadPath() string {
	return string(c.content)
}

func defaultPath() (string, error) {
	cur, err := common.GetCurrentPath()
	if err != nil {
		return "", err
	}
	return cur + string(os.PathSeparator) + defaultConf, nil
}

func NewDefault() (*Conf, error) {

	path, err := defaultPath()
	if err != nil {
		return nil, err
	}

	conf, err := New(path)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func New(path string) (*Conf, error) {

	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		return nil, errors.New(path + " is a directory")
	}

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	// TODO ioutil.ReadAll() f.Read() 等方法会在每行的结尾读出一个换行符，码点值10，\n
	// TODO 好像 bufio.Readline 会自动去掉换行符
	cont, err := ioutil.ReadAll(f)

	// TODO go生成切片的规则？不支持负数索引？包含关系是怎样的？
	cont = cont[:len(cont)-1]

	if err != nil {
		return nil, err
	}

	return &Conf{path: path, content: cont}, nil
}

func Create() (*Conf, error) {

	path, err := defaultPath()

	if err != nil {
		return nil, err
	}

	_, err = os.Stat(path)

	if err == nil {
		return nil, common.NewExistErr(path)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Input the docker config path: ")
	// readline 方法去掉了换行符
	l, _, err := r.ReadLine()

	// 因后续调用的函数去掉换行符的方式为去掉切片的最后一个元素
	// 所以此处需要增加一个换行符，避免最后一个字符被舍弃
	// TODO 用编辑器写的文件，即使没有换行，文件末尾也一定有个行换行符？
	// TODO 程序写进去的则没有最后的换行符？
	l = append(l, byte(10))
	//l = append(l, l...)

	_, err = f.Write(l)

	if err != nil {
		return nil, err
	}

	return NewDefault()
}
