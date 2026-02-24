package checks

import (
	"fmt"
	"net/http"
)

var googleBase string = "https://generativelanguage.googleapis.com"

// Returns whether resp is a 403 or not.
func Google(key string) bool {
	req, err := http.NewRequest("GET", googleBase+"/v1beta/models?key="+key, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	//ikik we should be retrying/rechecking etc if we get a bad response
	//but i dont think we really need to tbh lmao

	// no auth
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {

		fmt.Println("Status code is NOT OK. Provider: google StatusCode: ", resp.StatusCode)
		return false
	}

	// is authed
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "google",
		Check:    Google,
	})
}
