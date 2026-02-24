package checks

import (
	"fmt"
	"net/http"
)

var openrouter_base string = "https://openrouter.ai/api"

func Openrouter(key string) bool {
	req, err := http.NewRequest("GET", openrouter_base+"/v1/key", nil)
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

	// no authorization
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {

		fmt.Println("Status code is NOT OK. Provider: openrouter StatusCode: ", resp.StatusCode)
		return false
	}

	// has auth
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "openrouter",
		Check:    Openrouter,
	})
}
