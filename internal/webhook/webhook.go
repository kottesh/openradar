package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WebhookData struct {
	Key      string
	Provider string

	FilePath string
	RepoUrl  string
}

type discordEmbed struct {
	Title  string         `json:"title"`
	Color  int            `json:"color"`
	Fields []discordField `json:"fields"`
	Footer discordFooter  `json:"footer"`
}

type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type discordFooter struct {
	Text string `json:"text"`
}

type discordPayload struct {
	Embeds []discordEmbed `json:"embeds"`
}

var client = &http.Client{Timeout: 10 * time.Second}

func SendHook(webhookURL string, data WebhookData) error {
	if webhookURL == "" { // url not found
		return nil
	}

	payload := discordPayload{
		Embeds: []discordEmbed{
			{
				Title: "Found Key!",
				Color: 0x2BA84A,
				Fields: []discordField{
					{Name: "Provider", Value: data.Provider, Inline: true},
					{Name: "Key", Value: fmt.Sprintf("`%s`", data.Key), Inline: true},
					{Name: "File", Value: fmt.Sprintf("`%s`", data.FilePath), Inline: false},

					{Name: "Repository", Value: data.RepoUrl, Inline: false},
				},
				Footer: discordFooter{Text: "RadarBot"},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Failed to send webhook: %w", err)
	}

	resp, err := client.Post(webhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("Failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return fmt.Errorf("Rate limit")
	}

	return nil
}
