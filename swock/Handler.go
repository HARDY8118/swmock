package swock

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
)

type SwockResponse struct {
	Status int
	Header string
	Body   json.RawMessage
}

type SwockHandler struct {
	counter int
	// handlers map[string]map[string][]func(w http.ResponseWriter, r *http.Request) // {"/ping": "GET": () -> {}}
	handlers map[string]map[string][]SwockResponse
}

func NewSwockHandler() *SwockHandler {
	return &SwockHandler{
		0,
		make(map[string]map[string][]SwockResponse),
	}
}

func (handler *SwockHandler) AddPath(path string, method OpenApi3Path) {

	if !reflect.ValueOf(method.Get).IsZero() { // Is get request
		var responses []SwockResponse
		for status, response := range method.Get.Responses {
			istatus, _ := strconv.Atoi(status)
			for header, content := range response.Content {
				responses = append(responses, SwockResponse{
					Status: istatus,
					Header: header,
					Body:   content.Example,
				})
			}
		}

		handler.handlers[path] = make(map[string][]SwockResponse)
		handler.handlers[path]["GET"] = responses
	}

	if !reflect.ValueOf(method.Post).IsZero() { // Is post request
		var responses []SwockResponse
		for status, response := range method.Post.Responses {
			istatus, _ := strconv.Atoi(status)
			for header, content := range response.Content {
				responses = append(responses, SwockResponse{
					Status: istatus,
					Header: header,
					Body:   content.Example,
				})
			}
		}

		handler.handlers[path] = make(map[string][]SwockResponse)
		handler.handlers[path]["POST"] = responses
	}
	// fmt.Println(handler.handlers["/users"]["GET"])
}

func (handler SwockHandler) Handle(w http.ResponseWriter, r *http.Request) {
	_n := rand.Intn(len(handler.handlers[r.URL.Path][r.Method]))

	response := handler.handlers[r.URL.Path][r.Method][_n]

	w.Header().Set("Content-Type", response.Header)
	w.WriteHeader(response.Status)
	w.Write(response.Body)
}
