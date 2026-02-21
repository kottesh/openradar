package api

import (
	"openradar/internal/db/cache"
	"openradar/internal/domain"
)

// findings/repositories are seperately
// kept in different files.

// The following are cached by jobs.

func GetLeaderboardData() []domain.LeaderboardEntry {
	return cache.GetCachedLeaderboard()
}

func GetFindingsCount() int64 {
	return cache.FindingsCount
}

func GetRepositoriesCount() int64 {
	return cache.RepositoriesCount
}
