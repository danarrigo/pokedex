package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check slice length
		if len(actual) != len(c.expected) {
			t.Errorf("for input %q expected length %d, got %d (actual=%v)",
				c.input, len(c.expected), len(actual), actual)
			continue
		}

		// Check each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("for input %q at index %d expected %q, got %q (slice=%v)",
					c.input, i, c.expected[i], actual[i], actual)
			}
		}
	}
}
