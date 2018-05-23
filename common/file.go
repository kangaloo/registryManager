package common

import (
	"io"
	"os"
	"fmt"
)

func CopyFile(dst, src string) (int64, error) {

	fi, err := os.Stat(src)
	fmt.Println("----")
	if err != nil {
		return 0, err
	}

	_, err1 := os.Stat(dst)
	if err1 == nil {
		// 该类型的 ERROR() 方法接收的是指针，所以此处需要返回指针
		// 若该方法的 ERROR() 方法接收的是值，则此处可返回值
		return 0, &ExistError{dst}
	}

	s, err := os.Open(src)
	defer s.Close()
	if err != nil {
		return 0, err
	}

	d, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, fi.Mode())
	defer d.Close()
	if err != nil {
		return 0, err
	}

	return io.Copy(d, s)
}
