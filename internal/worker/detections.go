package worker

import (
	"gorm.io/gorm"

	"openradar/internal/scanner/checks"
	"openradar/internal/scanner/detectors"
)

// This file handles the
// detections for each of the
// findings & provider.

type DetectorResult struct {
	Key      string
	Provider string
}

// this will run the keys through all detectors

func RunAllDetectors(src string, fileName string, repoUrl string, database *gorm.DB) []DetectorResult {
	var results []DetectorResult

	for _, scan_function := range detectors.AllDetectors {
		// execute scan function for provider
		key, is_found, provider := scan_function(src)

		// check key is valid by going through providers api
		// check key isnt spam (aaa, test, 123, etc)
		if is_found &&
			checks.RunCheckForProvider(provider, key) &&
			detectors.EnsureKeyIsntSpam(key) {
			// key has been found
			results = append(results, DetectorResult{Key: key, Provider: provider})
		}
	}

	return results
}
