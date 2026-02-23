package tests

import (
	"openradar/internal/api"
	"testing"
)

func Test(t *testing.T) {

	// Invalid Page
	_, err := api.GetLatestFindings(0, 100, "anthropic", "24h", nil)

	if err == nil {
		t.Errorf("page=0 should have errored!")
	}

	// Invalid PageSize
	_, err = api.GetLatestFindings(1, 0, "anthropic", "24h", nil)

	if err == nil {
		t.Errorf("pagesize=0 should have errored!")
	}

	// Invalid MinAge
	_, err = api.GetLatestFindings(1, 100, "anthropic", "banana", nil)

	if err == nil {
		t.Errorf("minage='banana' should have errored!")
	}

	// Invalid Negative Duration
	_, err = api.GetLatestFindings(1, 100, "anthropic", "-1h", nil)

	if err == nil {
		t.Errorf("minage=-1h should have errored!")
	}

	// Invalid Providers
	_, err = api.GetLatestFindings(1, 100, "flavorpheus", "1h", nil)

	if err == nil {
		t.Errorf("Invalid provider should have errored!")
	}

	// Invalid Repository URL
	_, err = api.GetFindingsFromRepository("", 1, 100, nil)

	if err == nil {
		t.Errorf("Invalid repository URL should have errored!")
	}

	// Invalid page (getallrepositories)
	_, err = api.GetAllRepositories(1, 100, nil)

	if err == nil {
		t.Errorf("page=0 should have errored! (GetAllRepositories)")
	}

	// Invalid RepositoryInfo URL
	_, err = api.GetRepositoryInfo("", nil)

	if err == nil {
		t.Errorf("repo_url should have errored! (GetRepositoryInfo)")
	}
}
