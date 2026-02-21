// cloudflare
package detectors

import (
	"regexp"
)

var cfKey = regexp.MustCompile(`v1\.0-[a-f0-9]{24}-[a-f0-9]{146}`)

func Cloudflare(src string) (string, bool, string) {
	key := cfKey.FindString(src)
	if key == "" {
		return "", false, "cloudflare"
	}
	return key, true, "cloudflare"
}

func init() {
	AllDetectors = append(AllDetectors, Cloudflare)
}
