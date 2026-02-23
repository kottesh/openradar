package worker

// This handles saving/db functions for worker,
// along with webhook.

import (
	"openradar/internal/config"
	"openradar/internal/db"
	"openradar/internal/domain"
	"openradar/internal/webhook"
	"strings"

	"gorm.io/gorm"
)

// Saves the key if not exists.
func SaveKey(finding domain.Finding, database *gorm.DB, conf config.Config) {

	_, err := db.GetFindingByKey(finding.Key, database)

	// Entry doesnt exist!
	if err != nil {

		// Send the webhook
		webhook.SendHook(conf.HTTP.Webhook, webhook.WebhookData{
			Key:      finding.Key,
			Provider: finding.Provider,
			FilePath: finding.FilePath,
			RepoUrl:  strings.Replace(finding.RepoName, "api.github.com/repos", "github.com", 1),
		})

		//
		db.AddFinding(&finding, database)
	}
}
