// supabase key
package detectors

import (
	"regexp"
)

var supabaseReg = regexp.MustCompile(`sbp_[a-f0-9]{40}`)

func Supabase(src string) (string, bool, string) {
	key := supabaseReg.FindString(src)
	if key == "" {
		return "", false, "supabase"
	}
	return key, true, "supabase"
}

func init() {
	AllDetectors = append(AllDetectors, Supabase)
}
