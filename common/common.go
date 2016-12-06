package common

import "os"

func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
