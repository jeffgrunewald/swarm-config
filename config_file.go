package main

import (
	"fmt"
	"os"
)

// EnsureConfigFile ensures config file is present
func EnsureConfigFile(path string) (bool, error) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return false, err
		}
		defer file.Close()
	}
	return false, nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
