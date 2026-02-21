// npm
package detectors

import (
	"regexp"
)

var npmKey = regexp.MustCompile(`npm_[a-zA-Z0-9]{36}`)

func npm(src string) (string, bool, string) {
	key := npmKey.FindString(src)
	if key == "" {
		return "", false, "npm"
	}
	return key, true, "npm"
}

func init() {
	AllDetectors = append(AllDetectors, npm)
}
