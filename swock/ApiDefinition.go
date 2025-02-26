package swock

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

const SERVER_DESC_LEN = 60

type OpenApi3Info struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
}

type OpenApi3Server struct {
	Url         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
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

type OpenApi3Path map[string]OpenApi3Method

type OpenApi3Paths map[string]OpenApi3Path

func (oa3p OpenApi3Paths) PathsList() []string {
	var keys []string
	hasRoot := false

	for key, _ := range oa3p {
		if key == "/" {
			hasRoot = true
			continue
		}
		keys = append(keys, key)
	}

	if hasRoot {
		keys = append(keys, "/")
	}

	return keys
}

type OpenApi3Definition struct {
	OpenApi string           `json:"openapi"`
	Info    OpenApi3Info     `json:"info"`
	Servers []OpenApi3Server `json:"servers"`
	Paths   OpenApi3Paths    `json:"paths"`
}

func NewOpenApiDefinition(filePath string) OpenApi3Definition {
	var openApiDefinition OpenApi3Definition

	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("File does not exists")
		} else {
			log.Fatal(err)
		}

		os.Exit(1)
	}

	if strings.HasSuffix(filePath, "json") {
		err = json.Unmarshal(fileContent, &openApiDefinition)
		if err != nil {
			log.Fatalf("Failed to parse json, %s", err)
		}
	} else if strings.HasSuffix(filePath, "yaml") {
		err = yaml.Unmarshal(fileContent, &openApiDefinition)
		if err != nil {
			log.Fatalf("Failed to parse yaml, %s", err)
		}
	} else {
		log.Fatalf("Unsupported file: %s", filePath)
	}

	return openApiDefinition
}

func (oad OpenApi3Definition) Init() {
	log.Printf("OpenAPI version: %s", oad.OpenApi)
	log.Printf("Title: %s", oad.Info.Title)
	log.Printf("Description: %s", oad.Info.Description)
	log.Printf("Version: %s", oad.Info.Version)

	for _, server := range oad.Servers {
		description := server.Description
		if len(description) > SERVER_DESC_LEN {
			description = description[:SERVER_DESC_LEN-3] + "..."
		}
		log.Printf("Server: %s (%s)", server.Url, description)
	}
}

func attachHandler(mux *mux.Router, path string, methods OpenApi3Path) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		methodDefinition := methods
		methodLower := strings.ToLower(r.Method)

		if m, ok := methodDefinition[methodLower]; ok {
			randStatus := RandSelect(m.Responses.StatusCodes())
			randSchema := RandSelect(m.Responses[randStatus].Content.ContentTypes())

			response := m.Responses[randStatus].Content[randSchema].Example
			intStatus, _ := strconv.Atoi(randStatus)

			w.Header().Set("Content-Type", randSchema)
			w.WriteHeader(intStatus)
			w.Write(response)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		log.Printf("%s %s", r.Method, r.URL)
	})
}

func (oad OpenApi3Definition) Start() {
	mux := mux.NewRouter()

	for _, pathEntry := range oad.Paths.PathsList() {
		path := pathEntry
		methods := oad.Paths[path]

		attachHandler(mux, path, methods)
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(oad.Servers))

	for _, server := range oad.Servers {
		// go serve(server.Url, mux, wg)
		addr := Addr(server.Url)
		go func() {
			log.Printf("Starting server: %s", addr)
			err := http.ListenAndServe(addr, mux)
			if err != nil {
				log.Fatalf("Failed to start server %s, %s", addr, err)
			}
			log.Printf("Started server: %s", addr)
			wg.Done()
		}()
	}

	wg.Wait()
}
