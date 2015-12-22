// main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	results := make(map[string]float64)

	//method 1
	start := time.Now()
	args := os.Args[0:]
	for i := 0; i < 100; i++ {
		for n, arg := range args {
			fmt.Printf("arg[%d]: %v\n", n, arg)
		}
	}
	results["method 1"] = time.Since(start).Seconds()

	//method 2
	start = time.Now()
	for i := 0; i < 100; i++ {
		fmt.Println(strings.Join(os.Args[0:], " "))
	}
	results["method 2"] = time.Since(start).Seconds()

	//method 3
	start = time.Now()
	for j := 0; j < 100; j++ {
		s, sep := "", ""
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}
	results["method 3"] = time.Since(start).Seconds()

	for k, v := range results {
		fmt.Printf("%v, time elapsed: %v\n", k, v)
	}
}
