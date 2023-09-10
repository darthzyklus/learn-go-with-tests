package main

import (
	"fmt"
	"io"
)

func Greeter(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
