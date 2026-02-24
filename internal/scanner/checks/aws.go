package checks

import (
	"fmt"
	"net/http"
)

var aws_base string = "https://sts.amazonaws.com"

func Aws(key string) bool {
	req, err := http.NewRequest("GET", aws_base+"?Action=GetCallerIdentity&Version=2011-06-15", nil)

	if err != nil {
		fmt.Println("Error creating the request:", err)
		return false
	}

	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

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
		Provider: "aws",
		Check:    Aws,
	})
}
