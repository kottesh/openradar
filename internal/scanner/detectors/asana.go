// asana keys
package detectors

import (
	"regexp"
)

var asana = regexp.MustCompile(`[0-9]{16}:[a-zA-Z0-9]{32}`)

func Asana(src string) (string, bool, string) {
	key := asana.FindString(src)
	if key == "" {
		return "", false, "asana"
	}
	return key, true, "asana"
}

func init() {
	AllDetectors = append(AllDetectors, Asana)
}
