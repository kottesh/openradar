package checks

import (
	"fmt"
	"net/http"
)

var xai_base string = "https://api.x.ai"

func Xai(key string) bool {
	req, err := http.NewRequest("GET", xai_base+"/v1/models", nil)
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

	// no auth
	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	// auth
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "xai",
		Check:    Xai,
	})
}
