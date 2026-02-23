package worker

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"runtime/debug"
	"time"

	"openradar/internal/config"
	"openradar/internal/db"
	"openradar/internal/domain"
	"openradar/internal/queue"
	"openradar/internal/scanner"
	"openradar/internal/server"

	"gorm.io/gorm"
)

func Start(ctx context.Context, conf config.Config, DBtoSaveIn *gorm.DB, Hub *server.Hub) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case job := <-queue.JobQueue:
				log.Printf("Starting to process scan job %s for repository %s", job.ID, job.RepositoryURL)

				job.Status = domain.JobStatusInProgress
				// Scan!
				repo, err := scanner.ScanRepo(context.Background(), job.RepositoryURL, conf.GitHub.Key)
				if err != nil {
					log.Printf("failed to scan repo %s: %v", job.RepositoryURL, err)
					continue
				}

				// Job stuff.
				job.Status = domain.JobStatusCompleted
				job.UpdatedAt = time.Now()

				// Check repository should be scanned
				if repo.Size <= uint(conf.Scanner.MaxRepoSizeMB)*1024 {
					dir, err := os.MkdirTemp("", "openradar-")
					if err != nil {
						log.Printf("failed to create temp dir: %v", err)
						continue
					}

					msg, err := json.Marshal(repo)
					if err != nil {
						log.Printf("Failed to send?")
					}
					Hub.Broadcast <- msg

					addedRepo := domain.NewRepository(
						job.ID,
						job.RepositoryURL,
					)

					// Clone the repository
					cloneCtx, cloneCancel := context.WithTimeout(ctx, 60*time.Second)
					err = CloneRepo(cloneCtx, repo.Clone_Url, dir)
					cloneCancel()

					// Handle errors
					if err != nil {

						os.RemoveAll(dir)
						log.Printf("failed to clone repo %s: %v", job.RepositoryURL, err)
						continue
					}

					// Get scanned file results
					results, err := ScanFiles(dir, conf)
					if err != nil {
						continue
					}

					// Go through all the files!
					for _, file := range results {

						// Run through all the detectors
						for _, result := range RunAllDetectors(file.Content, file.RelPath, repo.Url, DBtoSaveIn) {

							// Save the key!
							SaveKey(*domain.NewFinding(
								addedRepo.ScanJobID,
								repo.Url,
								file.RelPath,
								result.Key,
								result.Provider,
							), DBtoSaveIn, conf)
						}
					}

					// Remove the repository files.
					os.RemoveAll(dir)

					// Save the repository!
					existingRepo, err := db.GetRepositoryByName(job.RepositoryURL, DBtoSaveIn)
					_ = existingRepo

					if err != nil {
						if err := db.AddRepository(addedRepo, DBtoSaveIn); err != nil {
							log.Printf("Failed to save repository: %v", err)
						}
					} else {
						if err := db.UpdateRepository(addedRepo, DBtoSaveIn); err != nil {
							log.Printf("Failed to save repository: %v", err)
						}
					}
				}

				// Free memory
				debug.FreeOSMemory()

				log.Printf("Finished processing scan job %s", repo.Url)
			}
		}
	}()
}
