package detectors

import "strings"

type DetectorFunc func(src string) (string, bool, string)

var AllDetectors []DetectorFunc

// kill me plz uwu
func hasRepeatingChars(s string, threshold int) bool {
	if len(s) == 0 {
		return false
	}
	count := 1
	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {

			count++

			if count >= threshold {
				return true
			}

		} else {
			count = 1
		}
	}
	return false
}

func EnsureKeyIsntSpam(key string) bool {
	lower := strings.ToLower(key)
	if strings.Contains(lower, "your_api_key") || strings.Contains(lower, "placeholder") {
		return false // spam
	}
	if strings.Contains(lower, "abcdefg") {
		return false // spam
	}
	if strings.Contains(lower, "12345678") {
		return false // spam
	}
	if strings.Contains(lower, "123") {
		return false //spam
	}
	if strings.Contains(lower, "abc") {
		return false
	}
	if strings.Contains(lower, "test") {
		return false
	}
	if strings.Contains(lower, "token") {
		return false
	}
	if hasRepeatingChars(lower, 6) {
		return false
	}

	return true // not spam
}
