package main

import (
	"./depl"
	"fmt"
	"log"
	"os"
)

func main() {
	sh, err := depl.GetCommand(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sh)
}
