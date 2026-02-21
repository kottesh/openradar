package api

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"

	"openradar/internal/domain"
)

func GetLatestFindings(page int, pageSize int, provider string, minAge string, dbToGrabFrom *gorm.DB) (*domain.PaginatedFindings, error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be greater than 0")
	}
	if pageSize < 1 || pageSize > 100 {
		return nil, fmt.Errorf("page_size must be between 1 and 100")
	}

	duration, err := time.ParseDuration(minAge)
	if err != nil {
		return nil, fmt.Errorf("invalid minAge format")
	}

	if duration < 0 || duration > 365*24*time.Hour {
		return nil, fmt.Errorf("minAge must be between 0 and 1 year")
	}

	cutOffTime := time.Now().Add(-duration)

	var findings []domain.Finding
	var totalCount int64

	query := dbToGrabFrom.Model(&domain.Finding{}).Where("detected_at >= ?", cutOffTime)

	if provider != "*" {
		validProviders := map[string]bool{
			"anthropic":   true,
			"cerebras":    true,
			"google":      true,
			"groq":        true,
			"mistral":     true,
			"openai":      true,
			"openrouter":  true,
			"xai":         true,
			"slack":       true,
			"discord":     true,
			"aws":         true,
			"asana":       true,
			"stripe":      true,
			"twilio":      true,
			"sendgrid":    true,
			"cloudflare":  true,
			"github":      true,
			"huggingface": true,
			"npm":         true,
			"pypi":        true,
			"shopify":     true,
			"supabase":    true,
			"telegram":    true,
			"flavortown":  true,
		}
		if !validProviders[provider] {
			return nil, fmt.Errorf("invalid provider: %s", provider)
		}
		query = query.Where("provider = ?", provider)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("error counting findings: %w", err)
	}

	offset := (page - 1) * pageSize

	result := query.
		Order("detected_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&findings)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching findings: %w", result.Error)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	return &domain.PaginatedFindings{
		Findings:   findings,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func GetFindingsFromRepository(repo_url string, page int, pageSize int, dbToGrabFrom *gorm.DB) (*domain.PaginatedFindings, error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be greater than 0")
	}
	if pageSize < 1 || pageSize > 100 {
		return nil, fmt.Errorf("page_size must be between 1 and 100")
	}
	if repo_url == "" {
		return nil, fmt.Errorf("repo_url cannot be empty")
	}

	var findings []domain.Finding
	var totalCount int64

	query := dbToGrabFrom.Model(&domain.Finding{}).Where("repo_name = ?", repo_url)

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("error counting findings: %w", err)
	}

	offset := (page - 1) * pageSize

	result := query.
		Order("detected_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&findings)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching findings: %w", result.Error)
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	return &domain.PaginatedFindings{
		Findings:   findings,
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}
