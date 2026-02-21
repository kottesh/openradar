// github api keys
package detectors

import (
	"regexp"
)

var githubPat = regexp.MustCompile(`ghp_[a-zA-Z0-9]{32}|github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59}`)

func Github(src string) (string, bool, string) {
	key := githubPat.FindString(src)
	if key == "" {
		return "", false, "github"
	}
	return key, true, "github"
}

func init() {
	AllDetectors = append(AllDetectors, Github)
}
