package checks

import (
	"fmt"
)

type CheckFunc func(src string) bool

type Check struct {
	Provider string
	Check    CheckFunc
}

var AllChecks []Check

// true = Does work
// false = Doesnt work

func RunCheckForProvider(prov string, key string) bool {
	for _, Provider := range AllChecks {
		if Provider.Provider == prov {
			return Provider.Check(key)
		}
	}
	fmt.Println("No provider found for", prov)
	return true // return as a fallback
}
