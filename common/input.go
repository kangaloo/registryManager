package common

import (
	"bufio"
	"os"
)

func Scaner() (string, error) {
	r := bufio.NewReader(os.Stdin)
	l, _, err := r.ReadLine()
	if err != nil {
		return "", err
	}
	return string(l), nil
}
