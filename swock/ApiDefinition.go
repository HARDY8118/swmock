package swock

import (
	"fmt"
)

type OpenApi3Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type OpenApi3Server struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type OpenApi3Parameter struct {
	In       string
	Name     string
	Required bool
	Schema   struct {
		Type        string
		Description string
		Example     any
	}
}

type OpenApi3Path struct {
	Get  OpenApi3Get  `json:"get"`
	Post OpenApi3Post `json:"post"`
}

type OpenApi3Paths map[string]OpenApi3Path

type OpenApi3Definition struct {
	OpenApi string           `json:"openapi"`
	Info    OpenApi3Info     `json:"info"`
	Servers []OpenApi3Server `json:"servers"`
	Paths   OpenApi3Paths    `json:"paths"`
}

func (definition *OpenApi3Definition) Print() {
	fmt.Printf("Open API version: %s\n", definition.OpenApi)
	fmt.Println("Info")
	fmt.Printf("  Title: %s\n", definition.Info.Title)
	fmt.Printf("  Description: %s\n", definition.Info.Description)
	fmt.Printf("  Version: %s\n", definition.Info.Version)
	fmt.Println("Servers")

	for _, server := range definition.Servers {
		fmt.Println("  ┌───")
		fmt.Printf("  │URL: %s\n", server.Url)
		fmt.Printf("  │Description: %s\n", server.Description)
		fmt.Println("  └───────────")
	}
}
