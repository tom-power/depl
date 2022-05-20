package main

import (
	"./flanalystsh"
	"fmt"
	"os"
)

func main() {
	fmt.Printf(flanalystsh.GetCommand(os.Args[1]))
}
