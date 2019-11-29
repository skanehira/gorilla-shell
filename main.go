package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err)
}

type cmdFunc func(args []interface{}) (interface{}, error)

var cmds map[string]cmdFunc

func init() {
	cmds = map[string]cmdFunc{
		"ls":  cmdFunc(ls),
		"cd":  cmdFunc(cd),
		"pwd": cmdFunc(pwd),
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ps()

	for scanner.Scan() {
		var cmd string
		var args []interface{}
		cmdText := strings.Split(scanner.Text(), " ")
		if len(cmdText) > 0 {
			cmd = cmdText[0]
		}
		if len(cmdText) > 1 {
			for _, c := range cmdText[1:] {
				args = append(args, c)
			}
		}

		if f, ok := cmds[cmd]; ok {
			res, err := f(args)
			if err != nil {
				printErr(err)
				goto ps
			}
			fmt.Println(res)
		}
	ps:
		ps()
	}
}

func ps() {
	fmt.Printf("%s>", filepath.Base(currentDir()))
}

func currentDir() string {
	current, err := os.Getwd()
	if err != nil {
		return ""
	}

	return current
}
