package main

import (
	"fmt"
	"github.com/tom-power/depl/depl"
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
