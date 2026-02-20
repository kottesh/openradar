// stripe keys
package detectors

import (
	"regexp"
)

var stripeRegex = regexp.MustCompile(`sk_(live|test)_[0-9a-zA-Z]{24,}`)

func Stripe(src string) (string, bool, string) {
	key := stripeRegex.FindString(src)
	if key == "" {
		return "", false, "stripe"
	}
	return key, true, "stripe"
}

func init() {
	AllDetectors = append(AllDetectors, Stripe)
}
