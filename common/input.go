package common

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner(prompt string) (string, error) {
	fmt.Printf(prompt)
	r := bufio.NewReader(os.Stdin)
	l, _, err := r.ReadLine()
	if err != nil {
		return "", err
	}
	return string(l), nil
}
