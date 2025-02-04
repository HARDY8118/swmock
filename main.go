package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	swock "github.com/HARDY8118/swock/swock"
)

func main() {

	definition := swock.ParseSwagger("swagger.json")

	for _, server := range definition.Servers {
		if !swock.ValidateUrl(server.Url) {
			log.Fatalf("Invalid URL: %s", server.Url)
		}
	}

	definition.Print()

	mux := http.NewServeMux()
	handler := swock.NewSwockHandler()

	for path, method := range definition.Paths {
		fmt.Printf("Found path: %s\n", path)
		handler.AddPath(path, method)
		mux.HandleFunc(path, handler.Handle)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(definition.Servers))

	for _, server := range definition.Servers {
		url := server.Url
		description := server.Description
		go func() {
			fmt.Printf("Creating server: %s (%s...)\n", url, description[:16])
			err := http.ListenAndServe(swock.Addr(url), mux)
			if err != nil {
				fmt.Printf("Failed to start server: %s\n%s\n", server.Url, err)
			}
			fmt.Printf("Server listening on %s", url)
			wg.Done()
		}()
	}

	wg.Wait()
}
