package main

import (
	"log"
	"os"
	"sort-manager/src/sortDir"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatal("Directory not specified")
	}
	nameDir := os.Args[1]

	err := sortDir.SortDirectory(nameDir)
	if err != nil {
		log.Println(err)
	}

}
