package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom-power/depl/depl"
)

func main() {
	script, err := depl.Run(os.Args[1], contains(os.Args, "--explain"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(script)
}

func contains(args []string, target string) bool {
	for _, arg := range args {
		if arg == target {
			return true
		}
	}
	return false
}
