package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		str := ""
		sep := ""
		for _, arg := range os.Args[1:] {
			str += arg + sep
			sep = " "
		}
	*/
	str := strings.Join(os.Args[1:], " ")
	fmt.Println(str)
}
