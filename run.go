package main

import (
	"fmt"
	"io"
)

func run(args []string, stdout io.Writer) error {
	fmt.Println("hello world")
	return nil
}
