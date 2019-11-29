package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func ls(args []interface{}) (interface{}, error) {
	current, err := os.Getwd()
	if err != nil {
		printErr(err)
		return nil, err
	}

	files, err := ioutil.ReadDir(current)
	if err != nil {
		printErr(err)
		return nil, err
	}

	var names []string
	for _, f := range files {
		names = append(names, f.Name())
	}

	if len(args) > 0 {
		if args[0] == "-l" {
			return strings.Join(names, "\r\n"), nil
		}
	}

	return strings.Join(names, " "), nil
}
