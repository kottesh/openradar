package checks

import (
	"fmt"
	"net/http"
)

var mistral_base = "https://api.mistral.ai"

func Mistral(key string) bool {
	request, err := http.NewRequest("GET", mistral_base+"/v1/models", nil)

	if err != nil {
		fmt.Println("Error creating the request:", err)
		return false
	}

	request.Header.Set("Authorization", "Bearer "+key)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {

		fmt.Println("Status code is NOT OK. Provider: ", mistral_base, " StatusCode: ", resp.StatusCode)
		return false
	}

	return false
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "mistral",
		Check:    Mistral,
	})
}
