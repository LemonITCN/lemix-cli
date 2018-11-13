package main

import (
	"fmt"
	"lemix-cli/lemix"
)

func main(){
	fmt.Print("hello\n")
	fmt.Println(lemix.Setup("hello111path"))
}