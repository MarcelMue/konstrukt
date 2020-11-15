package validate

import (
	"strconv"
	"testing"
)

func Test_Color(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "case 0: rgb format",
			input:    "rgb(255,0,255)",
			expected: true,
		},
		{
			name:     "case 1: hex format",
			input:    "#f2f2f2",
			expected: true,
		},
		{
			name:     "case 2: malformed hex format",
			input:    "f2f2f2",
			expected: false,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			result := Color(tc.input)

			if result != tc.expected {
				t.Fatal("expected", tc.expected, "got", result)
			}
		})
	}
}
