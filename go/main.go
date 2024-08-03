package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom-power/depl/depl"
)

func main() {
	script, err := depl.Run(command(), explain())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(script)
}

func command() string {
	return os.Args[1]
}

func explain() bool {
	return depl.Contains(os.Args, "--explain")
}
