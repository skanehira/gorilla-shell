package main

import (
	"os"
)

func cd(args []interface{}) (interface{}, error) {
	if len(args) > 0 {
		err := os.Chdir(args[0].(string))
		return "", err
	}
	return "", nil
}
