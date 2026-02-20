// twilio keys
package detectors

import (
	"regexp"
)

var twilioRegex = regexp.MustCompile(`SK[0-9a-fA-F]{32}`)

func Twilio(src string) (string, bool, string) {
	key := twilioRegex.FindString(src)
	if key == "" {
		return "", false, "twilio"
	}
	return key, true, "twilio"
}

func init() {
	AllDetectors = append(AllDetectors, Twilio)
}
