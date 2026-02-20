// sendgrid keys
package detectors

import (
	"regexp"
)

var sendgridRegex = regexp.MustCompile(`SG\.[\w_]{16,32}\.[\w_]{16,64}`)

func Sendgrid(src string) (string, bool, string) {
	key := sendgridRegex.FindString(src)
	if key == "" {
		return "", false, "sendgrid"
	}
	return key, true, "sendgrid"
}

func init() {
	AllDetectors = append(AllDetectors, Sendgrid)
}
