package checks

//github.go bugged out, so its called git.go

import (
	"fmt"
	"net/http"
)

var github_base string = "https://api.github.com"

func Github(key string) bool {
	req, err := http.NewRequest("GET", github_base+"/user", nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending req:", err)
		return false
	}
	defer resp.Body.Close()

	// no auth
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("Status code is NOT OK. Provider: github StatusCode: ", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "github",
		Check:    Github,
	})
}
