package checks

import (
	"context"
	"math"
	"time"

	"golang.org/x/time/rate"
)

type CheckFunc func(src string) bool

type Check struct {
	Provider string
	Check    CheckFunc
}

var AllChecks []Check

// Rate limits for each provider.
var limiters = map[string]*rate.Limiter{
	"anthropic":   rate.NewLimiter(rate.Every(time.Minute), 10), // not documented
	"asana":       rate.NewLimiter(rate.Every(time.Minute), 1500),
	"aws":         rate.NewLimiter(rate.Every(time.Second), 10000),
	"cerebras":    rate.NewLimiter(rate.Every(time.Minute), 6),
	"cloudflare":  rate.NewLimiter(rate.Every(time.Minute), 240),
	"discord":     rate.NewLimiter(rate.Every(time.Second), 25),
	"flavortown":  rate.NewLimiter(rate.Every(time.Minute), 5),
	"github":      rate.NewLimiter(rate.Every(time.Hour), int(5000-math.Round(((60*60)/35)))), // idk about this since we make requests anyways during scan_job? so lets just choose default and calculate reqs per hour w our backend
	"google":      rate.NewLimiter(rate.Every(time.Minute), 30),
	"groq":        rate.NewLimiter(rate.Every(time.Minute), 30),
	"huggingface": rate.NewLimiter(rate.Every(time.Minute), 200), // 1000 per 5 minutes
	"mistral":     rate.NewLimiter(rate.Every(time.Second), 1),
	"npm":         rate.NewLimiter(rate.Every(time.Minute), 20), // not documented
	"openai":      rate.NewLimiter(rate.Every(time.Minute), 60),
	"openrouter":  rate.NewLimiter(rate.Every(time.Minute), 20),
	"pypi":        rate.NewLimiter(rate.Every(time.Minute), 10), // not documented
	"sendgrid":    rate.NewLimiter(rate.Every(time.Minute), 10),
	"stripe":      rate.NewLimiter(rate.Every(time.Second), 25),
	"telegram":    rate.NewLimiter(rate.Every(time.Second), 30),
	"xai":         rate.NewLimiter(rate.Every(time.Minute), 15), // not documented
}

// Handles waiting
func waitForLimiter(provider string) {
	if l, ok := limiters[provider]; ok {
		l.Wait(context.Background())
	}
}

// true = Does work
// false = Doesnt work

func RunCheckForProvider(prov string, key string) bool {

	// Wait for rait limits
	waitForLimiter(prov)

	for _, Provider := range AllChecks {
		if Provider.Provider == prov {
			return Provider.Check(key)
		}
	}
	return true // return as a fallback
}
