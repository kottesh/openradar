package checks

import (
	"fmt"
	"net/http"
)

var asana_base string = "https://app.asana.com/api"

func Asana(key string) bool {
	request, err := http.NewRequest("GET", asana_base+"/1.0/users/me", nil)

	if err != nil {
		fmt.Println("Error creaing the request:", err)
		return false
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// no authorization
	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "asana",
		Check:    Asana,
	})
}
