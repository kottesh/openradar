package api

import (
	"math"
	"openradar/internal/domain"

	"gorm.io/gorm"

	"fmt"
)

func GetRepositoryInfo(repo_url string, dbToGrabFrom *gorm.DB) ([]domain.Repository, error) {
	var repository []domain.Repository

	if repo_url == "" {
		return nil, fmt.Errorf("repo_url cannot be empty")
	}

	query := dbToGrabFrom.Where("repo_name = ?", repo_url)

	result := query.Find(&repository)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching repository: %w", result.Error)
	}

	return repository, nil
}

func GetAllRepositories(page int, pageSize int, dbToGrabFrom *gorm.DB) (*domain.PaginatedRepositories, error) {
	// Validate APIs
	if page < 1 {
		return nil, fmt.Errorf("page must be greater than 0")
	}
	if pageSize < 1 || pageSize > 100 {
		return nil, fmt.Errorf("page_size must be between 1 and 100")
	}

	if dbToGrabFrom == nil {
		return nil, fmt.Errorf("database connection is required!")
	}

	var repositories []domain.Repository
	var totalCount int64

	// Validate DB

	if err := dbToGrabFrom.Model(&domain.Repository{}).Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("error counting repositories: %w", err)
	}

	offset := (page - 1) * pageSize

	result := dbToGrabFrom.
		Order("last_updated DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&repositories)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching repositories: %w", result.Error)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	return &domain.PaginatedRepositories{
		Repositories: repositories,
		Page:         page,
		PageSize:     pageSize,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
	}, nil
}
