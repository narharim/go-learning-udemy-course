package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: program_name filename")
		return
	}
	filename := os.Args[1]

	fmt.Println("Reading file:", filename)

	file, err := os.Open(filename)
	if err != nil {
		return
	}

	io.Copy(os.Stdout, file)

}
