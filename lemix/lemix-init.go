package lemix

import "fmt"

func Setup(path string) bool {
	fmt.Println("path is :" + path)
	exists, _ := PathExists("hello")
	return exists
}