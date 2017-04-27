package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	fmt.Fprintf(w, "Hello\n")
}
