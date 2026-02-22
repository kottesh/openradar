// openrouter api keys
package detectors

import (
	"regexp"
	"strings"
)

var openaiRegex = regexp.MustCompile(`sk-[a-zA-Z0-9]{20,}`)

func OpenAI(src string) (string, bool, string) {
	key := openaiRegex.FindString(src)
	if key == "" {
		return "", false, "openai"
	}
	if strings.HasPrefix(key, "sk-ant-") || strings.HasPrefix(key, "sk-or-") {
		return "", false, "openai"
	}
	return key, true, "openai"
}

func init() {
	AllDetectors = append(AllDetectors, OpenAI)
}
