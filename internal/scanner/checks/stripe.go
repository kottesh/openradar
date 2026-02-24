package checks

import (
	"fmt"
	"net/http"
)

var stripe_base string = "https://api.stripe.com"

func Stripe(key string) bool {
	req, err := http.NewRequest("GET", stripe_base+"/v1/balance", nil)

	if err != nil {
		fmt.Println("Error creating request: ", err)
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

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {

		fmt.Println("Status code is NOT OK. Provider: stripe StatusCode: ", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "stripe",
		Check:    Stripe,
	})
}
