package checks

import (
	"fmt"
	"net/http"
)

var pypi_base string = "https://pypi.org"

func PyPI(key string) bool {
	req, err := http.NewRequest("GET", pypi_base+"/manage/projects", nil)

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

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non OK status!", resp.StatusCode)
		return false
	}

	return true
}

func init() {
	AllChecks = append(AllChecks, Check{
		Provider: "pypi",
		Check:    PyPI,
	})
}
