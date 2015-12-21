// main.go
package main

import (
	"fmt"
	"os"
	//"strings"
)

func main() {
	//fmt.Println(strings.Join(os.Args[0:], " "))
	args := os.Args[0:]

	for n, arg := range args {
		fmt.Printf("arg[%d]: %v\n", n, arg)
	}
}
