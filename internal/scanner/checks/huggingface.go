package checks

import (
	"fmt"
	"net/http"
)

var hf_base string = "https://huggingface.co/api"

func Huggingface(key string) bool {
	req, err := http.NewRequest("GET", hf_base+"/whoami-v2", nil)

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
	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "huggingface",
		Check:    Huggingface,
	})
}
