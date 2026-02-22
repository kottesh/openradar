// flavortown keys
package detectors

import (
	"regexp"
)

var ftKey = regexp.MustCompile(`ft_sk_[a-f0-9]{32}`)

func Flavortown(src string) (string, bool, string) {
	Key := ftKey.FindString(src)
	if Key == "" {
		return "", false, "flavortown"
	}
	return Key, true, "flavortown"
}

func init() {
	AllDetectors = append(AllDetectors, Flavortown)
}
