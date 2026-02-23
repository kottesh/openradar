package tests

import (
	"openradar/internal/scanner/detectors"
	"testing"
)

func TestAllDetectors(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedKey      string
		expectedProvider string
		shouldFind       bool
	}{
		{
			name:             "OpenAI",
			input:            "sk-pLm8vKqR3nXwT5jYhD2cF9uEoAiBsGz4eN7kVt6rJ0mQpWx",
			expectedKey:      "sk-pLm8vKqR3nXwT5jYhD2cF9uEoAiBsGz4eN7kVt6rJ0mQpWx",
			expectedProvider: "openai",
			shouldFind:       true,
		},
		{
			name:             "Anthropic",
			input:            "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",
			expectedKey:      "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",
			expectedProvider: "anthropic",
			shouldFind:       true,
		},
		{
			name:             "Google",
			input:            "AIzaSyAuYeUI9sNaoXpQCkN_XrXOF34VGWN7oTI",
			expectedKey:      "AIzaSyAuYeUI9sNaoXpQCkN_XrXOF34VGWN7oTI",
			expectedProvider: "google",
			shouldFind:       true,
		},
		{
			name:             "Cerebras",
			input:            "csk-tvjydr5cer5c5y98r9td3e5mh3mv6cxjjendejycepnytnwp",
			expectedKey:      "csk-tvjydr5cer5c5y98r9td3e5mh3mv6cxjjendejycepnytnwp",
			expectedProvider: "cerebras",
			shouldFind:       true,
		},
		{
			name:             "Groq",
			input:            "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",
			expectedKey:      "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",
			expectedProvider: "groq",
			shouldFind:       true,
		},
		{
			name:             "Mistral",
			input:            "mis_mZ4c3qPC6rTeNGyP5BxXAR7JsucxZCpgsuw22ORhcgA89ea1066",
			expectedKey:      "mis_mZ4c3qPC6rTeNGyP5BxXAR7JsucxZCpgsuw22ORhcgA89ea1066",
			expectedProvider: "mistral",
			shouldFind:       true,
		},
		{
			name:             "xAI",
			input:            "xai-SkkXm1m1s1pxUHwz4nlMVSyK8biDct5yof5ja6ms1far6lMUzIs8YRBGL1cxpji79QLEtJRGAwBirNxU",
			expectedKey:      "xai-SkkXm1m1s1pxUHwz4nlMVSyK8biDct5yof5ja6ms1far6lMUzIs8YRBGL1cxpji79QLEtJRGAwBirNxU",
			expectedProvider: "xai",
			shouldFind:       true,
		},
		{
			name:             "OpenRouter",
			input:            "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", // copy the google ai key.
			expectedKey:      "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",
			expectedProvider: "openrouter",
			shouldFind:       true,
		},
		{
			name:             "Slack",
			input:            "xoxb-123456789012-123456789012-abcdefghijklmnop",
			expectedKey:      "xoxb-123456789012-123456789012-abcdefghijklmnop",
			expectedProvider: "slack",
			shouldFind:       true,
		},
		{
			name:             "Stripe",
			input:            "sk_test_51HPv9pLUn2YqaUhg1G9A4sJc1Vv3f8qjY8d7c6b5a4d3e2f1",
			expectedKey:      "sk_test_51HPv9pLUn2YqaUhg1G9A4sJc1Vv3f8qjY8d7c6b5a4d3e2f1",
			expectedProvider: "stripe",
			shouldFind:       true,
		},
		{
			name:             "AWS",
			input:            "AKIAIOSFODNN7EXAMPLE",
			expectedKey:      "AKIAIOSFODNN7EXAMPLE",
			expectedProvider: "aws",
			shouldFind:       true,
		},
		{
			name:             "Twilio",
			input:            "SK1234567890abcdef1234567890abcdef",
			expectedKey:      "SK1234567890abcdef1234567890abcdef",
			expectedProvider: "twilio",
			shouldFind:       true,
		},
		{
			name:             "SendGrid",
			input:            "SG.1234567890abcdef1234567890abcdef.1234567890abcdef1234567890abcdef",
			expectedKey:      "SG.1234567890abcdef1234567890abcdef.1234567890abcdef1234567890abcdef",
			expectedProvider: "sendgrid",
			shouldFind:       true,
		},
		{
			name:             "Asana",
			input:            "1234567890123456:abcdefghijklmnopqrstuvwxyz123456",
			expectedKey:      "1234567890123456:abcdefghijklmnopqrstuvwxyz123456",
			expectedProvider: "asana",
			shouldFind:       true,
		},
		{
			name:       "No key",
			input:      "this string has no key",
			shouldFind: false,
		},
		{
			name:             "Cloudflare",
			input:            "v1.0-a1b2c3d4e5f6a1b2c3d4e5f6-a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1",
			expectedKey:      "v1.0-a1b2c3d4e5f6a1b2c3d4e5f6-a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1",
			expectedProvider: "cloudflare",
			shouldFind:       true,
		},
		{
			name:             "Github (classic)",
			input:            "ghp_R8kN2mPqL5vX9wF3jH7dY0bA6cE4gI1s",
			expectedKey:      "ghp_R8kN2mPqL5vX9wF3jH7dY0bA6cE4gI1s",
			expectedProvider: "github",
			shouldFind:       true,
		},
		{
			name:             "Github (fine-grain)",
			input:            "github_pat_11AAAAAA0aaaaaaaaaaaaa_AbCdEfGhIjKlMnOpQrStUvWxYz0123456789AbCdEfGhIjKlMnOpQrStUvW",
			expectedKey:      "github_pat_11AAAAAA0aaaaaaaaaaaaa_AbCdEfGhIjKlMnOpQrStUvWxYz0123456789AbCdEfGhIjKlMnOpQrStUvW",
			expectedProvider: "github",
			shouldFind:       true,
		},
		{
			name:             "Hugging Face",
			input:            "hf_AbCdEfGhIjKlMnOpQrStUvWxYzAbCdEfGh",
			expectedKey:      "hf_AbCdEfGhIjKlMnOpQrStUvWxYzAbCdEfGh",
			expectedProvider: "huggingface",
			shouldFind:       true,
		},
		{
			name:             "npm",
			input:            "npm_R8kN2mPqL5vX9wF3jH7dY0bA6cE4gI1sZzWw",
			expectedKey:      "npm_R8kN2mPqL5vX9wF3jH7dY0bA6cE4gI1sZzWw",
			expectedProvider: "npm",
			shouldFind:       true,
		},
		{
			name:             "PyPI",
			input:            "pypi-AgEIcHlwaS5vcmcNMKLQPRSJHGFDEWVUYXZnmklqprsjhgfdewvuyxzNMKLQPRSJHGF",
			expectedKey:      "pypi-AgEIcHlwaS5vcmcNMKLQPRSJHGFDEWVUYXZnmklqprsjhgfdewvuyxzNMKLQPRSJHGF",
			expectedProvider: "pypi",
			shouldFind:       true,
		},
		{
			name:             "Shopify (admin)",
			input:            "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",
			expectedKey:      "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",
			expectedProvider: "shopify",
			shouldFind:       true,
		},
		{
			name:             "Shopify (secret)",
			input:            "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",
			expectedKey:      "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",
			expectedProvider: "shopify",
			shouldFind:       true,
		},
		{
			name:             "Supabase",
			input:            "sbp_a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2",
			expectedKey:      "sbp_a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2",
			expectedProvider: "supabase",
			shouldFind:       true,
		},
		{
			name:             "Telegram",
			input:            "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi",
			expectedKey:      "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi",
			expectedProvider: "telegram",
			shouldFind:       true,
		},
		{
			name:             "Flavortown",
			input:            "ft_sk_a2e00dfe20983c39940a86ada414319d",
			expectedKey:      "ft_sk_a2e00dfe20983c39940a86ada414319d",
			expectedProvider: "flavortown",
			shouldFind:       true,
		},
		{
			name:             "Discord",
			input:            "MTExODk2NzQyMDY4NDE2MjU2.GxkTqP.vR8mN3kL9pQwXjY2nDcF5hEoAiBsGz4eN7kV",
			expectedKey:      "MTExODk2NzQyMDY4NDE2MjU2.GxkTqP.vR8mN3kL9pQwXjY2nDcF5hEoAiBsGz4eN7kV",
			expectedProvider: "discord",
			shouldFind:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			found := false
			for _, detector := range detectors.AllDetectors {
				key, ok, provider := detector(tc.input)
				if ok {
					if !tc.shouldFind {
						t.Errorf("found key %s from provider %s when none was expected", key, provider)
					}
					if provider == tc.expectedProvider {
						found = true
						if key != tc.expectedKey {
							t.Errorf("expected key %s, but got %s", tc.expectedKey, key)
						}
					}
				}
			}
			if tc.shouldFind && !found {
				t.Errorf("expected to find key for provider %s, but none was found", tc.expectedProvider)
			}
		})
	}
}
