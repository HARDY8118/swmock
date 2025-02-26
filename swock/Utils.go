package swock

import (
	"math/rand"
	"regexp"
)

const urlRegex = `^https?:\/\/([a-z0-9\.]+)(:\d+)?$`

var urlRegexp = regexp.MustCompile(urlRegex)

func ValidateUrl(url string) bool {
	return urlRegexp.MatchString(url)
}

func Addr(url string) string {
	match := urlRegexp.FindStringSubmatch(url)
	return match[1] + match[2]
}

func RandSelect(v []string) string {
	n := len(v)

	return v[rand.Intn(n)]
}
