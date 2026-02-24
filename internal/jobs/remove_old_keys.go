package jobs

import (
	"openradar/internal/db"
	"openradar/internal/scanner/checks"
	"time"
)

// Runs if finding is > 3 days old & check fails
func removeOldFindingsFunc(jctx JobContext) {
	findings, err := db.GetAllFindings(jctx.DB)
	if err != nil {
		return
	}

	for _, finding := range findings {
		if time.Since(finding.DetectedAt) >= 3*24*time.Hour { // 3 hours x 24 hours
			// run check
			check := checks.RunCheckForProvider(finding.Provider, finding.Key) // lets hope to fucking hell we arent rate limited lmao :sob:
			if !check {                                                        // check is false/invalid
				db.DeleteFinding(finding.ScanJobID, jctx.DB)
				continue
			}
			// do nothing!
		}
	}
}

// func init() {
// 	RegisterJob(Job{
// 		Name:     "Remove invalid keys",
// 		Func:     removeOldFindingsFunc,
// 		Schedule: 9 * time.Hour, // every 2 hours
// 	})
// }
