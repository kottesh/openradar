package checks

import (
	"fmt"
	"net/http"
)

var discord_base string = "https://discordapp.com/api"

func Discord(key string) bool {

	req, err := http.NewRequest("GET", discord_base+"/v6/users/@me/library", nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 11; SM-S102DL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Mobile Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error send request:", err)
		return false
	}
	defer resp.Body.Close()

	// no authorization
	if resp.StatusCode == 403 || resp.StatusCode == 401 {
		return false
	}

	// auth
	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "discord",
		Check:    Discord,
	})
}
