package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom-power/depl/depl"
)

func main() {
	script, err := depl.ScriptFor(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(script)
}
