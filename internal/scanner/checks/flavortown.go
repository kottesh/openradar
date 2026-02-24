package checks

import (
	"fmt"
	"net/http"
)

var flavortown_base string = "https://flavortown.hackclub.com/api/v1"

func Flavortown(key string) bool {

	req, err := http.NewRequest("GET", flavortown_base, nil)

	if err != nil {
		fmt.Println("Error creating request!", err)
		return false
	}

	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request!", err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		fmt.Println("Non OK status!", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "flavortown",
		Check:    Flavortown,
	})
}
