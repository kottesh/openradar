package checks

import (
	"bytes"
	"fmt"
	"net/http"
)

var anthropic_base string = "https://api.anthropic.com"

func AnthropicCheck(key string) {

	var json = []byte(``)

	req, err := http.NewRequest("POST", anthropic_base+"/v1/models", bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

}
