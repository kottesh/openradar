package checks

import (
	"fmt"
	"net/http"
)

var groq_base string = "https://api.groq.com"

func Groq(key string) bool {
	req, err := http.NewRequest("GET", groq_base+"/openai/v1/models", nil)
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
		fmt.Println("Status code is NOT OK. Provider: groq StatusCode: ", resp.StatusCode)
		return false
	}

	// auth
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "groq",
		Check:    Groq,
	})
}
