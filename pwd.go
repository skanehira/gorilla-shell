package main

func pwd(args []interface{}) (interface{}, error) {
	return currentDir(), nil
}
