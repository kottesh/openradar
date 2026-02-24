package checks

import (
	"fmt"
	"net/http"
)

var npm_base string = "https://registry.npmjs.org"

func npm(key string) bool {
	req, err := http.NewRequest("GET", npm_base+"/-/whoami", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	// no auth found
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("Status code is NOT OK. Provider: npm StatusCode: ", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "npm",
		Check:    npm,
	})
}
