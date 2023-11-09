package main

import (
	"log"

	"github.com/kolbymcgarrah/cli-todo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
