package checks

import (
	"fmt"
	"net/http"
)

var openai_base string = "https://api.openai.com"

func OpenAI(key string) bool {
	req, err := http.NewRequest("GET", openai_base+"/v1/models", nil)

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

	// no auth
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("Status code is NOT OK. Provider: openai StatusCode: ", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "openai",
		Check:    OpenAI,
	})
}
