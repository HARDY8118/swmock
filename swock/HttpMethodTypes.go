package swock

import "encoding/json"

type OpenAPi3RequestBody struct {
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Content     map[string]struct {
		Schema  json.RawMessage `json:"schema"`
		Example json.RawMessage `json:"example"`
	} `json:"content"`
}

type OpenApi3ResponseContent map[string]struct {
	Schema  json.RawMessage `json:"schema"`
	Example json.RawMessage `json:"example"`
}

func (oa3rc OpenApi3ResponseContent) ContentTypes() []string {
	var keys []string

	for key, _ := range oa3rc {
		keys = append(keys, key)
	}

	return keys
}

type OpenApi3Responses map[string]struct {
	Description string                  `json:"summary"`
	Content     OpenApi3ResponseContent `json:"content"`
}

func (oa3r OpenApi3Responses) StatusCodes() []string {
	var keys []string

	for key, _ := range oa3r {
		keys = append(keys, key)
	}

	return keys
}

type OpenApi3Method struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Parameters  []OpenApi3Parameter
	RequestBody OpenAPi3RequestBody `json:"requestBody"`
	Responses   OpenApi3Responses   `json:"responses"`
}
