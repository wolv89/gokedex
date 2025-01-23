package commands_test

import (
	"testing"

	"github.com/wolv89/gokedex/commands"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "exit",
			expected: []string{"exit"},
		},
		{
			input:    "command - with -param",
			expected: []string{"command", "-", "with", "-param"},
		},
		{
			input:    " lots    of 	  spaced    	   words",
			expected: []string{"lots", "of", "spaced", "words"},
		},
	}

	for _, c := range cases {

		actual := commands.CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Found: %d\nExpecting: %d", len(actual), len(c.expected))
		}

		for i := range actual {

			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Found: %s\nExpecting: %s", word, expectedWord)
			}

		}

	}

}
