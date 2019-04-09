package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: filename required")
		return
	}

	filename := os.Args[1]
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error: cannot open file %s (%v)", filename, err)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		fmt.Println(line)
		fmt.Println(line)
	}
	f.Close()
}
