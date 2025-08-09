package main

import "testing"

type TestCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []TestCase{
		TestCase{
			input:    "   hello  world    ",
			expected: []string{"hello", "world"},
		},
		TestCase{
			input:    " my       name is   ellie",
			expected: []string{"my", "name", "is", "ellie"},
		},
		TestCase{
			input:    "i   like the  color  lavender",
			expected: []string{"i", "like", "the", "color", "lavender"},
		},
		TestCase{
			input:    "  trans rights  are human   rights!   ",
			expected: []string{"trans", "rights", "are", "human", "rights!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch between %v and %v! got: %d, want: %d", actual, c.expected, len(actual), len(c.expected))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word mismatch between %s and %s!", actual[i], c.expected[i])
			}
		}
	}
}
