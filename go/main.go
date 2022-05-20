package main

import (
	"./flanalystsh"
	"fmt"
	"log"
	"os"
)

func main() {
	sh, err := flanalystsh.GetCommand(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sh)
}
