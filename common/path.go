package common

import (
	"path/filepath"
	"os"
)

func GetCurrentPath() (string, error) {

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return path, nil
}
