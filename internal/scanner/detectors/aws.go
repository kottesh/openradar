// amazon web services
package detectors

import (
	"regexp"
)

var awsRegex = regexp.MustCompile(`AKIA[0-9A-Z]{16}`)

func AWS(src string) (string, bool, string) {
	key := awsRegex.FindString(src)
	if key == "" {
		return "", false, "aws"
	}
	return key, true, "aws"
}

func init() {
	AllDetectors = append(AllDetectors, AWS)
}
