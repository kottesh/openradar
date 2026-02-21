// hugging face keys
package detectors

import (
	"regexp"
)

var hfKey = regexp.MustCompile(`hf_[a-zA-Z]{34}`)

func HuggingFace(src string) (string, bool, string) {
	key := hfKey.FindString(src)
	if key == "" {
		return "", false, "huggingface"
	}
	return key, true, "huggingface"
}

func init() {
	AllDetectors = append(AllDetectors, HuggingFace)
}
