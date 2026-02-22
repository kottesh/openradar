// openrouter api keys
package detectors

import (
	"regexp"
)

var openaiRegex = regexp.MustCompile(`sk-[A-Za-z0-9]{20}T3BlbkFJ[A-Za-z0-9]{20}`)

func OpenAI(src string) (string, bool, string) {
	key := openaiRegex.FindString(src)
	if key == "" {
		return "", false, "openai"
	}
	return key, true, "openai"
}

func init() {
	AllDetectors = append(AllDetectors, OpenAI)
}
