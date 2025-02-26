package main

import (
	"fmt"
	"os"

	"github.com/HARDY8118/swock/swock"
)

func main() {

	argc := len(os.Args)

	if argc < 2 {
		fmt.Println("Too few arguments")
		fmt.Println("Usage: swmock <swagger_file>")
		os.Exit(0)
	} else if argc > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	openApiDefinition := swock.NewOpenApiDefinition(os.Args[1])
	openApiDefinition.Init()
	openApiDefinition.Start()

}
