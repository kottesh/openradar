package tests

import (
	"openradar/internal/scanner/detectors"
	"testing"
)

func TestEnsureKeyIsntSpam(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		isSpam bool
	}{
		{
			name:   "Alphabetic Order",
			input:  "abcdefghijklmnopqrstuvwxyz",
			isSpam: true,
		},
		{
			name:   "Numeral Order",
			input:  "1234567890",
			isSpam: true,
		},
		{
			name:   "Placeholder Key",
			input:  "placeholder",
			isSpam: true,
		},
		{
			name:   "Your API Key!",
			input:  "your_api_key",
			isSpam: true,
		},
		{
			name:   "Repeating Characters",
			input:  "saoipsoadaaaaaaaaaaa",
			isSpam: true,
		},
		{
			name:   "'xxx' key",
			input:  "xxxxxxxxxxxxxxx-xxxxxxxx",
			isSpam: true,
		},
		{
			name:   "Mixed Case Variant (your_api_key + placeholder)",
			input:  "your_api_key-placeholder",
			isSpam: true,
		},

		{
			name:   "Legitimate Key",
			input:  "AIzaSyCno1eYt7UhvSnUkH2Kfz_MtYoJP92Z27c",
			isSpam: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isKeySpam := detectors.EnsureKeyIsntSpam(tc.input)

			// Test & Reverse bool, because function return is weird lmao
			if isKeySpam != !tc.isSpam {
				t.Errorf("expected key %s to be %t, got %t", tc.name, tc.isSpam, isKeySpam)
			}
		})
	}
}
