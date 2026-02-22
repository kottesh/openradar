package checks

import (
	"fmt"
	"net/http"
)

var cloudflare_base string = "https://api.cloudflare.com"

func Cloudflare(key string) bool {
	req, err := http.NewRequest("GET", cloudflare_base+"/client/v4/user/tokens/verify", nil)
	if err != nil {
		fmt.Println("Err creating request:", err)
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

	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "cloudflare",
		Check:    Cloudflare,
	})
}
