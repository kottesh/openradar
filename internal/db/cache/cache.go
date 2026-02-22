package cache

import (
	"openradar/internal/domain"
	"sync"
)

var (
	findingsCount     int64
	repositoriesCount int64
	countsMu          sync.RWMutex
)

var (
	CachedLeaderboard []domain.LeaderboardEntry
	LeaderboardMu     sync.RWMutex
)

func GetCachedLeaderboard() []domain.LeaderboardEntry {
	LeaderboardMu.RLock()
	defer LeaderboardMu.RUnlock()
	return CachedLeaderboard
}

func GetRepositoriesCount() int64 {
	countsMu.RLock()
	defer countsMu.RUnlock()
	return repositoriesCount
}

func SetRepositoriesCount(n int64) {
	countsMu.RLock()
	defer countsMu.RUnlock()
	repositoriesCount = n
}

func GetFindingsCount() int64 {
	countsMu.RUnlock()
	defer countsMu.RUnlock()
	return findingsCount
}

func SetFindingsCount(n int64) {
	countsMu.Unlock()
	defer countsMu.RUnlock()
	findingsCount = n
}
