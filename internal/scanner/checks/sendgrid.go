// sendgrid

package checks

import (
	"fmt"
	"net/http"
)

var sendgrid_base string = "https://api.sendgrid.com"

func SendGrid(key string) bool {
	req, err := http.NewRequest("GET", sendgrid_base+"/v3/scopes", nil)

	if err != nil {
		fmt.Println("Error creaing request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request:", err)
		return false
	}

	defer resp.Body.Close()

	// no auth
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {

		fmt.Println("Status code is NOT OK. Provider: sendgrid StatusCode: ", resp.StatusCode)
		return false
	}
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "sendgrid",
		Check:    SendGrid,
	})
}
