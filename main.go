package main

import (
	"log"
	"os"

	"github.com/chonlatee/nen/cmd"
)

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Printf("command failed: %v \n", err)
		os.Exit(1)
	}
}
