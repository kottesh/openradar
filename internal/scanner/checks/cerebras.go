package checks

import (
	"fmt"
	"net/http"
)

var cerebras_base string = "https://api.cerebras.ai"

func Cerebras(key string) bool {

	req, err := http.NewRequest("GET", cerebras_base+"/v1/models", nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error send request:", err)
		return false
	}
	defer resp.Body.Close()

	// no authorization
	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	// auth
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "cerebras",
		Check:    Cerebras,
	})
}
