// telegram
package detectors

import (
	"regexp"
)

var telegramKey = regexp.MustCompile(`[0-9]{8,10}:[a-zA-Z0-9_-]{35}`)

func Telegram(src string) (string, bool, string) {
	key := telegramKey.FindString(src)
	if key == "" {
		return "", false, "telegram"
	}
	return key, true, "telegram"
}

func init() {
	AllDetectors = append(AllDetectors, Telegram)
}
