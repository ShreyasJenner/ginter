package main

import (
	"os"

	"ginter/repl"
)

func main() {

	repl.Start(os.Stdin, os.Stdout)
}
