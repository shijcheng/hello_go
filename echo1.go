// Echo1 Exercise 1.1

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	// Exercise 1.1
	//fmt.Println(os.Args)
}