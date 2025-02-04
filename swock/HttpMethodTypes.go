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

type OpenApi3Responses map[string]struct {
	Description string `json:"summary"`
	Content     map[string]struct {
		Schema  json.RawMessage `json:"schema"`
		Example json.RawMessage `json:"example"`
	} `json:"content"`
}

type OpenApi3Get struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Parameters  []OpenApi3Parameter
	Responses   OpenApi3Responses `json:"responses"`
}

type OpenApi3Head struct{}

// TODO change definition
type OpenApi3Post struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Parameters  []OpenApi3Parameter
	RequestBody OpenAPi3RequestBody `json:"requestBody"`
	Responses   OpenApi3Responses   `json:"responses"`
}

type OpenApi3Put struct{}

type OpenApi3Delete struct{}

type OpenApi3Connect struct{}

type OpenApi3Options struct{}

type OpenApi3Trace struct{}

type OpenApi3Patch struct{}
