package swock

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
)

const urlRegex = `^https?:\/\/([a-z0-9\.]+)(:\d+)?$`

var urlRegexp = regexp.MustCompile(urlRegex)

func readJson(filepath string) []byte {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := io.ReadAll(jsonFile)

	return jsonBytes
}

func ParseSwagger(filepath string) *OpenApi3Definition {

	jsonBytes := readJson(filepath)

	var definiton OpenApi3Definition

	json.Unmarshal(jsonBytes, &definiton)

	return &definiton
}

func ValidateUrl(url string) bool {
	return urlRegexp.MatchString(url)
}

func Addr(url string) string {
	match := urlRegexp.FindStringSubmatch(url)
	return match[1] + match[2]
}
