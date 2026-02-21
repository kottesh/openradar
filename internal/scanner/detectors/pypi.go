// pypi
package detectors

import (
	"regexp"
)

var pypiKey = regexp.MustCompile(`pypi-AgEIcHlwaS5vcmc[A-Za-z0-9_-]{50,}`)

func Pypi(src string) (string, bool, string) {
	key := pypiKey.FindString(src)
	if key == "" {
		return "", false, "pypi"
	}
	return key, true, "pypi"
}

func init() {
	AllDetectors = append(AllDetectors, Pypi)
}
