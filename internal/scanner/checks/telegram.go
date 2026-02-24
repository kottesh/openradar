package checks

import (
	"fmt"
	"net/http"
)

var telegram_base string = "https://api.telegram.org"

func Telegram(key string) bool {
	url := fmt.Sprintf("%s/bot%s/getMe", telegram_base, key)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "telegram",
		Check:    Telegram,
	})
}
