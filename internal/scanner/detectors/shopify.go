// shopify
package detectors

import (
	"regexp"
)

var shopifyKey = regexp.MustCompile(`shpat_[a-fA-F0-9]{32}`)
var shopifySecretKey = regexp.MustCompile(`shpss_[a-fA-F0-9]{32}`)

func Shopify(src string) (string, bool, string) {
	if key := shopifyKey.FindString(src); key != "" {
		return key, true, "shopify"
	}
	if key := shopifySecretKey.FindString(src); key != "" {
		return key, true, "shopify"
	}
	return "", false, "shopify"
}

func init() {
	AllDetectors = append(AllDetectors, Shopify)
}
