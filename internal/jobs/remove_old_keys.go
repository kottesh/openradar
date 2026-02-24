package jobs

import (
	"fmt"
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
		if time.Since(finding.DetectedAt) >= 2*24*time.Hour { // 2 x 24 hours
			// run check in corountine
			go func() {
				check := checks.RunCheckForProvider(finding.Provider, finding.Key) // the comment that was here previously is useless. however it was funny. if you want to read it go back some commits or something :sob: WE FIXED RATE LIMITS!
				if !check {                                                        // check is false/invalid
					db.DeleteFinding(finding.ID, jctx.DB)
				} else {
					fmt.Printf("Finding was valid! Provider: %s, Key: %s \n", finding.Provider, finding.Key)
				}
			}()
		}
	}
}

func init() {
	RegisterJob(Job{
		Name:     "Remove invalid/old keys > 2 days old",
		Func:     removeOldFindingsFunc,
		Schedule: 1 * time.Hour, // every hour
	})
}
