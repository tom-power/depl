package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tom-power/depl/depl"
)

func main() {
	script, err := depl.Run(os.Args[1], depl.Contains(os.Args, "--explain"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(script)
}
